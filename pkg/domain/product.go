package domain

import "time"

type ProductID uint

type Product struct {
	ID          ProductID `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	LastUpdate  time.Time `json:"last_update"`
}
