package domain

import "context"

type UserRepo interface {
	Login(ctx context.Context, email string, password string) (int, error)
}

type UserUsecase interface {
	Login(ctx context.Context, email string, password string) (string, error)
}
