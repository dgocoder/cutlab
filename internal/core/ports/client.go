package ports

import (
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
)

type ClientRepository interface {
	Get(id uuid.UUID) (domain.Client, error)
	Save(domain.Client) error
}

type ClientController interface {
	Get(id uuid.UUID) (domain.Client, error)
	Create(lastName string, firstName string, email string) (domain.Client, error)
}
