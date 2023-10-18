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

func (p *postgresqlProductRepo) GetByID(ctx context.Context, id string) (*swagger.ProductInfo, error) {
	row := p.db.QueryRow("SELECT product_id, name, catogory_id, describe, price, stock FROM product WHERE product_id = $1", id)
	product := &swagger.ProductInfo{}
	if err := row.Scan(&product.ProductId, &product.Name, &product.CatogoryId, &product.Description, &product.Price, &product.Storage); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return product, nil
}
