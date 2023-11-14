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

		quantity := cart.ProductQuantity
		if cart.DiscountId != 1 {
			discountQuantity, err := cu.cartRepo.GetEventDiscountQuantity(ctx, cart.DiscountId)
			if err != nil {
				logrus.Error(err)
				return 0, err
			}
			quantity = cart.ProductQuantity - (cart.ProductQuantity / (discountQuantity + 1))
		}
		totalPrice += product.Price * quantity
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
	totalPrice, errTotal := cu.GetTotalPriceByID(ctx, customerId, cartId, storeId)
	userAddress, errAddress := cu.cartRepo.GetUserAddressById(ctx, customerId)
	for _, err := range []error{errTotal, errAddress} {
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	dt := time.Now().Format("01-02-2006 15:04:05")

	order := &swagger.OrderInfo{
		CustomerId:          customerId,
		SeasoningDiscountId: checkoutInfo.SeasoningDiscountId,
		ShippingDiscountId:  checkoutInfo.ShippingDiscountId,
		CartId:              cartId,
		StoreId:             storeId,
		OrderStatus:         1,
		OrderDate:           dt,
		TakingAddress:       userAddress,
		TakingMethod:        checkoutInfo.TakingMethod,
		TotalPrice:          totalPrice,
	}

	err := cu.cartRepo.AddOrder(ctx, customerId, cartId, storeId, order)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = cu.cartRepo.AddUserCartId(ctx, customerId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (cu *cartUsecase) GetCartArrayByCustomerID(ctx context.Context, id int64) (*[]swagger.CartInfo, error) {
	carts, err := cu.cartRepo.GetCartArrayByCustomerID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return carts, nil
}
