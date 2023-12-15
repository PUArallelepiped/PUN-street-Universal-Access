package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type ProductRepo interface {
	GetProductByID(ctx context.Context, id int64) (*swagger.ProductInfoWithLabelAndDiscount, error)
	GetProductsByStoreID(ctx context.Context, id int64) (*[]swagger.ProductInfoWithLabelAndDiscount, error)
	ChangeProductStatusByProductID(ctx context.Context, id int64, status int64) error
	AddProductByStoreID(ctx context.Context, id int64, product *swagger.ProductInfoWithLabelAndDiscount) (int64, error)
	AddDiscountByProductID(ctx context.Context, event *swagger.EventDiscount) error
	AddProductLabel(ctx context.Context, productId int64, label_name string, required bool) error
	AddProductLabelItem(ctx context.Context, productId int64, labelName string, name string) error
}

type ProductUsecase interface {
	GetProductByID(ctx context.Context, id int64) (*swagger.ProductInfoWithLabelAndDiscount, error)
	GetProductsByStoreID(ctx context.Context, id int64) (*[]swagger.ProductInfoWithLabelAndDiscount, error)
	AddByStoreId(ctx context.Context, id int64, product *swagger.ProductInfoWithLabelAndDiscount) error
	AddProductDiscountLabel(ctx context.Context, id int64, product *swagger.ProductInfoWithLabelAndDiscount) error
	DeleteProduct(ctx context.Context, id int64) error
}
