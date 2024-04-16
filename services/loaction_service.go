package services

import (
	"context"
	"finance/models"
	"finance/repository"
 
	"github.com/google/uuid"
)

type LocationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) *LocationService {
	return &LocationService{repo: repo}
}

func (ls *LocationService) GetLocationByID(ctx context.Context, id uuid.UUID) (*models.Location, error) {
	return ls.repo.GetLocationByID(ctx, id)
}

func (ls *LocationService) GetLocations(ctx context.Context) ([]models.Location, error) {

	result, err := ls.repo.GetLocations(ctx)
	if err != nil {
		return nil, err

	}

	return result, err
}

func (ls *LocationService) CreateLocation(ctx context.Context, location *models.Location) (*models.Location, error) {

	err := ls.repo.CreateLocation(ctx, location)
	if err != nil {
		return nil, err
	}

	return location, nil
}
func (ls *LocationService) UpdateLocation(ctx context.Context, location *models.Location) (*models.Location, error) {
	err := ls.repo.UpdateLocation(ctx, location)
	if err != nil {
		return nil, err
	}

	return location, nil

}

func (ls *LocationService) DeleteLocation(ctx context.Context, id uuid.UUID) error {
	return ls.repo.DeleteLocation(ctx, id)
}
