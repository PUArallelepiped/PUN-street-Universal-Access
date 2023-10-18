package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type ProductRepo interface {
	GetByID(ctx context.Context, id string) (*swagger.ProductInfo, error)
}

type ProductUsecase interface {
	GetByID(ctx context.Context, id string) (*swagger.ProductInfo, error)
}
