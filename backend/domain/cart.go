package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type CartRepo interface {
	IsExitsOrderByStoreCartId(ctx context.Context, customerId int64, storeId int64) (bool, error)
	IsExitsCartByStoreCartId(ctx context.Context, customerId int64, storeId int64) (bool, error)
	IsProductCanAdd(ctx context.Context, id int64) (bool, error)

	DeleteProduct(ctx context.Context, customerId int64, productId int64) (int64, error)
	DeleteOrder(ctx context.Context, customerId int64, storeId int64) error

	GetAllHistoryById(ctx context.Context, id int64) (*[]swagger.HistoryInfo, error)
	GetRunOrderByID(ctx context.Context, id int64) (*[]swagger.RunOrderInfo, error)
	AddProductToCart(ctx context.Context, customerId int64, cartInfo *swagger.CartInfo) error
	AddOrderByCartInfo(ctx context.Context, customerId int64, productId int64) error
}

type CartUsecase interface {
	GetAllHistoryById(ctx context.Context, id int64) (*[]swagger.HistoryInfo, error)
	GetRunOrderByID(ctx context.Context, id int64) (*[]swagger.RunOrderInfo, error)
	DeleteProduct(ctx context.Context, customerId int64, productId int64) error
	AddProductToCart(ctx context.Context, customerId int64, cartInfo *swagger.CartInfo) error
}
