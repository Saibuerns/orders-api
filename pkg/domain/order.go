package domain

import "time"

type OrderID uint

type Order struct {
	ID             OrderID
	Client         Client
	DeliverDate    time.Time
	DeliverAddress Address
	Products       []Product
	DateCreated    time.Time
}
