package services

import (
	"context"
	"finance/models"
	"finance/repository"
 
	"github.com/google/uuid"
)

type TxCategoryService struct {
	repo repository.TxCategoryRepository
}

func NewTxCategoryService(repo repository.TxCategoryRepository) *TxCategoryService {
	return &TxCategoryService{repo: repo}
}

func (ts *TxCategoryService) GetTxCategoryByID(ctx context.Context, id uuid.UUID) (*models.TransactionCategory, error) {
	return ts.repo.GetTxCategoryByID(ctx, id)
}

func (ts *TxCategoryService) GetTxCategorys(ctx context.Context) ([]models.TransactionCategory, error) {

	result, err := ts.repo.GetTxCategorys(ctx)
	if err != nil {
		return nil, err

	}

	return result, err
}

func (ts *TxCategoryService) CreateTxCategory(ctx context.Context, txc *models.TransactionCategory) (*models.TransactionCategory, error) {

	err := ts.repo.CreateTxCategory(ctx, txc)
	if err != nil {
		return nil, err
	}

	return txc, nil
}
func (ts *TxCategoryService) UpdateTxCategory(ctx context.Context, txc *models.TransactionCategory) (*models.TransactionCategory, error) {
	err := ts.repo.UpdateTxCategory(ctx, txc)
	if err != nil {
		return nil, err
	}

	return txc, nil

}

func (ts *TxCategoryService) DeleteTxCategory(ctx context.Context, id uuid.UUID) error {
	return ts.repo.DeleteTxCategory(ctx, id)
}
