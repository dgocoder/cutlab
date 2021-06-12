package controllers

import (
	"errors"

	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type LocationController struct {
	repo ports.LocationRepository
}

func NewLocationController(repo ports.LocationRepository) *LocationController {
	return &LocationController{
		repo: repo,
	}
}

func (srv *LocationController) Get(id uuid.UUID) (domain.Location, error) {
	location, err := srv.repo.Get(id)
	if err != nil {
		return domain.Location{}, errors.New("get failed")
	}
	return location, nil
}

func (srv *LocationController) Create(companyId uuid.UUID, name string) (domain.Location, error) {
	location := domain.NewLocation(companyId, name)

	if err := srv.repo.Save(location); err != nil {
		return domain.Location{}, errors.New("create location into repository has failed")
	}
	return location, nil
}
