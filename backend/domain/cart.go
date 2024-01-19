package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/golang-jwt/jwt"
)

type CartRepo interface {
	IsExitsOrderByStoreCartId(ctx context.Context, customerId int64, storeId int64) (bool, error)
	IsExitsCartByStoreCartId(ctx context.Context, customerId int64, storeId int64) (bool, error)
	IsProductCanAdd(ctx context.Context, id int64) (bool, error)

	DeleteProduct(ctx context.Context, customerId int64, productId int64) (int64, error)
	DeleteOrder(ctx context.Context, customerId int64, storeId int64) error

	GetAllHistoryById(ctx context.Context, id int64) (*[]swagger.HistoryInfo, error)
	GetRunOrderByID(ctx context.Context, id int64) (*[]swagger.RunOrderInfo, error)
	GetHistoryCart(ctx context.Context, customerId int64, cartId int64, storeId int64) (*swagger.StoreOrderInfo, error)
	GetCurrentCartID(ctx context.Context, id int64) ([]IDs, error)
	GetSellerOrders(ctx context.Context, id int64) (*[]swagger.StoreOrderStatusInfo, error)
	GetSellerOrder(ctx context.Context, userId int64, cartId int64, storeId int64) (*swagger.StoreOrderStatusInfo, error)

	AddProductToCart(ctx context.Context, customerId int64, cartInfo *swagger.CartInfo) error
	AddOrderByCartInfo(ctx context.Context, customerId int64, productId int64) error
	AddUserCurrentCart(ctx context.Context, id int64) error

	UpdateOrderStatusByID(ctx context.Context, customerId int64, cartId int64, storeId int64) error
	CheckoutOrderInfo(ctx context.Context, customerId int64, storeId int64, totalPrice int64) error
}

type CartUsecase interface {
	GetAllHistoryById(ctx context.Context, id int64) (*[]swagger.HistoryInfo, error)
	GetRunOrderByID(ctx context.Context, id int64) (*[]swagger.RunOrderInfo, error)
	DeleteProduct(ctx context.Context, customerId int64, productId int64) error
	AddProductToCart(ctx context.Context, customerId int64, cartInfo *swagger.CartInfo) error
	GetHistoryCart(ctx context.Context, customerId int64, cartId int64, storeId int64) (*swagger.StoreOrderInfo, error)
	GetCurrentCartsByUserID(ctx context.Context, id int64) (*swagger.CartOrderInfo, error)
	Checkout(ctx context.Context, customerId int64) error
	UpdateOrderStatusByID(ctx context.Context, customerId int64, cartId int64, storeId int64) (*swagger.StoreOrderStatusInfo, error)
	GetSellerOrders(ctx context.Context, id int64) (*[]swagger.StoreOrderStatusInfo, error)
}

type IDs struct {
	UserID int64

	CartID int64

	StoreID int64
}

type Claims struct {
	Email     string `json:"email"`
	Id        int64  `json:"id"`
	Authority string `json:"authority"`
	jwt.StandardClaims
}
