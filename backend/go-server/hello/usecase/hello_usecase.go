package usecase

import (
	"context"
	"main/domain"

	"github.com/sirupsen/logrus"
)

type helloUsecase struct {
	helloRepo domain.HelloRepo
}

func NewHelloUsecase(helloRepo domain.HelloRepo) domain.HelloUsecase {
	return &helloUsecase{
		helloRepo: helloRepo,
	}
}

func (hu *helloUsecase) SayHello(ctx context.Context) (*domain.HelloMsg, error) {
	name, err := hu.helloRepo.SayHello(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return name, nil
}
