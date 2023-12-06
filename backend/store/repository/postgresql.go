package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"

	"github.com/sirupsen/logrus"
)

type postgresqlStoreRepo struct {
	db *sql.DB
}

func NewPostgressqlStoreRepo(db *sql.DB) domain.StoreRepo {
	return &postgresqlStoreRepo{db}
}

func (p *postgresqlStoreRepo) GetByID(ctx context.Context, id int64) (*swagger.StoreInfoWithCategory, error) {
	sqlStatement := `
	 SELECT store_id, name, rate, rate_count, address, picture, description, shipping_fee, status, 
    (SELECT 
        jsonb_agg(jsonb_build_object('category_id', categories.category_id,'category_name', categories.name))
        as categories_item FROM categories NATURAL JOIN 
        (SELECT labels.category_id FROM labels WHERE labels.store_id = $1)
    ) AS category_array
	FROM stores
	WHERE store_id = $1;
	`

	row := p.db.QueryRow(sqlStatement, id)
	s := &swagger.StoreInfoWithCategory{}
	var categoryString string
	if err := row.Scan(&s.StoreId, &s.Name, &s.Rate, &s.RateCount, &s.Address, &s.Picture, &s.Description, &s.ShippingFee, &s.Status,
		&categoryString); err != nil {
		logrus.Error(err)
		return nil, err
	}

	var categoryArray []swagger.Category
	err := json.Unmarshal([]byte(categoryString), &categoryArray)
	if err != nil {
		logrus.Error(err)
	}
	s.CategoryArray = categoryArray

	return s, nil
}

func (p *postgresqlStoreRepo) GetAllStore(ctx context.Context) ([]swagger.StoreInfo, error) {
	rows, err := p.db.Query("SELECT store_id, name, rate, rate_count, address, picture, description, shipping_fee, status FROM stores")

	if err != nil {
		logrus.Error(err)
	}
	l := []swagger.StoreInfo{}
	for rows.Next() {
		s := swagger.StoreInfo{}
		err := rows.Scan(&s.StoreId, &s.Name, &s.Rate, &s.RateCount, &s.Address, &s.Picture, &s.Description, &s.ShippingFee, &s.Status)
		if err != nil {
			logrus.Error(err)
		}
		l = append(l, s)
	}
	return l, nil
}

func (p *postgresqlStoreRepo) GetMonthTotalPriceById(ctx context.Context, id int64, year int64, month int64) (*swagger.InlineResponse200, error) {
	sqlStatement := `
	SELECT COALESCE(SUM(total_price), 0)
	FROM orders 
	WHERE store_id = $1 AND 
	EXTRACT(year FROM order_date) = $2 AND 
	EXTRACT(month FROM order_date) = $3;
	`

	row := p.db.QueryRow(sqlStatement, id, year, month)

	price := &swagger.InlineResponse200{}
	if err := row.Scan(&price.Price); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return price, nil

}

func (p *postgresqlStoreRepo) GetAllProductSellingById(ctx context.Context, id int64, year int64, month int64) (*[]swagger.ProductStatistic, error) {
	sqlStatement := `
	SELECT carts.product_id, products.name, SUM(carts.product_quantity) 
	FROM orders, carts, products 
	WHERE orders.store_id = $1 AND 
	orders.cart_id = carts.cart_id AND customer_id = user_id AND carts.product_id = products.product_id AND 
	EXTRACT(year FROM order_date) = $2 AND EXTRACT(month FROM order_date) = $3
	GROUP BY carts.product_id, products.name
	`
	rows, err := p.db.Query(sqlStatement, id, year, month)
	if err != nil {
		logrus.Error(err)
	}

	productSellingArray := &[]swagger.ProductStatistic{}
	for rows.Next() {
		productSelling := &swagger.ProductStatistic{}
		err := rows.Scan(&productSelling.ProductId, &productSelling.ProductName, &productSelling.ProductQuantity)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		*productSellingArray = append(*productSellingArray, *productSelling)

	}
	return productSellingArray, nil

}
