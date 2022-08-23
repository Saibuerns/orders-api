package domain

type ClientID uint

type Phone string

type Client struct {
	ID        ClientID
	LastName  string
	FirstName string
	Phone     Phone
}
