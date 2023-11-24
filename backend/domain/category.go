package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type CategoryRepo interface {
	GetAllCategory(ctx context.Context) (*[]swagger.Category, error)
	AddCategoryToStore(ctx context.Context, store_id int64, category_id int64) error
	DeleteCategoryToStore(ctx context.Context, store_id int64, category_id int64) error
}

type CategoryUsecase interface {
	GetAllCategory(ctx context.Context) (*[]swagger.Category, error)
	AddCategoryToStore(ctx context.Context, store_id int64, category_id int64) error
	DeleteCategoryToStore(ctx context.Context, store_id int64, category_id int64) error
}
