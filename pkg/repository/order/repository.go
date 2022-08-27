package order

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"orders.api/pkg/domain"
)

const (
	insertQuery        = `INSERT INTO order (client_id, deliver_date, deliver_address_id, date_created, last_updated) VALUES (?,?,?,?,?)`
	insertProductQuery = `INSERT INTO order_product (order_id, product_id) VALUES (?,?)`
)

type (
	sqlClient interface {
		BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
		Exec(query string, args ...any) (sql.Result, error)
		Query(query string, args ...any) (*sql.Rows, error)
		QueryRow(query string, args ...any) *sql.Row
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

func (r Repository) Save(ctx context.Context, order *domain.Order) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("fail creating tx: %v", err)
		return errors.New("fail creating tx")
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(ctx, insertQuery, order.Client.ID, order.DeliverDate, order.DeliverAddress.ID, time.Now(), time.Now())
	if err != nil {
		log.Printf("fail saving order: %v", err)
		return err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		log.Printf("fail getting order last insert id: %v", err)
		return err
	}

	if err = r.saveProducts(ctx, tx, orderID, order.Products); err != nil {
		return err
	}

	return nil
}

func (r Repository) saveProducts(ctx context.Context, tx *sql.Tx, orderID int64, products []domain.Product) error {
	for i := range products {
		_, err := tx.ExecContext(ctx, insertProductQuery, orderID, products[i].ID)
		if err != nil {
			log.Printf("fail saving order-product: %v", err)
			return err
		}
	}
	return nil
}
