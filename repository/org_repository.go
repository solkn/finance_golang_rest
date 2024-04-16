package repository

import (
	"context"
	"finance/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrgRepository interface {
	CreateOrg(ctx context.Context, location *models.Org) error
	GetOrgByID(ctx context.Context, ID uuid.UUID) (*models.Org, error)
	GetOrgs(ctx context.Context) ([]models.Org, error)
	UpdateOrg(ctx context.Context, location *models.Org) error
	DeleteOrg(ctx context.Context, ID uuid.UUID) error
} 


type orgRepository struct {
	db *gorm.DB
}

func NewOrgRepository(db *gorm.DB) OrgRepository {
	return &orgRepository{db: db}
}

func (or *orgRepository) CreateOrg(ctx context.Context, org *models.Org) error {
	return or.db.WithContext(ctx).Create(org).Error
}

func (or *orgRepository) GetOrgByID(ctx context.Context, ID uuid.UUID) (*models.Org, error) {
	var org models.Org
	err := or.db.WithContext(ctx).First(&org, ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // Handle user not found case
	}
	return &org, err
}

func (or *orgRepository) GetOrgs(ctx context.Context) ([]models.Org, error) {

	var orgs []models.Org
	result := or.db.WithContext(ctx).Find(&orgs)
	err := result.Error

	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		// Handle case where no users were found (not necessarily an error)
		return nil, nil // Or you can return an empty slice and a specific message
	}

	return orgs, nil

}

func (or *orgRepository) UpdateOrg(ctx context.Context, org *models.Org) error {
	return or.db.WithContext(ctx).Updates(org).Error
}

func (or *orgRepository) DeleteOrg(ctx context.Context, ID uuid.UUID) error {
	return or.db.WithContext(ctx).Delete(&models.Org{}, ID).Error
}
