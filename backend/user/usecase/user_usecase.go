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

func (su *UserUsecase) GetAllUser(ctx context.Context) ([]swagger.UserDataShort, error) {
	s, err := su.userRepo.GetAllUser(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}

func (su *UserUsecase) GetAllOrder(ctx context.Context) ([]swagger.OrderInfoShort, error) {
	orders, err := su.userRepo.GetAllOrder(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return orders, nil
}

func (su *UserUsecase) BanUser(ctx context.Context, id int64) error {
	err := su.userRepo.BanUser(ctx, id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (su *UserUsecase) UnBanUser(ctx context.Context, id int64) error {
	err := su.userRepo.UnBanUser(ctx, id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
