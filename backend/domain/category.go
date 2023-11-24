package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type CategoryRepo interface {
	GetAllCategory(ctx context.Context) (*[]swagger.Category, error)
}

type CategoryUsecase interface {
	GetAllCategory(ctx context.Context) (*[]swagger.Category, error)
}
