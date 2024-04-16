package repository

import (
	"context"
	"finance/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TagRepository interface {
	CreateTag(ctx context.Context, tag *models.Tag) error
	GetTagByID(ctx context.Context, ID uuid.UUID) (*models.Tag, error)
	GetTags(ctx context.Context) ([]models.Tag, error)
	UpdateTag(ctx context.Context, tag *models.Tag) error
	DeleteTag(ctx context.Context, ID uuid.UUID) error
} 


type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (tr *tagRepository) CreateTag(ctx context.Context, tag *models.Tag) error {
	return tr.db.WithContext(ctx).Create(tag).Error
}

func (tr *tagRepository) GetTagByID(ctx context.Context, ID uuid.UUID) (*models.Tag, error) {
	var tag models.Tag
	err := tr.db.WithContext(ctx).First(&tag, ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil 
	}
	return &tag, err
}

func (tr *tagRepository) GetTags(ctx context.Context) ([]models.Tag, error) {

	var tags []models.Tag
	result := tr.db.WithContext(ctx).Find(&tags)
	err := result.Error

	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, nil 
	}

	return tags, nil

}

func (tr *tagRepository) UpdateTag(ctx context.Context, tag *models.Tag) error {
	return tr.db.WithContext(ctx).Updates(tag).Error
}

func (tr *tagRepository) DeleteTag(ctx context.Context, ID uuid.UUID) error {
	return tr.db.WithContext(ctx).Delete(&models.Tag{}, ID).Error
}
