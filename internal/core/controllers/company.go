package controllers

import (
	"errors"

	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type CompanyController struct {
	repo ports.CompanyRepository
}

func NewCompanyController(repo ports.CompanyRepository) *CompanyController {
	return &CompanyController{
		repo: repo,
	}
}

func (c *CompanyController) Get(id uuid.UUID) (domain.Company, error) {
	company, err := c.repo.Get(id)
	if err != nil {
		return domain.Company{}, errors.New("get failed")
	}
	return company, nil
}

func (c *CompanyController) Create(name string, email string) (domain.Company, error) {
	company := domain.NewCompany(name, email)

	if err := c.repo.Save(company); err != nil {
		return domain.Company{}, errors.New("create company into repository has failed")
	}
	return company, nil
}
