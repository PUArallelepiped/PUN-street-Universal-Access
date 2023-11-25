package usecase

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	userRepo domain.UserRepo
}

func NewUserUsecase(userRepo domain.UserRepo) domain.UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (su *UserUsecase) GetByID(ctx context.Context, id int64) (*swagger.UserData, error) {
	s, err := su.userRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}

func (su *UserUsecase) GetAllUser(ctx context.Context) ([]swagger.UserData, error) {
	s, err := su.userRepo.GetAllUser(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}
