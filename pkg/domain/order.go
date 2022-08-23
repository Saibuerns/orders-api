package domain

import "time"

type OrderID uint

type Order struct {
	ID             OrderID
	ClientID       ClientID
	Date           time.Time
	DeliverDate    time.Time
	DeliverAddress Address
	Products       []Product
}
