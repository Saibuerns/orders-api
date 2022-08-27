package order

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"orders.api/pkg/controller/order/dto"
	"orders.api/pkg/domain"
)

type (
	orderService interface {
		CreateOrder(context.Context, *domain.Order) error
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
}

func (c Controller) Post(ctx *gin.Context) {
	postRequest := &dto.PostRequest{}
	if err := ctx.BindJSON(postRequest); err != nil {
		ctx.JSON(400, errors.New("invalid body"))
	}

	order, err := postRequest.ToOrder()
	if err != nil {
		ctx.JSON(400, err)
	}

	if err := c.orderService.CreateOrder(ctx.Request.Context(), order); err != nil {
		ctx.JSON(500, err)
	}

	ctx.JSON(201, "order created")
}
