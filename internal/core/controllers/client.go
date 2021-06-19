package controllers

import (
	"errors"

	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type ClientController struct {
	repo ports.ClientRepository
}

func NewClientController(repo ports.ClientRepository) *ClientController {
	return &ClientController{
		repo: repo,
	}
}

func (c *ClientController) Get(id uuid.UUID) (domain.Client, error) {
	client, err := c.repo.Get(id)
	if err != nil {
		return domain.Client{}, errors.New("get failed")
	}
	return client, nil
}

func (c *ClientController) Create(firstName string, lastName string, email string) (domain.Client, error) {
	client := domain.NewClient(email, lastName, firstName)

	if err := c.repo.Save(client); err != nil {
		return domain.Client{}, errors.New("create client into repository has failed")
	}
	return client, nil
}
