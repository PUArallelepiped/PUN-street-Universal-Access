package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/sirupsen/logrus"
)

type postgresqlCartRepo struct {
	db *sql.DB
}

func NewPostgressqlCartRepo(db *sql.DB) domain.CartRepo {
	return &postgresqlCartRepo{db}
}

func (p *postgresqlCartRepo) GetAllHistoryById(ctx context.Context, id int64) (*[]swagger.HistoryInfo, error) {
	sqlStatement := `
	SELECT orders.store_id, orders.cart_id, orders.order_date, orders.total_price, orders.user_id AS customer_id, orders.status, 
	stores.name AS store_name, stores.picture AS store_picture, stores.rate AS store_rate
	FROM orders LEFT JOIN stores ON orders.store_id = stores.store_id WHERE orders.user_id = $1 AND orders.status = 6;
	`

	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	historyArray := &[]swagger.HistoryInfo{}
	for rows.Next() {
		history := &swagger.HistoryInfo{}
		err := rows.Scan(&history.StoreId, &history.CartId, &history.OrderDate,
			&history.TotalPrice, &history.CustomerId, &history.Status,
			&history.StoreName, &history.StorePicture, &history.StoreRate)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		*historyArray = append(*historyArray, *history)
	}

	return historyArray, nil
}

func (p *postgresqlCartRepo) GetRunOrderByID(ctx context.Context, id int64) (*[]swagger.RunOrderInfo, error) {
	sqlStatement := `
	SELECT orders.store_id, orders.cart_id, orders.user_id, orders.status, 
	stores.name AS store_name, stores.picture AS store_picture
	FROM orders LEFT JOIN stores ON orders.store_id = stores.store_id 
	WHERE orders.user_id = $1 AND 
	orders.status != 0 AND orders.status != 1 AND orders.status != 6;
	`

	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	runOrderArray := &[]swagger.RunOrderInfo{}
	for rows.Next() {
		runOrder := &swagger.RunOrderInfo{}
		err := rows.Scan(&runOrder.StoreId, &runOrder.CartId, &runOrder.UserId, &runOrder.Status,
			&runOrder.StoreName, &runOrder.StorePicture)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		*runOrderArray = append(*runOrderArray, *runOrder)
	}

	return runOrderArray, nil
}

func (p *postgresqlCartRepo) DeleteOrder(ctx context.Context, customerId int64, storeId int64) error {
	sqlStatement := `
	DELETE FROM orders
	WHERE user_id = $1 AND 
	cart_id = ( SELECT current_cart_id FROM user_data WHERE user_id = $1) AND 
	store_id = $2 AND 
	status = 0;
	`
	_, err := p.db.Exec(sqlStatement, customerId, storeId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlCartRepo) IsExitsOrderByStoreCartId(ctx context.Context, customerId int64, storeId int64) (bool, error) {
	sqlStatement := `
	SELECT COUNT(*) > 0 FROM orders 
	WHERE user_id = $1 AND 
	cart_id = ( SELECT current_cart_id FROM user_data WHERE user_id = $1) AND 
	store_id = $2;
	`
	row := p.db.QueryRow(sqlStatement, customerId, storeId)

	exist := false
	err := row.Scan(&exist)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	return exist, nil
}

func (p *postgresqlCartRepo) IsExitsCartByStoreCartId(ctx context.Context, customerId int64, storeId int64) (bool, error) {
	sqlStatement := `
	SELECT COUNT(*) > 0 FROM carts 
	WHERE customer_id = $1 AND 
	cart_id = ( SELECT current_cart_id FROM user_data WHERE user_id = $1) AND 
	store_id = $2;
	`
	row := p.db.QueryRow(sqlStatement, customerId, storeId)

	exist := false
	err := row.Scan(&exist)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	return exist, nil
}

func (p *postgresqlCartRepo) DeleteProduct(ctx context.Context, customerId int64, productId int64) (int64, error) {
	sqlStatement := `
	DELETE FROM carts
	WHERE customer_id = $1 AND 
	cart_id = ( SELECT current_cart_id FROM user_data WHERE user_id = $1) AND 
	product_id = $2
	RETURNING store_id
	`
	row := p.db.QueryRow(sqlStatement, customerId, productId)

	var storeId int64
	err := row.Scan(&storeId)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return storeId, nil
}

func (p *postgresqlCartRepo) AddProductToCart(ctx context.Context, customerId int64, cartInfo *swagger.CartInfo) error {
	sqlStatement := `
	INSERT INTO carts (customer_id, cart_id, store_id, product_id, product_quantity, event_discount_id) VALUES
	($1, 
	(SELECT current_cart_id FROM user_data WHERE user_id = $1),
	$2, $3, $4, $5)
	ON CONFLICT (customer_id, product_id, store_id, cart_id) DO UPDATE 
	SET product_quantity = EXCLUDED.product_quantity;
	`
	_, err := p.db.Exec(sqlStatement, customerId, cartInfo.StoreId, cartInfo.ProductId, cartInfo.ProductQuantity, cartInfo.DiscountId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlCartRepo) AddOrderByCartInfo(ctx context.Context, customerId int64, storeId int64) error {
	sqlStatement := `
	INSERT INTO orders (user_id, cart_id, store_id, seasoning_discount_id, 
		shipping_discount_id, status, total_price, Order_date, taking_address, taking_method) VALUES 
    ($1, 
	(SELECT current_cart_id FROM user_data WHERE user_id = $1),
	$2, 
	(SELECT COALESCE(
		(SELECT discount_id
		FROM seasoning_discount
		WHERE start_date <= $3 AND end_date >= $3),
		1
	) AS discount_id),
	(SELECT COALESCE(shipping_discount.discount_id, 1) FROM shipping_discount LEFT JOIN
	discounts ON shipping_discount.discount_id = discounts.discount_id
	WHERE shipping_discount.store_id = $2 AND discounts.status = 1),
	0, 
	0, 
	'2020-01-01 00:00:00', 
	(SELECT address FROM user_data WHERE user_id = $1), 
	1)
	`
	dt := time.Now().Format("01-02-2006 15:04:05")

	_, err := p.db.Exec(sqlStatement, customerId, storeId, dt)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlCartRepo) IsProductCanAdd(ctx context.Context, id int64) (bool, error) {
	sqlStatement := `
	SELECT (status != 2) FROM products WHERE product_id = $1
	`
	row := p.db.QueryRow(sqlStatement, id)

	canAdd := false
	err := row.Scan(&canAdd)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	return canAdd, nil
}
