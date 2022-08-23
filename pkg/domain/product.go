package domain

type ProductID uint

type Product struct {
	ID          ProductID
	Code        string
	Description string
}
