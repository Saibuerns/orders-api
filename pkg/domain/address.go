package domain

import (
	"crypto/sha256"
	"fmt"
)

type AddressID uint

type Address struct {
	ID           AddressID
	StreetName   string
	StreetNumber string
	Comments     string
}

func (a Address) Hash() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s-%s", a.StreetName, a.StreetNumber)))

	return string(h.Sum(nil))
}
