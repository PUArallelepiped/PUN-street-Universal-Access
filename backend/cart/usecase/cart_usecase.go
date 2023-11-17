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

func (cu *cartUsecase) PostCart(ctx context.Context, cart *swagger.CartInfo, id int64) error {
	errCart := cu.cartRepo.PostCart(ctx, cart, id)
	userAddress, errAddress := cu.cartRepo.GetUserAddressById(ctx, id)
	dt := time.Now().Format("01-02-2006 15:04:05")

	for _, err := range []error{errCart, errAddress} {
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	_, errOrder := cu.cartRepo.GetOrderById(ctx, id, cart.CartId, cart.StoreId)
	if errOrder != nil {
		order := &swagger.OrderInfo{
			CustomerId:          id,
			SeasoningDiscountId: 1,
			ShippingDiscountId:  1,
			CartId:              cart.CartId,
			StoreId:             cart.StoreId,
			OrderStatus:         0,
			OrderDate:           dt,
			TakingAddress:       userAddress,
			TakingMethod:        1,
			TotalPrice:          0,
		}

		err := cu.cartRepo.AddOrder(ctx, id, cart.CartId, cart.StoreId, order)
		if err != nil {
			logrus.Error(err)
			return err
		}
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

	// cal event discount
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

	//cal seasoning and shipping
	order, cartErr := cu.cartRepo.GetOrderById(ctx, customerId, cartId, storeId)
	shippingFee, shippingFeeErr := cu.cartRepo.GetStoreShippingFeeByID(ctx, storeId)
	for _, err := range []error{cartErr, shippingFeeErr} {
		if err != nil {
			logrus.Error(err)
			return 0, err
		}
	}

	// shipping discount
	if order.ShippingDiscountId != 1 {
		max_price, err := cu.cartRepo.GetMaxPriceByID(ctx, order.ShippingDiscountId)
		if err != nil {
			logrus.Error(err)
			return 0, err
		}
		if totalPrice < max_price {
			totalPrice += shippingFee
		}
	}
	// seasoning  discount
	if order.SeasoningDiscountId != 1 {
		discountPercentage, err := cu.cartRepo.GetPercentageByID(ctx, order.SeasoningDiscountId)
		if err != nil {
			logrus.Error(err)
			return 0, err
		}

		totalPrice = int64(float32(totalPrice) * (float32(discountPercentage) / 100))
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
	dt := time.Now().Format("01-02-2006 15:04:05")
	totalPrice, err := cu.GetTotalPriceByID(ctx, customerId, cartId, storeId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = cu.cartRepo.CheckoutOrder(ctx, customerId, cartId, storeId, totalPrice, dt)
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
