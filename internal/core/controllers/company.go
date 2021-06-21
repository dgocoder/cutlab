package controllers

import (
	"errors"

	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type CompanyController struct {
	repo     ports.CompanyRepository
	userrepo ports.UserRepository
}

func NewCompanyController(repo ports.CompanyRepository, userrepo ports.UserRepository) *CompanyController {
	return &CompanyController{
		repo:     repo,
		userrepo: userrepo,
	}
}

func (c *CompanyController) Get(id uuid.UUID) (domain.Company, error) {
	company, err := c.repo.Get(id)
	if err != nil {
		return domain.Company{}, errors.New("get failed")
	}
	return company, nil
}

func (c *CompanyController) Create(params ports.CreateCompanyController) (domain.Company, error) {
	company := domain.NewCompany(params.Name, params.AddressLine1, params.AddressLine2, params.City, params.State, params.PostalCode, params.Country)

	if err := c.repo.Save(company); err != nil {
		return domain.Company{}, errors.New("create company into repository has failed")
	}

	user, err := c.userrepo.Get(params.OwnerId)
	if err != nil {
		return domain.Company{}, err
	}

	user.AddCompanyId(company.ID)

	if err := c.userrepo.Save(user); err != nil {
		return domain.Company{}, err
	}

	return company, nil
}
