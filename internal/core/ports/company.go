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
	Create(name string, email string) (domain.Company, error)
}
