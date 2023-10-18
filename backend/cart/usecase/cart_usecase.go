package usecase

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"

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

func (cu *cartUsecase) GetByID(ctx context.Context, customerId string, productId string, storeId string) (*domain.Cart, error) {
	cart, err := cu.cartRepo.GetByID(ctx, customerId, productId, storeId)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return cart, nil
}

func (cu *cartUsecase) PostCart(ctx context.Context, cart *domain.Cart) (*domain.Cart, error) {
	cart, err := cu.cartRepo.PostCart(ctx, cart)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return cart, nil
}
