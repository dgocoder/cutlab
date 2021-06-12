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

func (srv *ServiceController) Get(id uuid.UUID) (domain.Service, error) {
	service, err := srv.repo.Get(id)
	if err != nil {
		return domain.Service{}, errors.New("get failed")
	}
	return service, nil
}

func (srv *ServiceController) Create(companyId uuid.UUID, locationId uuid.UUID, name string, description *string, duration int, feeAmount int, cancellationFeeAmount int, CancellationCutOff int) (domain.Service, error) {
	service := domain.NewService(companyId, locationId, name, description, duration, feeAmount, cancellationFeeAmount, CancellationCutOff)

	if err := srv.repo.Save(service); err != nil {
		return domain.Service{}, errors.New("create service into repository has failed")
	}
	return service, nil
}
