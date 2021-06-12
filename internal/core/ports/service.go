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
	Create(companyId uuid.UUID, locationId uuid.UUID, name string, description *string, duration int, feeAmount int, cancellationFeeAmount int, CancellationCutOff int) (domain.Service, error)
}
