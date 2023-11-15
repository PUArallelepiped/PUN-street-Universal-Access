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

func (p *postgresqlProductRepo) GetByProductID(ctx context.Context, id int64) (*swagger.ProductInfo, error) {
	row := p.db.QueryRow("SELECT product_id, store_id, name, description, picture, price, stock, status FROM products WHERE product_id = $1", id)

	product := &swagger.ProductInfo{}
	if err := row.Scan(&product.ProductId, &product.StoreId, &product.Name, &product.Description, &product.Picture, &product.Price, &product.Stock, &product.Status); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return product, nil
  
func (p *postgresqlProductRepo) AddByStoreId(ctx context.Context, id int64, product *swagger.ProductInfo) error {
	sqlStatement := `
	INSERT INTO products (store_id, name, description, picture, price, stock, status) VALUES
    ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := p.db.Exec(sqlStatement, id, product.Name, product.Description, product.Picture, product.Price, product.Stock, product.Status)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (p *postgresqlProductRepo) UpdateById(ctx context.Context, storeId int64, productId int64, product *swagger.ProductInfo) error {
	sqlStatement := `
	UPDATE products SET 
	name = $1, description = $2, picture = $3, price = $4, stock = $5, status = $6
	WHERE store_id = $7 AND product_id = $8 
	`

	_, err := p.db.Exec(sqlStatement, product.Name, product.Description, product.Picture, product.Price, product.Stock, product.Status, storeId, productId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
