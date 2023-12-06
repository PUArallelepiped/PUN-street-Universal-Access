package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type ProductRepo interface {
	GetProductByID(ctx context.Context, id int64) (*swagger.ProductInfoWithLabelAndDiscount, error)
	GetProductsByStoreID(ctx context.Context, id int64) (*[]swagger.ProductInfoWithLabelAndDiscount, error)

	// AddByStoreId(ctx context.Context, id int64, product *swagger.ProductInfo) error
}

type ProductUsecase interface {
	GetProductByID(ctx context.Context, id int64) (*swagger.ProductInfoWithLabelAndDiscount, error)
	GetProductsByStoreID(ctx context.Context, id int64) (*[]swagger.ProductInfoWithLabelAndDiscount, error)
	// AddByStoreId(ctx context.Context, id int64, product *swagger.ProductInfo) error
}
