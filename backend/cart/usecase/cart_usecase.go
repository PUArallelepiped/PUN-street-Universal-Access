package usecase

import (
	"context"
	"fmt"

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

func (cu *cartUsecase) PostCart(ctx context.Context, cart *swagger.CartInfo) error {
	err := cu.cartRepo.PostCart(ctx, cart)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (cu *cartUsecase) GetTotalPriceByID(ctx context.Context, customerId int64, cartId int64, storeId int64) (int64, error) {
	var totalPrice int64 = 0

	carts, err := cu.cartRepo.GetCartByID(ctx, customerId, cartId, storeId)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	for _, cart := range *carts {
		// var productRepo domain.ProductRepo
		// product, err := productRepo.GetByProductID(ctx, cart.ProductId)
		fmt.Println(cart)
		product, err := cu.cartRepo.GetByProductID(ctx, cart.ProductId)
		fmt.Println(product)
		if err != nil {
			logrus.Error(err)
			return 0, err
		}
		totalPrice += product.Price * cart.ProductQuantity
		fmt.Println(totalPrice)

	}

	return totalPrice, nil
}
