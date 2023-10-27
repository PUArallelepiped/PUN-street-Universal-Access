package repository

import (
	"context"
	"database/sql"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/sirupsen/logrus"
)

type postgresqlProductRepo struct {
	db *sql.DB
}

func NewPostgressqlProductRepo(db *sql.DB) domain.ProductRepo {
	return &postgresqlProductRepo{db}
}

func (p *postgresqlProductRepo) GetByID(ctx context.Context, id int64) (*[]swagger.ProductInfo, error) {
	row, err := p.db.Query("SELECT product_id, store_id, name, description, picture, price, stock, status FROM products WHERE store_id = $1", id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	products := &[]swagger.ProductInfo{}
	for row.Next() {
		product := &swagger.ProductInfo{}
		err = row.Scan(&product.ProductId, &product.StoreId, &product.Name, &product.Description, &product.Picture, &product.Price, &product.Stock, &product.Status)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		*products = append(*products, *product)
	}
	return products, nil
}
