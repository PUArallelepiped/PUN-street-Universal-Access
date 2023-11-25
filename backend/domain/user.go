package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type UserRepo interface {
	GetByID(ctx context.Context, id string) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserData, error)
}

type UserUsecase interface {
	GetByID(ctx context.Context, id string) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserData, error)
}
