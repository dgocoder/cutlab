package controllers

import (
	"errors"

	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type CustomerController struct {
	repo ports.CustomerRepository
}

func NewCustomerController(repo ports.CustomerRepository) *CustomerController {
	return &CustomerController{
		repo: repo,
	}
}

func (srv *CustomerController) Get(id uuid.UUID) (domain.Customer, error) {
	customer, err := srv.repo.Get(id)
	if err != nil {
		return domain.Customer{}, errors.New("get failed")
	}
	return customer, nil
}

func (srv *CustomerController) Create(firstName string, lastName string, email string) (domain.Customer, error) {
	customer := domain.NewCustomer(email, lastName, firstName)

	if err := srv.repo.Save(customer); err != nil {
		return domain.Customer{}, errors.New("create customer into repository has failed")
	}
	return customer, nil
}
