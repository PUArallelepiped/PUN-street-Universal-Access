package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type UserRepo interface {
	GetByID(ctx context.Context, id int64) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserDataShort, error)
	GetAllOrder(ctx context.Context) ([]swagger.OrderInfoShort, error)
	BanUser(ctx context.Context, id int64) error
	UnBanUser(ctx context.Context, id int64) error
}

type UserUsecase interface {
	GetByID(ctx context.Context, id int64) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserDataShort, error)
	GetAllOrder(ctx context.Context) ([]swagger.OrderInfoShort, error)
	BanUser(ctx context.Context, id int64) error
	UnBanUser(ctx context.Context, id int64) error
}
