package usecase

import (
	"context"

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

func (du *discountUsecase) GetShippingByStoreID(ctx context.Context, id int64) (*swagger.ShippingDiscount, error) {
	discount, err := du.discountRepo.GetShippingByStoreID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return discount, nil
}

func (du *discountUsecase) AddSeasoning(ctx context.Context, seasoning *swagger.SeasoningDiscount) error {
	err := du.discountRepo.AddSeasoning(ctx, seasoning)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (du *discountUsecase) AddShipping(ctx context.Context, shipping *swagger.ShippingDiscount, id int64) error {
	exist, err := du.discountRepo.IsExistShippingDiscountByStoreID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// check exist then for each shipping discount change status = 0
	if exist {
		discount, err := du.discountRepo.GetShippingByStoreID(ctx, id)
		if err != nil {
			logrus.Error(err)
			return err
		}
		err = du.discountRepo.DisableDiscountByDiscountID(ctx, discount.DiscountId)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	// insert new discounts and shipping_discount table
	err = du.discountRepo.AddShipping(ctx, shipping, id)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (du *discountUsecase) AddEvent(ctx context.Context, event *swagger.EventDiscount, id int64) error {
	err := du.discountRepo.AddEvent(ctx, event, id)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (du *discountUsecase) GetAllSeasoning(ctx context.Context) ([]swagger.SeasoningDiscount, error) {
	discounts, err := du.discountRepo.GetAllSeasoning(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return discounts, nil
}

func (du *discountUsecase) GetAllEventByProductID(ctx context.Context, id int64) ([]swagger.EventDiscount, error) {
	discounts, err := du.discountRepo.GetAllEventByProductID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return discounts, nil
}

func (du *discountUsecase) DisableDiscountByDiscountID(ctx context.Context, id int64) error {
	err := du.discountRepo.DisableDiscountByDiscountID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
