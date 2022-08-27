package dto

import (
	"database/sql"
	"time"

	"orders.api/pkg/domain"
)

type ProductDBItem struct {
	ID          domain.ProductID `db:"id"`
	Code        sql.NullString   `db:"code"`
	Description sql.NullString   `db:"description"`
	Price       sql.NullFloat64  `db:"price"`
	DateCreated time.Time        `db:"date_created"`
	LastUpdate  time.Time        `db:"last_update"`
}

func (p ProductDBItem) ToProduct() *domain.Product {
	product := &domain.Product{ID: p.ID, LastUpdate: p.LastUpdate}

	if p.Code.Valid {
		product.Code = p.Code.String
	}

	if p.Description.Valid {
		product.Description = p.Description.String
	}

	if p.Price.Valid {
		product.Price = p.Price.Float64
	}

	return product
}
