package usecase

import (
	"context"
	"errors"

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
	// delete product
	storeId, err := cu.cartRepo.DeleteProduct(ctx, customerId, productId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// check now the delete product's store exits in cart
	exits, err := cu.cartRepo.IsExitsCartByStoreCartId(ctx, customerId, storeId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// if not exits delete the store order status = 0
	if !exits {
		err = cu.cartRepo.DeleteOrder(ctx, customerId, storeId)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	return nil
}

func (cu *cartUsecase) AddProductToCart(ctx context.Context, customerId int64, cartInfo *swagger.CartInfo) error {
	//check product status != 2
	status, err := cu.cartRepo.IsProductCanAdd(ctx, cartInfo.ProductId)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if !status {
		return errors.New("The inventory is not enough for the supply")
	}

	// add or update product in cart
	err = cu.cartRepo.AddProductToCart(ctx, customerId, cartInfo)
	if err != nil {
		logrus.Error(err)
		return err
	}
	// check order exits by store_id
	exits, err := cu.cartRepo.IsExitsOrderByStoreCartId(ctx, customerId, cartInfo.StoreId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// if not exits add order status = 0
	if !exits {
		err := cu.cartRepo.AddOrderByCartInfo(ctx, customerId, cartInfo.StoreId)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	return nil
}
