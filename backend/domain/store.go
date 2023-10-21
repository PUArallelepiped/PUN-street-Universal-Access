package domain

import "context"

type Store struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

type StoreRepo interface {
	GetByID(ctx context.Context, id string) (*Store, error)
}

type StoreUsecase interface {
	GetByID(ctx context.Context, id string) (*Store, error)
}
