package dto

import (
	"errors"
	"strings"
	"time"

	"orders.api/pkg/domain"
)

type deliverDate int64

func (dd deliverDate) isValid() bool {
	deliverDate := time.Unix(int64(dd), 0)

	if !deliverDate.IsZero() && time.Now().Sub(deliverDate) >= 2 {
		return true
	}
	return false
}

type deliverAddress struct {
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
	Comments     string `json:"comment,omitempty"`
}

func (da deliverAddress) toAddress() (*domain.Address, error) {
	da.StreetName = strings.TrimSpace(da.StreetName)
	da.StreetNumber = strings.TrimSpace(da.StreetNumber)

	if len(da.StreetName) > 0 && len(da.StreetNumber) > 0 {
		return nil, errors.New("invalid deliver address")
	}

	return &domain.Address{
		StreetName:   da.StreetName,
		StreetNumber: da.StreetNumber,
		Comments:     da.Comments,
	}, nil
}

type PostRequest struct {
	ClientID       domain.ClientID    `json:"client_id"`
	DeliverDate    deliverDate        `json:"deliver_date"`
	DeliverAddress deliverAddress     `json:"deliver_address"`
	ProductIDs     []domain.ProductID `json:"products"`
}

func (pr PostRequest) ToOrder() (*domain.Order, error) {
	if pr.ClientID == 0 {
		return nil, errors.New("invalid client id")
	}
	if !pr.DeliverDate.isValid() {
		return nil, errors.New("invalid deliver date")
	}
	deliverAddress, err := pr.DeliverAddress.toAddress()
	if err != nil {
		return nil, err
	}
	products, err := pr.toProducts()
	if err != nil {
		return nil, err
	}

	return &domain.Order{
		Client:         domain.Client{ID: pr.ClientID},
		DeliverDate:    time.Unix(int64(pr.DeliverDate), 0),
		DeliverAddress: *deliverAddress,
		Products:       products,
	}, nil
}

func (pr PostRequest) toProducts() ([]domain.Product, error) {
	for i, id := range pr.ProductIDs {
		if id == 0 {
			pr.ProductIDs = append(pr.ProductIDs[:i], pr.ProductIDs[i:]...)
		}
	}

	if len(pr.ProductIDs) == 0 {
		return nil, errors.New("product ids is empty")
	}

	products := make([]domain.Product, len(pr.ProductIDs))
	for i, id := range pr.ProductIDs {
		products[i] = domain.Product{ID: id}
	}
	return products, nil
}
