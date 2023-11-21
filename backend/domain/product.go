package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type PostInfo struct {
	Key    string `json:"key"`
	Action string `json:"action"`
	Source string `json:"source"`
	Format string `json:"format"`
}

type ResponseInfo struct {
	StatusCode int       `json:"status_code"`
	Image      ImageInfo `json:"image"`
}

type ImageInfo struct {
	Url       string `json:"url"`
	UrlViewer string `json:"url_viewer"`
}

type ProductRepo interface {
	GetByID(ctx context.Context, id int64) (*[]swagger.ProductInfo, error)
	GetByProductID(ctx context.Context, id int64) (*swagger.ProductInfo, error)
	AddByStoreId(ctx context.Context, id int64, product *swagger.ProductInfo) error
	UpdateById(ctx context.Context, storeId int64, productId int64, product *swagger.ProductInfo) error
}

type ProductUsecase interface {
	GetByID(ctx context.Context, id int64) (*[]swagger.ProductInfo, error)
	AddByStoreId(ctx context.Context, id int64, product *swagger.ProductInfo) error
	UpdateById(ctx context.Context, storeId int64, productId int64, product *swagger.ProductInfo) error
}
