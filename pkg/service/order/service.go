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
	addressRepository interface {
		GetByHash(context.Context, string) (*domain.Address, error)
		Save(context.Context, *domain.Address) (*domain.Address, error)
	}

	Service struct {
		orderRepo   orderRepository
		addressRepo addressRepository
	}
)

func NewService(or orderRepository, ar addressRepository) (*Service, error) {
	if or == nil {
		return nil, errors.New("order repository can't be nil")
	}
	if ar == nil {
		return nil, errors.New("address repository can't be nil")
	}

	return &Service{
		orderRepo:   or,
		addressRepo: ar,
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
			return s.addressRepo.Save(ctx, &address)
		}
		return nil, err
	}
	return savedAddress, nil
}
