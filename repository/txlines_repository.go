package repository

import (
	"context"
	"finance/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TxLinesRepository interface {
	CreateTxLines(ctx context.Context, line *models.TxLines) error
	GetTxLinesByID(ctx context.Context, ID uuid.UUID) (*models.TxLines, error)
	GetTxLines(ctx context.Context) ([]models.TxLines, error)
	UpdateTxLines(ctx context.Context, line *models.TxLines) error
	DeleteTxLines(ctx context.Context, ID uuid.UUID) error
} 


type txLinesRepository struct {
	db *gorm.DB
}

func NewTxLinesRepository(db *gorm.DB) TxLinesRepository {
	return &txLinesRepository{db: db}
}

func (tr *txLinesRepository) CreateTxLines(ctx context.Context, line *models.TxLines) error {
	return tr.db.WithContext(ctx).Create(line).Error
}

func (tr *txLinesRepository) GetTxLinesByID(ctx context.Context, ID uuid.UUID) (*models.TxLines, error) {
	var line models.TxLines
	err := tr.db.WithContext(ctx).First(&line, ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil 
	}
	return &line, err
}

func (tr *txLinesRepository) GetTxLines(ctx context.Context) ([]models.TxLines, error) {

	var lines []models.TxLines
	result := tr.db.WithContext(ctx).Find(&lines)
	err := result.Error

	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, nil 
	}

	return lines, nil

}

func (tr *txLinesRepository) UpdateTxLines(ctx context.Context, line *models.TxLines) error {
	return tr.db.WithContext(ctx).Updates(line).Error
}

func (tr *txLinesRepository) DeleteTxLines(ctx context.Context, ID uuid.UUID) error {
	return tr.db.WithContext(ctx).Delete(&models.TxLines{}, ID).Error
}
