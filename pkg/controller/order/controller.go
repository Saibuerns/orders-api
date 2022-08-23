package order

import (
	"context"
	"errors"

	"orders.api/pkg/domain"
)

type (
	orderService interface {
		Save(context.Context, domain.Order) error
	}

	Controller struct {
		orderService orderService
	}
)

func NewController(orderService orderService) (*Controller, error) {
	if orderService == nil {
		return nil, errors.New("order service can't be nil")
	}

	return &Controller{
		orderService: orderService,
	}, nil

	// test push commit
}
