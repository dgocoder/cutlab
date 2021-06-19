package ports

import (
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
)

type UserRepository interface {
	Get(id uuid.UUID) (domain.User, error)
	Save(domain.User) error
}

type UserController interface {
	Get(id uuid.UUID) (domain.User, error)
	Create(name string, email string, password string, phone *domain.Phone) (domain.User, error)
}
