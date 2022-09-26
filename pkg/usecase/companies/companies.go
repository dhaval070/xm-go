package companies

import (
	"context"
	"xmgo/pkg/models"
)

type ICompanies interface {
	GetAll() string
	Get(ctx context.Context, id int) (models.Company, error)
}

type usecase struct {
}

func NewUseCase() *usecase {
	return &usecase{}
}

func (uc *usecase) GetAll() string {
	return "all"
}

func (uc *usecase) Get(ctx context.Context, id int) (models.Company, error) {
	var err error
	return models.Company{
		ID:   id,
		Name: "xm",
	}, err
}
