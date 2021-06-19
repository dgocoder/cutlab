package controllers

import (
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type UserController struct {
	userRepository ports.UserRepository
	authGateway    ports.AuthGateway
}

func NewUserController(userRepository ports.UserRepository, authGateway ports.AuthGateway) *UserController {
	return &UserController{
		userRepository: userRepository,
		authGateway:    authGateway,
	}
}

func (c *UserController) Get(id uuid.UUID) (domain.User, error) {
	user, err := c.userRepository.Get(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (c *UserController) Create(name string, email string, password string, phone *domain.Phone) (domain.User, error) {
	user, err := c.authGateway.CreateUser(email, password, name)
	if err != nil {
		return domain.User{}, nil
	}
	err = c.userRepository.Save(user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
