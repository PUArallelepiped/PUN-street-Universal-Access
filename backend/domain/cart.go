package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type CartRepo interface {
	PostCart(ctx context.Context, cart *swagger.CartInfo, id int64) error
	GetCartByID(ctx context.Context, customerId int64, cartId int64, storeId int64) (*[]swagger.CartInfo, error)
	GetByProductID(ctx context.Context, id int64) (*swagger.ProductInfo, error)
	DeleteProduct(ctx context.Context, customerId int64, cartId int64, productId int64) error
	AddOrder(ctx context.Context, customerId int64, cartId int64, storeId int64, order *swagger.OrderInfo) error
	GetUserAddressById(ctx context.Context, id int64) (string, error)
	AddUserCartId(ctx context.Context, id int64) error
	GetEventDiscountQuantity(ctx context.Context, id int64) (int64, error)
	GetCartArrayByCustomerID(ctx context.Context, id int64) (*[]swagger.CartInfo, error)
	GetOrderById(ctx context.Context, customerId int64, cartId int64, storeId int64) (*swagger.OrderInfo, error)
	CheckoutOrder(ctx context.Context, customerId int64, cartId int64, storeId int64, totalPrice int64, orderDate string) error
	GetStoreShippingFeeByID(ctx context.Context, id int64) (int64, error)
	GetMaxPriceByID(ctx context.Context, id int64) (int64, error)
	GetPercentageByID(ctx context.Context, id int64) (int64, error)
	DeleteOrder(ctx context.Context, customerId int64, cartId int64, storeId int64) error
}

type CartUsecase interface {
	PostCart(ctx context.Context, cart *swagger.CartInfo, id int64) error
	GetTotalPriceByID(ctx context.Context, customerId int64, cartId int64, storeId int64) (int64, error)
	DeleteProduct(ctx context.Context, customerId int64, cartId int64, productId int64) error
	Checkout(ctx context.Context, customerId int64, cartId int64, storeId int64, checkoutInfo *swagger.CheckoutInfo) error
	GetCartArrayByCustomerID(ctx context.Context, id int64) (*[]swagger.CartInfo, error)
}
