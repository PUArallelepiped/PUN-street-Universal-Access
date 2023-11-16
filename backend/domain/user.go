package domain

import "context"

type UserRepo interface {
	Login(ctx context.Context, email string, password string) error
}

type UserUsecase interface {
	Login(ctx context.Context, email string, password string) error
}
