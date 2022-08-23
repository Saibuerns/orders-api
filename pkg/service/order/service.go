package order

import (
	"context"
	"errors"

	"orders.api/pkg/domain"
)

type (
	orderRepository interface {
		save(context.Context, domain.Order)
	}

	Service struct {
		orderRepository orderRepository
	}
)

func NewService(orderRepo orderRepository) (*Service, error) {
	if orderRepo == nil {
		return nil, errors.New("order repository can't be nil")
	}

	return &Service{
		orderRepository: orderRepo,
	}, nil
}
