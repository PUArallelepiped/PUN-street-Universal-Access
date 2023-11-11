package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type DiscountRepo interface {
	GetShippingByStoreID(ctx context.Context, id int64) ([]swagger.ShippingDiscount, error)
	AddSeasoning(ctx context.Context, seasoning *swagger.SeasoningDiscount) error
	AddShipping(ctx context.Context, shipping *swagger.ShippingDiscount, id int64) error
	AddEvent(ctx context.Context, event *swagger.EventDiscount) error
}

type DiscountUsecase interface {
	GetShippingByStoreID(ctx context.Context, id int64) ([]swagger.ShippingDiscount, error)
	AddSeasoning(ctx context.Context, seasoning *swagger.SeasoningDiscount) error
	AddShipping(ctx context.Context, shipping *swagger.ShippingDiscount, id int64) error
	AddEvent(ctx context.Context, event *swagger.EventDiscount) error
}
