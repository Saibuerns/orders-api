package product

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"orders.api/pkg/domain"
)

type (
	productRepository interface {
		BatchSave(context.Context, []domain.Product) error
		GetAll(context.Context) ([]domain.Product, error)
	}

	Controller struct {
		productRepo productRepository
	}
)

func NewController(productRepository productRepository) (*Controller, error) {
	if productRepository == nil {
		return nil, errors.New("order service can't be nil")
	}

	return &Controller{
		productRepo: productRepository,
	}, nil
}

func (c Controller) Post(ctx *gin.Context) {
	var products []domain.Product
	if err := ctx.BindJSON(products); err != nil {
		ctx.JSON(400, errors.New("invalid body"))
	}

	if err := c.productRepo.BatchSave(ctx.Request.Context(), products); err != nil {
		ctx.JSON(500, err)
	}

	ctx.JSON(201, "products saved")
}

func (c Controller) GetAll(ctx *gin.Context) {
	products, err := c.productRepo.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, err)
	}

	ctx.JSON(200, products)
}
