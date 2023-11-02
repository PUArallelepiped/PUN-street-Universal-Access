package usecase

import (
	"context"
	"fmt"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/sirupsen/logrus"
)

type discountUsecase struct {
	discountRepo domain.DiscountRepo
}

func NewDiscountUsecase(discountRepo domain.DiscountRepo) domain.DiscountUsecase {
	return &discountUsecase{
		discountRepo: discountRepo,
	}
}

func (du *discountUsecase) GetByStoreID(ctx context.Context, id int64) ([]swagger.DiscountInfo, error) {
	shipping_discounts, err := du.discountRepo.GetShippingByStoreID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	discounts := []swagger.DiscountInfo{}
	for _, shipping_discount := range shipping_discounts {
		discount, err := du.discountRepo.GetByDiscountID(ctx, shipping_discount.DiscountId)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		fmt.Println(shipping_discount, discount)
		discounts = append(discounts, *discount)
	}

	return discounts, nil
}

func (du *discountUsecase) AddSeasoning(ctx context.Context, seasoning *swagger.SeasoningDiscount) error {
	err := du.discountRepo.AddSeasoning(ctx, seasoning)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
