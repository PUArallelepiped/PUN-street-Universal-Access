package usecase

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/sirupsen/logrus"
)

type categoryUsecase struct {
	categoryRepo domain.CategoryRepo
}

func NewCategoryUsecase(categoryRepo domain.CategoryRepo) domain.CategoryRepo {
	return &categoryUsecase{
		categoryRepo: categoryRepo,
	}
}

func (cu *categoryUsecase) GetAllCategory(ctx context.Context) (*[]swagger.Category, error) {
	categorys, err := cu.categoryRepo.GetAllCategory(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return categorys, nil
}
