package companies

import (
	"context"
	"xmgo/pkg/models"
	"xmgo/pkg/repo"
)

// ICompanies is a companies usecase interface
type ICompanies interface {
	GetAll() string
	Get(ctx context.Context, id uint) (models.Company, error)
}

type usecase struct {
	repo repo.IRepository
}

// NewUseCase creates an instance of companies usecase
func NewUseCase(repo repo.IRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}

// GetAll returns all companies records
func (uc *usecase) GetAll() string {
	return "all"
}

// Get returns company record by ID
func (uc *usecase) Get(ctx context.Context, id uint) (models.Company, error) {
	var err error
	company, err := uc.repo.GetCompany(id)

	return company, err
}
