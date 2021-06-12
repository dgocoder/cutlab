package ports

import "github.com/google/uuid"

type PaymentGateway interface {
	CreateCustomer(id uuid.UUID) error
}
