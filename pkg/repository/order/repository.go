package order

import (
	"context"
	"database/sql"
	"errors"

	"orders.api/pkg/domain"
)

const parseSQLDateLayout = "2006-01-02"

const ()

type (
	sqlClient interface {
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
	return nil
}
