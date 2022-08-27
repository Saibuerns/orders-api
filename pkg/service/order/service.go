package order

import (
	"context"
	"errors"

	"orders.api/pkg/domain"
	apiErrors "orders.api/pkg/infrastructure/errors"
)

type (
	orderRepository interface {
		Save(context.Context, *domain.Order) error
	}
	clientRepository interface {
		Get(context.Context, domain.ClientID) (*domain.Client, error)
	}
	addressRepository interface {
		GetByHash(context.Context, string) (*domain.Address, error)
		Save(context.Context, domain.Address) (*domain.Address, error)
	}
	productRepository interface {
		GetByIDs(context.Context, []domain.ProductID) ([]domain.Product, error)
	}

	Service struct {
		orderRepo   orderRepository
		clientRepo  clientRepository
		addressRepo addressRepository
		productRepo productRepository
	}
)

func NewService(or orderRepository, cr clientRepository, ar addressRepository, pr productRepository) (*Service, error) {
	if or == nil {
		return nil, errors.New("order repository can't be nil")
	}
	if cr == nil {
		return nil, errors.New("client repository can't be nil")
	}
	if ar == nil {
		return nil, errors.New("address repository can't be nil")
	}
	if pr == nil {
		return nil, errors.New("product repository can't be nil")
	}

	return &Service{
		orderRepo:   or,
		clientRepo:  cr,
		addressRepo: ar,
		productRepo: pr,
	}, nil
}

func (s Service) CreateOrder(ctx context.Context, order *domain.Order) error {
	address, err := s.getAddress(ctx, order.DeliverAddress)
	if err != nil {
		return err
	}

	order.DeliverAddress = *address

	return s.orderRepo.Save(ctx, order)
}

// getAddress get the address by hash if this not exists, is saved and then gets by hash again
func (s Service) getAddress(ctx context.Context, address domain.Address) (*domain.Address, error) {
	savedAddress, err := s.addressRepo.GetByHash(ctx, address.Hash())
	if err != nil {
		if errors.As(err, &apiErrors.NotFoundError{}) {
			return s.addressRepo.Save(ctx, address)
		}
		return nil, err
	}
	return savedAddress, nil
}
