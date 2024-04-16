package services

import (
	"context"
	"finance/models"
	"finance/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return us.repo.GetUserByID(ctx, id)
}

func (us *UserService) GetUsers(ctx context.Context) ([]models.User, error) {

	result, err := us.repo.GetUsers(ctx)
	if err != nil {
		return nil, err

	}

	return result, err
}


func (s *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {

	// Call the repository to create the user
	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (us *UserService) UpdateUser(ctx context.Context, id uuid.UUID, user *models.User) (*models.User, error) {
	err := us.repo.UpdateUser(ctx, id,user)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (us *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return us.repo.DeleteUser(ctx, id)
}
