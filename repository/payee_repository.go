package repository

import (
	"context"
	"finance/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PayeeRepository interface {
	CreatePayee(ctx context.Context, payee *models.Payee) error
	GetPayeeByID(ctx context.Context, ID uuid.UUID) (*models.Payee, error)
	GetPayees(ctx context.Context) ([]models.Payee, error)
	UpdatePayee(ctx context.Context, payee *models.Payee) error
	DeletePayee(ctx context.Context, ID uuid.UUID) error
} 


type payeeRepository struct {
	db *gorm.DB
}

func NewPayeeRepository(db *gorm.DB) PayeeRepository {
	return &payeeRepository{db: db}
}

func (pr *payeeRepository) CreatePayee(ctx context.Context, payee *models.Payee) error {
	return pr.db.WithContext(ctx).Create(payee).Error
}

func (pr *payeeRepository) GetPayeeByID(ctx context.Context, ID uuid.UUID) (*models.Payee, error) {
	var payee models.Payee
	err := pr.db.WithContext(ctx).First(&payee, ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // Handle user not found case
	}
	return &payee, err
}

func (pr *payeeRepository) GetPayees(ctx context.Context) ([]models.Payee, error) {

	var payees []models.Payee
	result := pr.db.WithContext(ctx).Find(&payees)
	err := result.Error

	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		// Handle case where no users were found (not necessarily an error)
		return nil, nil // Or you can return an empty slice and a specific message
	}

	return payees, nil

}

func (pr *payeeRepository) UpdatePayee(ctx context.Context, payee *models.Payee) error {
	return pr.db.WithContext(ctx).Updates(payee).Error
}

func (pr *payeeRepository) DeletePayee(ctx context.Context, ID uuid.UUID) error {
	return pr.db.WithContext(ctx).Delete(&models.Payee{}, ID).Error
}
