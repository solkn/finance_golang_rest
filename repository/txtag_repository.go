package repository

import (
	"context"
	"finance/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TxTagRepository interface {
	CreateTxTag(ctx context.Context, tag *models.TxTag) error
	GetTxTagByID(ctx context.Context, ID uuid.UUID) (*models.TxTag, error)
	GetTxTags(ctx context.Context) ([]models.TxTag, error)
	UpdateTxTag(ctx context.Context, tag *models.TxTag) error
	DeleteTxTag(ctx context.Context, ID uuid.UUID) error
} 


type txTagRepository struct {
	db *gorm.DB
}

func NewTxTagRepository(db *gorm.DB) TxTagRepository {
	return &txTagRepository{db: db}
}

func (tr *txTagRepository) CreateTxTag(ctx context.Context, tag *models.TxTag) error {
	return tr.db.WithContext(ctx).Create(tag).Error
}

func (tr *txTagRepository) GetTxTagByID(ctx context.Context, ID uuid.UUID) (*models.TxTag, error) {
	var tag models.TxTag
	err := tr.db.WithContext(ctx).First(&tag, ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil 
	}
	return &tag, err
}

func (tr *txTagRepository) GetTxTags(ctx context.Context) ([]models.TxTag, error) {

	var tags []models.TxTag
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

func (tr *txTagRepository) UpdateTxTag(ctx context.Context, tag *models.TxTag) error {
	return tr.db.WithContext(ctx).Updates(tag).Error
}

func (tr *txTagRepository) DeleteTxTag(ctx context.Context, ID uuid.UUID) error {
	return tr.db.WithContext(ctx).Delete(&models.TxTag{}, ID).Error
}
