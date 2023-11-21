package usecase

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
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

func (pu *productUsecase) GetByID(ctx context.Context, id int64) (*[]swagger.ProductInfo, error) {
	products, err := pu.productRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return products, nil
}

func (pu *productUsecase) AddByStoreId(ctx context.Context, id int64, product *swagger.ProductInfo) error {
	product.Picture = UploadImg(product.Picture)
	err := pu.productRepo.AddByStoreId(ctx, id, product)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (pu *productUsecase) UpdateById(ctx context.Context, StoreId int64, productId int64, product *swagger.ProductInfo) error {
	err := pu.productRepo.UpdateById(ctx, StoreId, productId, product)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func UploadImg(byteImg string) string {
	httpposturl := "https://freeimage.host/api/1/upload"

	s := domain.PostInfo{
		Key:    "6d207e02198a847aa98d0a2a901485a5",
		Action: "upload",
		Source: byteImg,
		Format: "json"}
	v := make(map[string][]string)
	v["key"] = []string{s.Key}
	v["action"] = []string{s.Action}
	v["source"] = []string{s.Source}
	v["format"] = []string{s.Format}

	qs := url.Values(v)
	response, err := http.PostForm(httpposturl, qs)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	body, _ := io.ReadAll(response.Body)

	var responseInfo domain.ResponseInfo
	err = json.Unmarshal([]byte(string(body)), &responseInfo)
	if err != nil {
		logrus.Error(err)
		return ""
	}

	return responseInfo.Image.Url
}
