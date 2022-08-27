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
	Floor        string
	Appartment   string
	Comments     string
}

func (a Address) Hash() string {
	h := sha256.New()

	stringToHashed := fmt.Sprintf("%s-%s", a.StreetName, a.StreetNumber)
	if len(a.Floor) > 0 && len(a.Appartment) > 0 {
		stringToHashed = fmt.Sprintf("%s-%s-%s-%s", a.StreetName, a.StreetNumber, a.Floor, a.Appartment)
	}

	h.Write([]byte(stringToHashed))

	return string(h.Sum(nil))
}
