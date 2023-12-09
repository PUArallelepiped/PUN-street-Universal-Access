package usecase

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/sirupsen/logrus"
)

type cartUsecase struct {
	cartRepo domain.CartRepo
}

func NewCartUsecase(cartRepo domain.CartRepo) domain.CartUsecase {
	return &cartUsecase{
		cartRepo: cartRepo,
	}
}

func (cu *cartUsecase) GetAllHistoryById(ctx context.Context, id int64) (*[]swagger.HistoryInfo, error) {
	historyArray, err := cu.cartRepo.GetAllHistoryById(ctx, id)
	if err != nil {
		return nil, err
	}

	return historyArray, nil
}

func (cu *cartUsecase) GetRunOrderByID(ctx context.Context, id int64) (*[]swagger.RunOrderInfo, error) {
	runOrderArray, err := cu.cartRepo.GetRunOrderByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return runOrderArray, nil
}

func (cu *cartUsecase) DeleteProduct(ctx context.Context, customerId int64, productId int64) error {
	storeId, err := cu.cartRepo.DeleteProduct(ctx, customerId, productId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	exits, err := cu.cartRepo.IsExitsCartByStoreCartId(ctx, customerId, storeId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	if !exits {
		err = cu.cartRepo.DeleteOrder(ctx, customerId, storeId)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	return nil
}
