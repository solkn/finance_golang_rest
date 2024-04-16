package services

import (
	"context"
	"finance/models"
	"finance/repository"
 
	"github.com/google/uuid"
)

type OrgService struct {
	repo repository.OrgRepository
}

func NewOrgService(repo repository.OrgRepository) *OrgService {
	return &OrgService{repo: repo}
}

func (os *OrgService) GetOrgByID(ctx context.Context, id uuid.UUID) (*models.Org, error) {
	return os.repo.GetOrgByID(ctx, id)
}

func (os *OrgService) GetOrgs(ctx context.Context) ([]models.Org, error) {

	result, err := os.repo.GetOrgs(ctx)
	if err != nil {
		return nil, err

	}

	return result, err
}

func (os *OrgService) CreateOrg(ctx context.Context, org *models.Org) (*models.Org, error) {

	err := os.repo.CreateOrg(ctx, org)
	if err != nil {
		return nil, err
	}

	return org, nil
}
func (os *OrgService) UpdateOrg(ctx context.Context, org *models.Org) (*models.Org, error) {
	err := os.repo.UpdateOrg(ctx, org)
	if err != nil {
		return nil, err
	}

	return org, nil

}

func (os *OrgService) DeleteOrg(ctx context.Context, id uuid.UUID) error {
	return os.repo.DeleteOrg(ctx, id)
}
