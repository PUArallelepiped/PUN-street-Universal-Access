package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type StoreRepo interface {
	GetByID(ctx context.Context, id int64) (*swagger.StoreInfoWithCategory, error)
	GetAllStore(ctx context.Context) ([]swagger.StoreInfoWithCategory, error)
	GetMonthTotalPriceById(ctx context.Context, id int64, year int64, month int64) (*swagger.InlineResponse200, error)
	GetAllProductSellingById(ctx context.Context, id int64, year int64, month int64) (*[]swagger.ProductStatistic, error)
	CalculateRate(ctx context.Context, id int64, rate swagger.RateInfo) error
}

type StoreUsecase interface {
	GetByID(ctx context.Context, id int64) (*swagger.StoreInfoWithCategory, error)
	GetAllStore(ctx context.Context) ([]swagger.StoreInfoWithCategory, error)
	GetStatisticsById(ctx context.Context, id int64, year int64) (*[]swagger.InlineResponse200, error)
	GetAllProductSellingById(ctx context.Context, id int64, year int64, month int64) (*[]swagger.ProductStatistic, error)
	CalculateRate(ctx context.Context, id int64, rate swagger.RateInfo) error
}
