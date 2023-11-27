package usecase

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"

	"github.com/sirupsen/logrus"
)

type storeUsecase struct {
	storeRepo domain.StoreRepo
}

func NewStoreUsecase(storeRepo domain.StoreRepo) domain.StoreUsecase {
	return &storeUsecase{
		storeRepo: storeRepo,
	}
}

func (su *storeUsecase) GetByID(ctx context.Context, id string) (*swagger.StoreInfo, error) {
	s, err := su.storeRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}

func (su *storeUsecase) GetAllStore(ctx context.Context) ([]swagger.StoreInfo, error) {
	s, err := su.storeRepo.GetAllStore(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}

func (su *storeUsecase) GetStatisticsById(ctx context.Context, id int64, year int64) (*[]swagger.InlineResponse200, error) {
	var priceArray []swagger.InlineResponse200
	for month := 1; month < 13; month++ {
		price, err := su.storeRepo.GetMonthTotalPriceById(ctx, id, year, int64(month))
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		priceArray = append(priceArray, *price)
	}

	return &priceArray, nil
}

func (su *storeUsecase) GetAllProductSellingById(ctx context.Context, id int64, year int64, month int64) (*[]swagger.ProductStatistic, error) {
	productStatistics, err := su.storeRepo.GetAllProductSellingById(ctx, id, year, month)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return productStatistics, nil
}
