package address

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"orders.api/pkg/domain"
	"orders.api/pkg/repository/address/dto"
)

const (
	insertQuery       = `INSERT INTO address (street_name, street_number, floor, appartment, comments, hash) VALUES (?,?,?,?,?,?)`
	selectByHashQuery = `SELECT street_name, street_number, floor, appartment, comments FROM address WHERE hash = ?`
)

type (
	sqlClient interface {
		Exec(query string, args ...any) (sql.Result, error)
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

func (r Repository) Save(ctx context.Context, address *domain.Address) (*domain.Address, error) {
	hash := address.Hash()
	result, err := r.db.Exec(insertQuery, address.StreetName, address.StreetNumber, address.Comments, hash)
	if err != nil {
		log.Printf("fail saving address: %v", err)
		return nil, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		log.Printf("fail getting order last insert id: %v", err)
		return nil, err
	}

	return r.GetByHash(ctx, hash)
}

func (r Repository) GetByHash(ctx context.Context, hash string) (*domain.Address, error) {
	row := r.db.QueryRow(selectByHashQuery, hash)

	dbItem := &dto.AddressDBItem{}
	if err := row.Scan(dbItem); err != nil {
		log.Printf("fail getting address: %v", err)
		return nil, err
	}
	return dbItem.ToAddress(), nil
}
