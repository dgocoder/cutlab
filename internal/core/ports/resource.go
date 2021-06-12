package ports

import (
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
)

type ResourceRepository interface {
	Get(id uuid.UUID) (domain.Resource, error)
	Save(domain.Resource) error
}

type ResourceController interface {
	Get(id uuid.UUID) (domain.Resource, error)
	Create(name string, locationId uuid.UUID, companyId uuid.UUID, email string) (domain.Resource, error)
}
