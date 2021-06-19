package ports

import "github.com/testd/cutlab/internal/core/domain"

type AuthGateway interface {
	CreateUser(email string, password string, name string) (domain.User, error)
}
