package ports

import (
	"time"

	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
)

type EventRepository interface {
	Get(id uuid.UUID) (domain.Event, error)
	Save(domain.Event) error
}

type EventController interface {
	Get(id uuid.UUID) (domain.Event, error)
	Create(locationId uuid.UUID, serviceId uuid.UUID, resourceId uuid.UUID, customerId uuid.UUID, startAt time.Time, endAt time.Time, eventType string) (domain.Event, error)
}
