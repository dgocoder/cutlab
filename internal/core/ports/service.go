package ports

import (
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
)

type ServiceRepository interface {
	Get(id uuid.UUID) (domain.Service, error)
	Save(domain.Service) error
}

type ServiceController interface {
	Get(id uuid.UUID) (domain.Service, error)
	Create(params CreateServiceView) (domain.Service, error)
	CreateMany(params []CreateServiceView) ([]domain.Service, error)
}

type CreateServiceView struct {
	CompanyId             uuid.UUID
	LocationId            uuid.UUID
	Name                  string
	Description           *string
	Duration              int
	FeeAmount             int
	CancellationFeeAmount int
	CancellationCutOff    int
}
