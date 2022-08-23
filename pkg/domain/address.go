package domain

type AddressID uint

type Address struct {
	ID           AddressID
	StreetName   string
	StreetNumber string
	Comments     string
}
