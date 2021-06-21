package ports

import (
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
)

type CompanyRepository interface {
	Get(id uuid.UUID) (domain.Company, error)
	Save(domain.Company) error
}

type CompanyController interface {
	Get(id uuid.UUID) (domain.Company, error)
	Create(params CreateCompanyController) (domain.Company, error)
}

type CreateCompanyController struct {
	OwnerId      uuid.UUID
	Name         string
	AddressLine1 *string
	AddressLine2 *string
	City         *string
	State        *string
	PostalCode   *string
	Country      *string
}
