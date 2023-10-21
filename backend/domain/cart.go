package domain

import (
	"context"
)

type Cart struct {
	CustomerId      int64 `json:"customer_id"`
	ProductId       int64 `json:"product_id"`
	StoreId         int64 `json:"store_id"`
	ProductQuantity int64 `json:"product_quantity"`
}

type CartRepo interface {
	GetByID(ctx context.Context, customerId string, productId string, storeId string) (*Cart, error)
	PostCart(ctx context.Context, cart *Cart) (*Cart, error)
}

type CartUsecase interface {
	GetByID(ctx context.Context, customerId string, productId string, storeId string) (*Cart, error)
	PostCart(ctx context.Context, cart *Cart) (*Cart, error)
}
