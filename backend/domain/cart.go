package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type CartRepo interface {
	PostCart(ctx context.Context, cart *swagger.CartInfo) error
	GetCartByID(ctx context.Context, customerId int64, cartId int64, storeId int64) (*[]swagger.CartInfo, error)
	GetByProductID(ctx context.Context, id int64) (*swagger.ProductInfo, error)
	DeleteProduct(ctx context.Context, customerId int64, cartId int64, productId int64) error
}

type CartUsecase interface {
	PostCart(ctx context.Context, cart *swagger.CartInfo) error
	GetTotalPriceByID(ctx context.Context, customerId int64, cartId int64, storeId int64) (int64, error)
	DeleteProduct(ctx context.Context, customerId int64, cartId int64, productId int64) error
}
