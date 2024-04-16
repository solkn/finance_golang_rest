package repository

import (
	"context"
	"finance/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegisterRepository interface {
	CreateRegister(ctx context.Context, regster *models.Register) error
	GetRegisterByID(ctx context.Context, ID uuid.UUID) (*models.Register, error)
	// GetRegisterByName(ctx context.Context, name string) (*models.Register, error)
	GetRegisters(ctx context.Context) ([]models.Register, error)
	UpdateRegister(ctx context.Context, register *models.Register) error
	DeleteRegister(ctx context.Context, ID uuid.UUID) error
} 


type registerRepository struct {
	db *gorm.DB
}

func NewRegisterRepository(db *gorm.DB) RegisterRepository {
	return &registerRepository{db: db}
}

func (rr *registerRepository) CreateRegister(ctx context.Context, register *models.Register) error {
	return rr.db.WithContext(ctx).Create(register).Error
}

func (rr *registerRepository) GetRegisters(ctx context.Context) ([]models.Register, error) {

	var register []models.Register
	result := rr.db.WithContext(ctx).Find(&register)
	err := result.Error

	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		// Handle case where no users were found (not necessarily an error)
		return nil, nil // Or you can return an empty slice and a specific message
	}

	return register, nil

}

func (rr *registerRepository) GetRegisterByID(ctx context.Context, ID uuid.UUID) (*models.Register, error) {
	var register models.Register
	err := rr.db.WithContext(ctx).First(&register, ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // Handle user not found case
	}
	return &register, err
}

// func (rr *registerRepository) GetRegisterByName(ctx context.Context, name string) (*models.Register, error) {
// 	var register models.Register
// 	err := rr.db.WithContext(ctx).First(&register, name).Error
// 	if err == gorm.ErrRecordNotFound {
// 		return nil, nil // Handle user not found case
// 	}
// 	return &register, err
// }

func (rr *registerRepository) UpdateRegister(ctx context.Context, register *models.Register) error {
	return rr.db.WithContext(ctx).Updates(register).Error
}

func (rr *registerRepository) DeleteRegister(ctx context.Context, ID uuid.UUID) error {
	return rr.db.WithContext(ctx).Delete(&models.Register{}, ID).Error
}
