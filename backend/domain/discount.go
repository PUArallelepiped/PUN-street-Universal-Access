package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type DiscountRepo interface {
	GetByDiscountID(ctx context.Context, id int64) (*swagger.DiscountInfo, error)
	GetShippingByStoreID(ctx context.Context, id int64) ([]swagger.ShippingDiscount, error)
}

type DiscountUsecase interface {
	GetByStoreID(ctx context.Context, id int64) ([]swagger.DiscountInfo, error)
}
