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
	Create(params CreateLocationView) (domain.Location, error)
}

type CreateLocationView struct {
	CompanyId     uuid.UUID
	Name          *string
	BusinessHours domain.Availability
}
