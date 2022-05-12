package domain

import (
	"events-notification/adapters/db"
	"events-notification/models"
)

type CustomerService struct {
}

func (s CustomerService) CreateCustomer(customer models.Customer) error {
	db, err := db.NewDB()
	if err != nil {
		return err
	}

}
