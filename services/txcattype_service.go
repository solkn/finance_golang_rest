package services

import (
	"context"
	"finance/models"
	"finance/repository"
 
	"github.com/google/uuid"
)

type TxCatTypeService struct {
	repo repository.TxCatTypeRepository
}

func NewTxCatTypeService(repo repository.TxCatTypeRepository) *TxCatTypeService {
	return &TxCatTypeService{repo: repo}
}

func (ts *TxCatTypeService) GetTxCatTypeByID(ctx context.Context, id uuid.UUID) (*models.TxCategoryType, error) {
	return ts.repo.GetTxCatTypeByID(ctx, id)
}

func (ts *TxCatTypeService) GetTxCatTypes(ctx context.Context) ([]models.TxCategoryType, error) {

	result, err := ts.repo.GetTxCatTypes(ctx)
	if err != nil {
		return nil, err

	}

	return result, err
}

func (ts *TxCatTypeService) CreateTxCatType(ctx context.Context, txc *models.TxCategoryType) (*models.TxCategoryType, error) {

	err := ts.repo.CreateTxCatType(ctx, txc)
	if err != nil {
		return nil, err
	}

	return txc, nil
}
func (ts *TxCatTypeService) UpdateTxCatType(ctx context.Context, txc *models.TxCategoryType) (*models.TxCategoryType, error) {
	err := ts.repo.UpdateTxCatType(ctx, txc)
	if err != nil {
		return nil, err
	}

	return txc, nil

}

func (ts *TxCatTypeService) DeleteTxCatType(ctx context.Context, id uuid.UUID) error {
	return ts.repo.DeleteTxCatType(ctx, id)
}
