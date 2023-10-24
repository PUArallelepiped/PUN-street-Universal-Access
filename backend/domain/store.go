package domain

import (
	"context"
)

type Store struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

type StoreRepo interface {
	GetByID(ctx context.Context, id string) (*Store, error)
	GetAll(ctx context.Context) ([]Store, error)
}

type StoreUsecase interface {
	GetByID(ctx context.Context, id string) (*Store, error)
	GetAll(ctx context.Context) ([]Store, error)
}
