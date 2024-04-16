package repository

import (
	"context"
	"finance/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TxCatTypeRepository interface {
	CreateTxCatType(ctx context.Context, txc *models.TxCategoryType) error
	GetTxCatTypeByID(ctx context.Context, ID uuid.UUID) (*models.TxCategoryType, error)
	GetTxCatTypes(ctx context.Context) ([]models.TxCategoryType, error)
	UpdateTxCatType(ctx context.Context, txc *models.TxCategoryType) error
	DeleteTxCatType(ctx context.Context, ID uuid.UUID) error
} 


type txCatTypeRepository struct {
	db *gorm.DB
}

func NewTxCatTypeRepository(db *gorm.DB) TxCatTypeRepository {
	return &txCatTypeRepository{db: db}
}

func (tr *txCatTypeRepository) CreateTxCatType(ctx context.Context, txc *models.TxCategoryType) error {
	return tr.db.WithContext(ctx).Create(txc).Error
}

func (tr *txCatTypeRepository) GetTxCatTypeByID(ctx context.Context, ID uuid.UUID) (*models.TxCategoryType, error) {
	var txc models.TxCategoryType
	err := tr.db.WithContext(ctx).First(&txc, ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil 
	}
	return &txc, err
}

func (tr *txCatTypeRepository) GetTxCatTypes(ctx context.Context) ([]models.TxCategoryType, error) {

	var txcs []models.TxCategoryType
	result := tr.db.WithContext(ctx).Find(&txcs)
	err := result.Error

	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, nil 
	}

	return txcs, nil

}

func (tr *txCatTypeRepository) UpdateTxCatType(ctx context.Context, txc *models.TxCategoryType) error {
	return tr.db.WithContext(ctx).Updates(txc).Error
}

func (tr *txCatTypeRepository) DeleteTxCatType(ctx context.Context, ID uuid.UUID) error {
	return tr.db.WithContext(ctx).Delete(&models.TxCategoryType{}, ID).Error
}
