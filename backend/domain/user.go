package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type UserRepo interface {
	GetByID(ctx context.Context, id int64) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserData, error)
	Login(ctx context.Context, email string, password string) (int, error)
}

type UserUsecase interface {
	GetByID(ctx context.Context, id int64) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserData, error)
	Login(ctx context.Context, email string, password string) (string, error)
	ValidateToken(ctx context.Context, token string) error
}
