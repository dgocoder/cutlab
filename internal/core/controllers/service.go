package controllers

import (
	"errors"

	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type ServiceController struct {
	repo ports.ServiceRepository
}

func NewServiceController(repo ports.ServiceRepository) *ServiceController {
	return &ServiceController{
		repo: repo,
	}
}

func (c *ServiceController) Get(id uuid.UUID) (domain.Service, error) {
	service, err := c.repo.Get(id)
	if err != nil {
		return domain.Service{}, errors.New("get failed")
	}
	return service, nil
}

func (c *ServiceController) Create(params ports.CreateServiceView) (domain.Service, error) {
	service := domain.NewService(params.CompanyId, params.LocationId, params.Name, params.Description, params.Duration, params.FeeAmount, params.CancellationFeeAmount, params.CancellationCutOff)

	if err := c.repo.Save(service); err != nil {
		return domain.Service{}, errors.New("create service into repository has failed")
	}
	return service, nil
}

func (c *ServiceController) CreateMany(params []ports.CreateServiceView) ([]domain.Service, error) {
	services := make([]domain.Service, len(params))
	for i, s := range params {
		// TODO: NEED TO ADD BATCH SUPPORT TO REPO PATTERN
		services[i] = domain.NewService(s.CompanyId, s.LocationId, s.Name, s.Description, s.Duration, s.FeeAmount, s.CancellationFeeAmount, s.CancellationCutOff)
		if err := c.repo.Save(services[i]); err != nil {
			return []domain.Service{}, err
		}
	}
	return services, nil
}
