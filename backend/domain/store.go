package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type StoreRepo interface {
	GetByID(ctx context.Context, id string) (*swagger.StoreInfo, error)
	GetAllStore(ctx context.Context) ([]swagger.StoreInfo, error)
	GetMonthTotalPriceById(ctx context.Context, id int64, year int64, month int64) (*swagger.InlineResponse200, error)
}

type StoreUsecase interface {
	GetByID(ctx context.Context, id string) (*swagger.StoreInfo, error)
	GetAllStore(ctx context.Context) ([]swagger.StoreInfo, error)
	GetStatisticsById(ctx context.Context, id int64, year int64) (*[]swagger.InlineResponse200, error)
}
