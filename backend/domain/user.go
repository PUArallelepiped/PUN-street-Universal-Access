package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type UserRepo interface {
	GetByID(ctx context.Context, id string) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserData, error)
	LogIn(ctx context.Context, userName string, userPassword string) (bool, error)
}

type UserUsecase interface {
	GetByID(ctx context.Context, id string) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserData, error)
	LogIn(ctx context.Context, userName string, userPassword string) (bool, error)
}
