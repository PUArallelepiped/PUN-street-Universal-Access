package domain

import "context"

type HelloMsg struct {
	Name string `json:"name"`
}

type HelloRepo interface {
	SayHello(ctx context.Context) (*HelloMsg, error)
}

type HelloUsecase interface {
	SayHello(ctx context.Context) (*HelloMsg, error)
}
