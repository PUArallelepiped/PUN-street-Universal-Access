package usecase

import (
	"context"
	"time"

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
		// product, err := productRepo.GetByProductID(ctx, cart.ProductId)
		product, err := cu.cartRepo.GetByProductID(ctx, cart.ProductId)
		if err != nil {
			logrus.Error(err)
			return 0, err
		}
		totalPrice += product.Price * cart.ProductQuantity
	}

	return totalPrice, nil
}
func (cu *cartUsecase) DeleteProduct(ctx context.Context, customerId int64, cartId int64, productId int64) error {
	err := cu.cartRepo.DeleteProduct(ctx, customerId, cartId, productId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (cu *cartUsecase) Checkout(ctx context.Context, customerId int64, cartId int64, storeId int64, checkoutInfo *swagger.CheckoutInfo) error {
	totalPrice, err := cu.GetTotalPriceByID(ctx, customerId, cartId, storeId)
	user, err := cu.cartRepo.GetUserById(ctx, customerId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	dt := time.Now().Format("01-02-2006")

	// FIX ME when api merge
	order := &swagger.OrderInfo{
		CustomerId:    customerId,
		DiscountId:    checkoutInfo.SeasoningDiscountId,
		CartId:        cartId,
		StoreId:       storeId,
		OrderStatus:   1,
		OrderDate:     dt,
		TakingAddress: user.Address,
		TakingMethod:  checkoutInfo.TakingMethod,
		TotalPrice:    totalPrice,
	}

	err = cu.cartRepo.AddOrder(ctx, customerId, cartId, storeId, order)

	return nil
}
