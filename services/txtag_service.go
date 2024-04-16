package services

import (
	"context"
	"finance/models"
	"finance/repository"
 
	"github.com/google/uuid"
)

type TxTagService struct {
	repo repository.TxTagRepository
}

func NewTxTagService(repo repository.TxTagRepository) *TxTagService {
	return &TxTagService{repo: repo}
}

func (ts *TxTagService) GetTxTagByID(ctx context.Context, id uuid.UUID) (*models.TxTag, error) {
	return ts.repo.GetTxTagByID(ctx, id)
}

func (ts *TxTagService) GetTxTags(ctx context.Context) ([]models.TxTag, error) {

	result, err := ts.repo.GetTxTags(ctx)
	if err != nil {
		return nil, err

	}

	return result, err
}

func (ts *TxTagService) CreateTxTag(ctx context.Context, tag *models.TxTag) (*models.TxTag, error) {

	err := ts.repo.CreateTxTag(ctx, tag)
	if err != nil {
		return nil, err
	}

	return tag, nil
}
func (ts *TxTagService) UpdateTxTag(ctx context.Context, tag *models.TxTag) (*models.TxTag, error) {
	err := ts.repo.UpdateTxTag(ctx, tag)
	if err != nil {
		return nil, err
	}

	return tag, nil

}

func (ts *TxTagService) DeleteTxTag(ctx context.Context, id uuid.UUID) error {
	return ts.repo.DeleteTxTag(ctx, id)
}
