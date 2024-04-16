package repository

import (
	"context"
	"fmt"

	"finance/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, ID uuid.UUID) (*models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, user *models.User) error
	DeleteUser(ctx context.Context, ID uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	// return r.db.CreateContext(ctx, user).Error

	return ur.db.WithContext(ctx).Create(user).Error
}

func (ur *userRepository) GetUserByID(ctx context.Context, ID uuid.UUID) (*models.User, error) {
	var user models.User
	err := ur.db.WithContext(ctx).First(&user, ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // Handle user not found case
	}
	return &user, err
}

func (ur *userRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	// var users []models.User
	// err := r.db.WithContext(ctx).Error
	// if err == gorm.ErrRecordNotFound {
	// 	return nil, nil // Handle user not found case
	// }

	// print("users:", users)

	// return users, err

	var users []models.User
	// Execute a query to fetch users (replace with your specific query)
	result := ur.db.WithContext(ctx).Find(&users)
	err := result.Error

	if err != nil {
		// Handle potential errors during query execution
		return nil, err
	}

	if result.RowsAffected == 0 {
		// Handle case where no users were found (not necessarily an error)
		return nil, nil // Or you can return an empty slice and a specific message
	}

	fmt.Println("users:", users)
	// Users have been successfully retrieved
	return users, nil

}

func (ur *userRepository) UpdateUser(ctx context.Context, id uuid.UUID, user *models.User) error {
	// return ur.db.WithContext(ctx).Updates(user).Error
	return ur.db.WithContext(ctx).Where("id = ?", id).Updates(user).Error

}

func (ur *userRepository) DeleteUser(ctx context.Context, ID uuid.UUID) error {
	return ur.db.WithContext(ctx).Delete(&models.User{}, ID).Error

}
