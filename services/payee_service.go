package services

import (
	"context"
	"finance/models"
	"finance/repository"
 
	"github.com/google/uuid"
)

type PayeeService struct {
	repo repository.PayeeRepository
}

func NewPayeeService(repo repository.PayeeRepository) *PayeeService {
	return &PayeeService{repo: repo}
}

func (ps *PayeeService) GetPayeeByID(ctx context.Context, id uuid.UUID) (*models.Payee, error) {
	return ps.repo.GetPayeeByID(ctx, id)
}

func (ps *PayeeService) GetPayees(ctx context.Context) ([]models.Payee, error) {

	result, err := ps.repo.GetPayees(ctx)
	if err != nil {
		return nil, err

	}

	return result, err
}

func (ps *PayeeService) CreatePayee(ctx context.Context, payee *models.Payee) (*models.Payee, error) {

	err := ps.repo.CreatePayee(ctx, payee)
	if err != nil {
		return nil, err
	}

	return payee, nil
}
func (ps *PayeeService) UpdatePayee(ctx context.Context, payee *models.Payee) (*models.Payee, error) {
	err := ps.repo.UpdatePayee(ctx, payee)
	if err != nil {
		return nil, err
	}

	return payee, nil

}

func (ps *PayeeService) DeletePayee(ctx context.Context, id uuid.UUID) error {
	return ps.repo.DeletePayee(ctx, id)
}
