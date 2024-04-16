package repository

import (
	"context"
	"finance/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LocationRepository interface {
	CreateLocation(ctx context.Context, location *models.Location) error
	GetLocationByID(ctx context.Context, ID uuid.UUID) (*models.Location, error)
	GetLocations(ctx context.Context) ([]models.Location, error)
	UpdateLocation(ctx context.Context, location *models.Location) error
	DeleteLocation(ctx context.Context, ID uuid.UUID) error
} 


type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{db: db}
}

func (lr *locationRepository) CreateLocation(ctx context.Context, location *models.Location) error {
	return lr.db.WithContext(ctx).Create(location).Error
}

func (lr *locationRepository) GetLocationByID(ctx context.Context, ID uuid.UUID) (*models.Location, error) {
	var location models.Location
	err := lr.db.WithContext(ctx).First(&location, ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // Handle user not found case
	}
	return &location, err
}

func (lr *locationRepository) GetLocations(ctx context.Context) ([]models.Location, error) {

	var locations []models.Location
	// Execute a query to fetch users (replace with your specific query)
	result := lr.db.WithContext(ctx).Find(&locations)
	err := result.Error

	if err != nil {
		// Handle potential errors during query execution
		return nil, err
	}

	if result.RowsAffected == 0 {
		// Handle case where no users were found (not necessarily an error)
		return nil, nil // Or you can return an empty slice and a specific message
	}

	return locations, nil

}

func (lr *locationRepository) UpdateLocation(ctx context.Context, location *models.Location) error {
	return lr.db.WithContext(ctx).Updates(location).Error
}

func (lr *locationRepository) DeleteLocation(ctx context.Context, ID uuid.UUID) error {
	return lr.db.WithContext(ctx).Delete(&models.User{}, ID).Error
}
