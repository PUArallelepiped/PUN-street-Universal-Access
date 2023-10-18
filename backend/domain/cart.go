package domain

import (
	"context"
)

type Cart struct {
	CustomerID      string `json:"customer_id"`
	ProductID       string `json:"product_id"`
	StoreID         string `json:"store_id"`
	ProductQuantity string `json:"product_quantity"`
}

type CartRepo interface {
	GetByID(ctx context.Context, customerId string, productId string, storeId string) (*Cart, error)
	PostCart(ctx context.Context, cart *Cart) (*Cart, error)
}

type CartUsecase interface {
	GetByID(ctx context.Context, customerId string, productId string, storeId string) (*Cart, error)
	PostCart(ctx context.Context, cart *Cart) (*Cart, error)
}
