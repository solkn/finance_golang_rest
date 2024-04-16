package repository

import (
	"context"
	"finance/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TxCategoryRepository interface {
	CreateTxCategory(ctx context.Context, txc *models.TransactionCategory) error
	GetTxCategoryByID(ctx context.Context, ID uuid.UUID) (*models.TransactionCategory, error)
	GetTxCategorys(ctx context.Context) ([]models.TransactionCategory, error)
	UpdateTxCategory(ctx context.Context, txc *models.TransactionCategory) error
	DeleteTxCategory(ctx context.Context, ID uuid.UUID) error
} 


type txCategoryRepository struct {
	db *gorm.DB
}

func NewTxCategoryRepository(db *gorm.DB) TxCategoryRepository {
	return &txCategoryRepository{db: db}
}

func (tr *txCategoryRepository) CreateTxCategory(ctx context.Context, txc *models.TransactionCategory) error {
	return tr.db.WithContext(ctx).Create(txc).Error
}

func (tr *txCategoryRepository) GetTxCategoryByID(ctx context.Context, ID uuid.UUID) (*models.TransactionCategory, error) {
	var txc models.TransactionCategory
	err := tr.db.WithContext(ctx).First(&txc, ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil 
	}
	return &txc, err
}

func (tr *txCategoryRepository) GetTxCategorys(ctx context.Context) ([]models.TransactionCategory, error) {

	var txcs []models.TransactionCategory
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

func (tr *txCategoryRepository) UpdateTxCategory(ctx context.Context, txc *models.TransactionCategory) error {
	return tr.db.WithContext(ctx).Updates(txc).Error
}

func (tr *txCategoryRepository) DeleteTxCategory(ctx context.Context, ID uuid.UUID) error {
	return tr.db.WithContext(ctx).Delete(&models.TransactionCategory{}, ID).Error
}
