package domain

import "context"

type ProductRepo interface {
	GetByID(ctx context.Context, id string)
}

type ProductUsecase interface {
	GetByID(ctx context.Context, id string)
}
