package services

import (
	"context"
	"finance/models"
	"finance/repository"
 
	"github.com/google/uuid"
)

type TagService struct {
	repo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) *TagService {
	return &TagService{repo: repo}
}

func (ts *TagService) GetTagByID(ctx context.Context, id uuid.UUID) (*models.Tag, error) {
	return ts.repo.GetTagByID(ctx, id)
}

func (ts *TagService) GetTags(ctx context.Context) ([]models.Tag, error) {

	result, err := ts.repo.GetTags(ctx)
	if err != nil {
		return nil, err

	}

	return result, err
}

func (ts *TagService) CreateTag(ctx context.Context, tag *models.Tag) (*models.Tag, error) {

	err := ts.repo.CreateTag(ctx, tag)
	if err != nil {
		return nil, err
	}

	return tag, nil
}
func (ts *TagService) UpdateTag(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
	err := ts.repo.UpdateTag(ctx, tag)
	if err != nil {
		return nil, err
	}

	return tag, nil

}

func (ts *TagService) DeleteTag(ctx context.Context, id uuid.UUID) error {
	return ts.repo.DeleteTag(ctx, id)
}
