package usecase

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
)

type userUsecase struct {
	userRepo domain.UserRepo
}

func NewUserUsecase(userRepo domain.UserRepo) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (uu *userUsecase) Login(ctx context.Context, email string, password string) error {
	err := uu.userRepo.Login(ctx, email, password)
	if err != nil {
		return err
	}

	return nil
}
