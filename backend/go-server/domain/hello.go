package domain

import "context"

type HelloRepo interface {
	SayHello(ctx context.Context) error
}

type HelloUsecase interface {
	SayHello(ctx context.Context) error
}
