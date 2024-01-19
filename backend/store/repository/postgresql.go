package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"math"

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
        jsonb_agg(jsonb_build_object('category_id', categories.category_id,'category_name', categories.name)) AS categories_item 
		FROM categories NATURAL JOIN 
        (SELECT labels.category_id FROM labels WHERE labels.store_id = $1) AS label
    ) AS category_array
	FROM stores
	WHERE store_id = $1;
	`

	row := p.db.QueryRow(sqlStatement, id)
	s := &swagger.StoreInfoWithCategory{}
	var categoryString sql.NullString
	if err := row.Scan(&s.StoreId, &s.Name, &s.Rate, &s.RateCount, &s.Address, &s.Picture, &s.Description, &s.ShippingFee, &s.Status,
		&categoryString); err != nil {
		logrus.Error(err)
		return nil, err
	}

	if categoryString.Valid {
		var categoryArray []swagger.Category
		err := json.Unmarshal([]byte(categoryString.String), &categoryArray)
		if err != nil {
			logrus.Error(err)
		}
		s.CategoryArray = categoryArray
	}

	return s, nil
}

func (p *postgresqlStoreRepo) GetAllStoreByPrice(ctx context.Context, searchInfo *swagger.SearchInfo) (*[]swagger.OneStoreListInfo, error) {
	sqlStatement := `
	SELECT store_id, name, rate, picture, category_array FROM
	(SELECT store_id, name, rate, picture,
		(SELECT 
			jsonb_agg(jsonb_build_object(
			'category_id', categories.category_id,
			'category_name', categories.name))
			as categories_item FROM categories NATURAL JOIN 
			(SELECT labels.category_id FROM labels 
			WHERE labels.store_id = stores.store_id) AS labels
		) AS category_array,
		(SELECT COUNT(*) > 0 FROM products
		WHERE products.store_id = stores.store_id
		AND price <= $1 AND price >= $2) AS is_in_price_range
	FROM stores) AS store_info
	WHERE is_in_price_range = true
	`

	rows, err := p.db.Query(sqlStatement, searchInfo.PriceHigh, searchInfo.PriceLow)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	stores := &[]swagger.OneStoreListInfo{}
	var categoryString sql.NullString
	for rows.Next() {
		store := &swagger.OneStoreListInfo{}
		err := rows.Scan(&store.StoreId, &store.Name, &store.Rate, &store.Picture, &categoryString)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		if categoryString.Valid {
			var categoryArray []swagger.Category
			err := json.Unmarshal([]byte(categoryString.String), &categoryArray)
			if err != nil {
				logrus.Error(err)
			}
			store.CategoryArray = categoryArray
		}

		*stores = append(*stores, *store)
	}

	return stores, nil
}

func (p *postgresqlStoreRepo) GetAllStoreByCategoryId(ctx context.Context, id int64) (*[]swagger.OneStoreListInfo, error) {
	sqlStatement := `
	SELECT store_id, name, rate, picture, category_array FROM
	(SELECT store_id, name, rate, picture,
		(SELECT 
			jsonb_agg(jsonb_build_object(
			'category_id', categories.category_id,
			'category_name', categories.name))
			as categories_item FROM categories NATURAL JOIN 
			(SELECT labels.category_id FROM labels 
			WHERE labels.store_id = stores.store_id) AS labels
		) AS category_array,
		(SELECT COUNT(*)>0 FROM labels
			WHERE category_id = $1 AND labels.store_id = stores.store_id
		) AS is_category_exits
	FROM stores) AS store_info
	WHERE is_category_exits = true
	`
	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	stores := &[]swagger.OneStoreListInfo{}
	var categoryString sql.NullString
	for rows.Next() {
		store := &swagger.OneStoreListInfo{}
		err := rows.Scan(&store.StoreId, &store.Name, &store.Rate, &store.Picture, &categoryString)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		if categoryString.Valid {
			var categoryArray []swagger.Category
			err := json.Unmarshal([]byte(categoryString.String), &categoryArray)
			if err != nil {
				logrus.Error(err)
			}
			store.CategoryArray = categoryArray
		}

		*stores = append(*stores, *store)
	}

	return stores, nil
}

func (p *postgresqlStoreRepo) GetAllStoreByString(ctx context.Context, searchInfo *swagger.SearchInfo) (*[]swagger.OneStoreListInfo, error) {
	sqlStatement := `
	SELECT store_id, name, rate, picture, category_array FROM
		(SELECT store_id, name, rate, picture,
			(SELECT 
				jsonb_agg(jsonb_build_object(
				'category_id', categories.category_id,
				'category_name', categories.name))
				as categories_item FROM categories NATURAL JOIN 
				(SELECT labels.category_id FROM labels 
				WHERE labels.store_id = stores.store_id) AS labels
			) AS category_array,
			(SELECT COUNT(*) > 0
				FROM products LEFT JOIN stores as s ON products.store_id = stores.store_id
				WHERE products.store_id = stores.store_id AND(
				products.name LIKE '%' || $1 || '%' OR
				stores.name LIKE '%' || $1 || '%' OR
				products.description LIKE '%' || $1 || '%'
			)) AS is_string_exits
		FROM stores) as store
	WHERE is_string_exits = true
	`

	rows, err := p.db.Query(sqlStatement, searchInfo.SearchString)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	stores := &[]swagger.OneStoreListInfo{}
	var categoryString sql.NullString
	for rows.Next() {
		store := &swagger.OneStoreListInfo{}
		err := rows.Scan(&store.StoreId, &store.Name, &store.Rate, &store.Picture, &categoryString)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		if categoryString.Valid {
			var categoryArray []swagger.Category
			err := json.Unmarshal([]byte(categoryString.String), &categoryArray)
			if err != nil {
				logrus.Error(err)
			}
			store.CategoryArray = categoryArray
		}

		*stores = append(*stores, *store)
	}

	return stores, nil
}

func (p *postgresqlStoreRepo) GetMonthTotalPriceById(ctx context.Context, id int64, year int64, month int64) (*swagger.InlineResponse200, error) {
	sqlStatement := `
	SELECT COALESCE(SUM(total_price), 0)
	FROM orders 
	WHERE store_id = $1 AND orders.status >= 6 AND 
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
	FROM carts 
	LEFT JOIN products ON carts.product_id = products.product_id 
	LEFT JOIN orders ON carts.store_id = orders.store_id AND carts.cart_id = orders.cart_id AND carts.customer_id = orders.user_id
	WHERE orders.store_id = $1 AND orders.status >= 6  AND
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

func (p *postgresqlStoreRepo) CalculateRate(ctx context.Context, id int64, rate swagger.RateInfo) error {
	sqlStatement := `
	UPDATE stores
	SET rate = ROUND((rate*rate_count+$1)::numeric/(rate_count+1),1) , rate_count = rate_count+1
	WHERE store_id=$2
	`
	tempRate := float32(math.Max(0, math.Min(5, float64(rate.Rate))))
	if _, err := p.db.Exec(sqlStatement, tempRate, id); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (p *postgresqlStoreRepo) UpdateOrderRate(ctx context.Context, customerId int64, cartID int64, storeID int64, rate swagger.RateInfo) error {
	sqlStatement := `
	UPDATE orders set rate = $1
	WHERE user_id =$2 AND cart_id =$3 AND store_id = $4 AND rate = 0
	`

	_, err := p.db.Exec(sqlStatement, rate.Rate, customerId, cartID, storeID)
	if err != nil {
		return err
	}

	return nil
}
