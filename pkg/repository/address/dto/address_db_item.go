package dto

import (
	"database/sql"

	"orders.api/pkg/domain"
)

type AddressDBItem struct {
	ID           domain.AddressID `db:"id"`
	StreetName   string           `db:"street_name"`
	StreetNumber string           `db:"street_number"`
	Floor        sql.NullString   `db:"floor"`
	Appartment   sql.NullString   `db:"appartment"`
	Comments     sql.NullString   `db:"comments"`
}

func (a AddressDBItem) ToAddress() *domain.Address {
	address := &domain.Address{
		ID:           a.ID,
		StreetName:   a.StreetName,
		StreetNumber: a.StreetNumber,
	}
	if a.Floor.Valid {
		address.Floor = a.Floor.String
	}
	if a.Appartment.Valid {
		address.Appartment = a.Appartment.String
	}
	if a.Comments.Valid {
		address.Comments = a.Comments.String
	}

	return address
}
