package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type ProductRepo interface {
	GetByID(ctx context.Context, id int64) (*[]swagger.ProductInfo, error)
	PostByStoreId(ctx context.Context, id int64, product *swagger.ProductInfo) error
	PutById(ctx context.Context, storeId int64, productId int64, product *swagger.ProductInfo) error
}

type ProductUsecase interface {
	GetByID(ctx context.Context, id int64) (*[]swagger.ProductInfo, error)
	PostByStoreId(ctx context.Context, id int64, product *swagger.ProductInfo) error
	PutById(ctx context.Context, storeId int64, productId int64, product *swagger.ProductInfo) error
}
