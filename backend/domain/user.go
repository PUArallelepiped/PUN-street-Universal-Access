package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type UserRepo interface {
<<<<<<< HEAD
	GetByID(ctx context.Context, id string) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserData, error)
	LogIn(ctx context.Context, userName string, userPassword string) (bool, error)
}

type UserUsecase interface {
	GetByID(ctx context.Context, id string) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserData, error)
	LogIn(ctx context.Context, userName string, userPassword string) (bool, error)
=======
	GetByID(ctx context.Context, id int64) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserDataShort, error)
}

type UserUsecase interface {
	GetByID(ctx context.Context, id int64) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserDataShort, error)
>>>>>>> upstream/main
}
