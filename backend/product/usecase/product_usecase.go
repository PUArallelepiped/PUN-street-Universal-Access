package usecase

import (
	"context"
	"fmt"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
)

type productUsecase struct {
	productRepo domain.ProductRepo
}

func NewProductUsecase(productRepo domain.ProductRepo) domain.ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (pu *productUsecase) GetProductByID(ctx context.Context, id int64) (*swagger.ProductInfoWithLabelAndDiscount, error) {
	product, err := pu.productRepo.GetProductByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return product, nil

}

func (pu *productUsecase) GetProductsByStoreID(ctx context.Context, id int64) (*[]swagger.ProductInfoWithLabelAndDiscount, error) {
	products, err := pu.productRepo.GetProductsByStoreID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return products, nil
}

func (pu *productUsecase) AddByStoreId(ctx context.Context, id int64, product *swagger.ProductInfoWithLabelAndDiscount) error {
	//update product
	if product.ProductId != 0 {
		productById, err := pu.productRepo.GetProductByID(ctx, product.ProductId)
		if err != nil {
			logrus.Error(err)
			return err
		}
		productById.Status = product.Status

		if cmp.Equal(productById, product) {
			//just change status to post body status
			fmt.Println("same")
			err := pu.productRepo.ChangeProductStatusByProductID(ctx, product.ProductId, product.Status)
			if err != nil {
				logrus.Error(err)
				return err
			}
		} else {
			//change status = 0 and add product
			fmt.Println("not same")
			err := pu.productRepo.ChangeProductStatusByProductID(ctx, product.ProductId, 0)
			if err != nil {
				logrus.Error(err)
				return err
			}
			// add product
			err = pu.AddProductDiscountLabel(ctx, id, product)
			if err != nil {
				logrus.Error(err)
				return err
			}
		}
	} else {
		//just add new product
		err := pu.AddProductDiscountLabel(ctx, id, product)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	return nil
}

func (pu *productUsecase) AddProductDiscountLabel(ctx context.Context, id int64, product *swagger.ProductInfoWithLabelAndDiscount) error {
	// add product
	fmt.Println("add product")
	productID, err := pu.productRepo.AddProductByStoreID(ctx, id, product)
	if err != nil {
		logrus.Error(err)
		return err
	}
	// add event discount array
	if product.EventDiscountArray != nil {
		fmt.Println("add discount")
		eventDiscountArray := product.EventDiscountArray
		for _, eventDiscount := range eventDiscountArray {
			eventDiscount.ProductId = productID
			err = pu.productRepo.AddDiscountByProductID(ctx, &eventDiscount)
			if err != nil {
				logrus.Error(err)
				return err
			}
		}
	}
	//add label array
	if product.ProductLabelArray != nil {
		productLabelArray := product.ProductLabelArray
		for _, productLabel := range productLabelArray {
			fmt.Println("add label")
			//add product label
			productLabel.ProductId = productID
			err = pu.productRepo.AddProductLabel(ctx, productLabel.ProductId, productLabel.LabelName, productLabel.Required)
			if err != nil {
				logrus.Error(err)
				return err
			}
			fmt.Println("add label item")
			//add product label item
			productLabelItemArray := productLabel.ItemArray
			for _, productLabelItem := range productLabelItemArray {
				fmt.Println(productLabelItemArray)
				fmt.Println(productLabelItem)
				err = pu.productRepo.AddProductLabelItem(ctx, productLabel.ProductId, productLabel.LabelName, productLabelItem.Name)
				if err != nil {
					logrus.Error(err)
					return err
				}
			}
		}
	}

	return nil
}
