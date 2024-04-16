package services

import (
	"context"
	"finance/models"
	"finance/repository"
 
	"github.com/google/uuid"
)

type RegisterService struct {
	repo repository.RegisterRepository
}

func NewRegisterService(repo repository.RegisterRepository) *RegisterService {
	return &RegisterService{repo: repo}
}

func (rs *RegisterService) GetRegisters(ctx context.Context) ([]models.Register, error) {

	result, err := rs.repo.GetRegisters(ctx)
	if err != nil {
		return nil, err

	}

	return result, err
}

func (rs *RegisterService) GetRegisterByID(ctx context.Context, id uuid.UUID) (*models.Register, error) {
	return rs.repo.GetRegisterByID(ctx, id)
}

func (rs *RegisterService) CreateRegister(ctx context.Context, register *models.Register) (*models.Register, error) {

	err := rs.repo.CreateRegister(ctx, register)
	if err != nil {
		return nil, err
	}

	return register, nil
}
func (rs *RegisterService) UpdateRegister(ctx context.Context, register *models.Register) (*models.Register, error) {
	err := rs.repo.UpdateRegister(ctx, register)
	if err != nil {
		return nil, err
	}

	return register, nil
 
}

func (rs *RegisterService) DeleteRegister(ctx context.Context, id uuid.UUID) error {
	return rs.repo.DeleteRegister(ctx, id)
}
