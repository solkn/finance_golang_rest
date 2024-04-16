package services

import (
	"context"
	"finance/models"
	"finance/repository"
 
	"github.com/google/uuid"
)

type TxLinesService struct {
	repo repository.TxLinesRepository
}

func NewTxLinesService(repo repository.TxLinesRepository) *TxLinesService {
	return &TxLinesService{repo: repo}
}

func (ts *TxLinesService) GetTxLinesByID(ctx context.Context, id uuid.UUID) (*models.TxLines, error) {
	return ts.repo.GetTxLinesByID(ctx, id)
}

func (ts *TxLinesService) GetTxTags(ctx context.Context) ([]models.TxLines, error) {

	result, err := ts.repo.GetTxLines(ctx)
	if err != nil {
		return nil, err

	}

	return result, err
}

func (ts *TxLinesService) CreateTxLines(ctx context.Context, tag *models.TxLines) (*models.TxLines, error) {

	err := ts.repo.CreateTxLines(ctx, tag)
	if err != nil {
		return nil, err
	}

	return tag, nil
}
func (ts *TxLinesService) UpdateTxLines(ctx context.Context, tag *models.TxLines) (*models.TxLines, error) {
	err := ts.repo.UpdateTxLines(ctx, tag)
	if err != nil {
		return nil, err
	}

	return tag, nil

}

func (ts *TxLinesService) DeleteTxLines(ctx context.Context, id uuid.UUID) error {
	return ts.repo.DeleteTxLines(ctx, id)
}
