package repository

import (
	"context"
	"finance/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction []models.Transaction) error
	CreateTransactionRegister(ctx context.Context, regster *models.Register) error
	CreateTransactionPayee(ctx context.Context, payee *models.Payee) error
	CreateTransactionCategory(ctx context.Context, category *models.TransactionCategory) error
	CreateTransactionTag(ctx context.Context, payee *models.Tag) error
	GetTransactionByID(ctx context.Context, ID uuid.UUID) (*models.Transaction, error)
	GetTransactions(ctx context.Context) ([]models.Transaction, error)
	GetRegisterByName(ctx context.Context, name string) (*models.Register, error)
	GetPayeeByName(ctx context.Context, name string) (*models.Payee, error)
	GetTransactionCategoryByName(ctx context.Context, name string) (*models.TransactionCategory, error)
	GetTagByName(ctx context.Context, name string) (*models.Tag, error)
	UpdateTransaction(ctx context.Context, id uuid.UUID, transaction *models.Transaction) error
	DeleteTransaction(ctx context.Context, ID uuid.UUID) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (tr *transactionRepository) CreateTransaction(ctx context.Context, transaction []models.Transaction) error {
	return tr.db.WithContext(ctx).Create(transaction).Error
}
func (tr *transactionRepository) CreateTransactionRegister(ctx context.Context, register *models.Register) error {
	return tr.db.WithContext(ctx).Create(register).Error
}
func (tr *transactionRepository) CreateTransactionPayee(ctx context.Context, payee *models.Payee) error {
	return tr.db.WithContext(ctx).Create(payee).Error
}
func (tr *transactionRepository) CreateTransactionCategory(ctx context.Context, category *models.TransactionCategory) error {
	return tr.db.WithContext(ctx).Create(category).Error
}
func (tr *transactionRepository) CreateTransactionTag(ctx context.Context, tag *models.Tag) error {
	return tr.db.WithContext(ctx).Create(tag).Error
}
func (tr *transactionRepository) GetTransactionByID(ctx context.Context, ID uuid.UUID) (*models.Transaction, error) {
	var transaction models.Transaction
	err := tr.db.WithContext(ctx).Preload("Payee").Preload("Tags").Preload("Register").Preload("TxLines").First(&transaction, ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &transaction, err
}

func (tr *transactionRepository) GetTransactions(ctx context.Context) ([]models.Transaction, error) {

	var transactions []models.Transaction
	// result := tr.db.WithContext(ctx).Find(&transactions)
	result := tr.db.WithContext(ctx).Preload("Payee").Preload("Tags").Preload("Register").Preload("TxLines").Find(&transactions)

	err := result.Error

	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return transactions, nil

}
func (tr *transactionRepository) GetRegisterByName(ctx context.Context, name string) (*models.Register, error) {
	var register models.Register
	err := tr.db.WithContext(ctx).Where("name = ?", name).First(&register).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // Handle user not found case
	} else if err != nil {
		// Handle other errors from Gorm
		return nil, err
	}
	return &register, err
}
func (tr *transactionRepository) GetPayeeByName(ctx context.Context, name string) (*models.Payee, error) {
	var payee models.Payee
	err := tr.db.WithContext(ctx).Where("name = ?", name).First(&payee).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // Handle user not found case
	}
	return &payee, err
}

func (tr *transactionRepository) GetTransactionCategoryByName(ctx context.Context, name string) (*models.TransactionCategory, error) {
	var transactionCategory models.TransactionCategory
	err := tr.db.WithContext(ctx).Where("name = ?", name).First(&transactionCategory).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // Handle user not found case
	}
	return &transactionCategory, err
}

func (tr *transactionRepository) GetTagByName(ctx context.Context, name string) (*models.Tag, error) {
	var tag models.Tag
	err := tr.db.WithContext(ctx).Where("name = ?", name).First(&tag).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // Handle user not found case
	}
	return &tag, err
}

func (tr *transactionRepository) UpdateTransaction(ctx context.Context, id uuid.UUID, transaction *models.Transaction) error {
	return tr.db.WithContext(ctx).Where("id = ?", id).Updates(transaction).Error
}

func (tr *transactionRepository) DeleteTransaction(ctx context.Context, ID uuid.UUID) error {
	return tr.db.WithContext(ctx).Delete(&models.Transaction{}, ID).Error
}
