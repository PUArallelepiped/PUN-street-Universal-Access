package domain

import (
	"context"
)

type Cart struct {
	CustomerID      int `json:"customer_id"`
	ProductID       int `json:"product_id"`
	StoreID         int `json:"store_id"`
	ProductQuantity int `json:"product_quantity"`
}

type CartRepo interface {
	GetByID(ctx context.Context, customerId string, productId string, storeId string) (*Cart, error)
	PostCart(ctx context.Context, cart *Cart) (*Cart, error)
}

type CartUsecase interface {
	GetByID(ctx context.Context, customerId string, productId string, storeId string) (*Cart, error)
	PostCart(ctx context.Context, cart *Cart) (*Cart, error)
}
