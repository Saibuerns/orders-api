package product

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"orders.api/pkg/domain"
	"orders.api/pkg/repository/product/dto"
)

const (
	insertQuery    = `INSERT INTO product (code, description, price, date_created, last_updated) VALUES (?,?,?,?,?);`
	selectAllQuery = `SELECT id, code, description, price, date_created, last_updated FROM product;`
)

type (
	sqlClient interface {
		BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
		Query(query string, args ...any) (*sql.Rows, error)
	}

	Repository struct {
		db sqlClient
	}
)

func NewRepository(sqlClient sqlClient) (*Repository, error) {
	if sqlClient == nil {
		return nil, errors.New("sql client can't be nil")
	}
	return &Repository{db: sqlClient}, nil
}

func (r Repository) BatchSave(ctx context.Context, products []domain.Product) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("fail creating tx: %v", err)
		return errors.New("fail creating tx")
	}
	defer tx.Rollback()

	for _, p := range products {
		_, err := tx.ExecContext(ctx, insertQuery, p.Code, p.Description, p.Price, time.Now(), time.Now())
		if err != nil {
			log.Printf("fail saving product: %v", err)
			return err
		}
	}

	return nil
}

func (r Repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.db.Query(selectAllQuery)
	if err != nil {
		log.Printf("fail getting products: %v", err)
		return nil, err
	}

	var dbItems []dto.ProductDBItem
	for rows.Next() {
		dbItem := &dto.ProductDBItem{}
		if err = rows.Scan(dbItem); err != nil {
			log.Printf("fail parsing product db item: %v", err)
			return nil, err
		}
		dbItems = append(dbItems, *dbItem)
	}

	products := make([]domain.Product, len(dbItems))
	for i, item := range dbItems {
		products[i] = *item.ToProduct()
	}

	return products, nil
}
