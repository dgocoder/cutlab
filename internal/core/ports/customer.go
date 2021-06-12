package ports

import (
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
)

type CustomerRepository interface {
	Get(id uuid.UUID) (domain.Customer, error)
	Save(domain.Customer) error
}

type CustomerController interface {
	Get(id uuid.UUID) (domain.Customer, error)
	Create(lastName string, firstName string, email string) (domain.Customer, error)
}
