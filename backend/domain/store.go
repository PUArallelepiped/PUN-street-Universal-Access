package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type StoreRepo interface {
	GetByID(ctx context.Context, id string) (*swagger.StoreInfo, error)
}

type StoreUsecase interface {
	GetByID(ctx context.Context, id string) (*swagger.StoreInfo, error)
}
