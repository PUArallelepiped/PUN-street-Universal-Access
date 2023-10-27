package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type CartRepo interface {
	PostCart(ctx context.Context, cart *swagger.CartInfo) error
}

type CartUsecase interface {
	PostCart(ctx context.Context, cart *swagger.CartInfo) error
}
