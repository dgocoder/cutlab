package ports

import "github.com/google/uuid"

type PaymentGateway interface {
	CreateClient(id uuid.UUID) error
}
