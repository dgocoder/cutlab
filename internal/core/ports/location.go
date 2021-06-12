package ports

import (
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
)

type LocationRepository interface {
	Get(id uuid.UUID) (domain.Location, error)
	Save(domain.Location) error
}

type LocationController interface {
	Get(id uuid.UUID) (domain.Location, error)
	Create(companyId uuid.UUID, name string) (domain.Location, error)
}
