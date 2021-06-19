package controllers

import (
	"errors"

	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type ResourceController struct {
	resourceRepository ports.ResourceRepository
}

func NewResourceController(resourceRepository ports.ResourceRepository) *ResourceController {
	return &ResourceController{
		resourceRepository: resourceRepository,
	}
}

func (c *ResourceController) Get(id uuid.UUID) (domain.Resource, error) {
	resource, err := c.resourceRepository.Get(id)
	if err != nil {
		return domain.Resource{}, errors.New("get failed")
	}
	return resource, nil
}

func (c *ResourceController) Create(name string, locationId uuid.UUID, companyId uuid.UUID, email string) (domain.Resource, error) {
	game := domain.NewResource(name, locationId, companyId, email)

	if err := c.resourceRepository.Save(game); err != nil {
		return domain.Resource{}, errors.New("create game into repository has failed")
	}
	return game, nil
}
