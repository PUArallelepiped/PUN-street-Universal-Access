package usecase

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/google/go-cmp/cmp"

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

func (su *storeUsecase) GetByID(ctx context.Context, id int64) (*swagger.StoreInfoWithCategory, error) {
	s, err := su.storeRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}

func (su *storeUsecase) GetAllStore(ctx context.Context, searchInfo *swagger.SearchInfo) (*[]swagger.OneStoreListInfo, error) {
	//search by price
	store, err := su.storeRepo.GetAllStoreByPrice(ctx, searchInfo)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	//if user search by category
	if len(searchInfo.CategoryArray) != 0 {
		for _, category := range searchInfo.CategoryArray {
			storeByCategory, err := su.storeRepo.GetAllStoreByCategoryId(ctx, category.CategoryId)
			if err != nil {
				logrus.Error(err)
				return nil, err
			}
			*store = findIntersection(*store, *storeByCategory)
		}
	}
	//if user search by string
	if searchInfo.SearchString != "" {
		storeByString, err := su.storeRepo.GetAllStoreByString(ctx, searchInfo)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		*store = findIntersection(*store, *storeByString)
	}

	return store, nil
}

func findIntersection(arr1 []swagger.OneStoreListInfo, arr2 []swagger.OneStoreListInfo) []swagger.OneStoreListInfo {
	stores := []swagger.OneStoreListInfo{}

	for _, num1 := range arr1 {
		// Check if num1 exists in arr2
		for _, num2 := range arr2 {
			if cmp.Equal(num1, num2) {
				stores = append(stores, num1)
				break
			}
		}
	}

	return stores
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

func (su *storeUsecase) CalculateRate(ctx context.Context, id int64, rate swagger.RateInfo) error {
	return su.storeRepo.CalculateRate(ctx, id, rate)
}
