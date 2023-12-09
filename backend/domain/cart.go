package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type CartRepo interface {
	GetAllHistoryById(ctx context.Context, id int64) (*[]swagger.HistoryInfo, error)
	GetRunOrderByID(ctx context.Context, id int64) (*[]swagger.RunOrderInfo, error)
}

type CartUsecase interface {
	GetAllHistoryById(ctx context.Context, id int64) (*[]swagger.HistoryInfo, error)
	GetRunOrderByID(ctx context.Context, id int64) (*[]swagger.RunOrderInfo, error)
}
