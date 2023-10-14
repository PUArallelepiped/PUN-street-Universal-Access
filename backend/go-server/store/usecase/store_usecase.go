package usecase

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"

	"github.com/sirupsen/logrus"
)

type storeUsecase struct {
	storeRepo domain.StoreRepo
}

func NewStoreUsecase(storeRepo domain.StoreRepo) domain.StoreUsecase {
	return &storeUsecase{
		storeRepo: storeRepo,
	}
}

func (su *storeUsecase) GetByID(ctx context.Context, id string) (*domain.Store, error) {
	s, err := su.storeRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}
