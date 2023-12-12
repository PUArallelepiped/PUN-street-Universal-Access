package repository

import (
	"context"
	"database/sql"
	"encoding/json"
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
	($1, $2, $3, $4, $5, $6)
	ON CONFLICT (customer_id, product_id, store_id, cart_id) DO UPDATE 
	SET product_quantity = EXCLUDED.product_quantity;
	`
	_, err := p.db.Exec(sqlStatement, customerId, cartInfo.CartId, cartInfo.StoreId, cartInfo.ProductId, cartInfo.ProductQuantity, cartInfo.DiscountId)
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
	$2, 1, 1, 0, 0, '2020-01-01 00:00:00', '台北市', 1)
	`
	_, err := p.db.Exec(sqlStatement, customerId, storeId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlCartRepo) GetHistoryCart(ctx context.Context, customerId int64, cartId int64, storeId int64) (*swagger.StoreOrderInfo, error) {
	sqlStatement := `
		SELECT store_id, name AS store_name, shipping_fee AS store_shipping_fee,
			(SELECT jsonb_build_object(
				'discount_id', COALESCE(discounts.discount_id, 1),
				'discount_name', COALESCE(discounts.name, 'default'), 
				'discount_description', COALESCE(discounts.description, 'no discount'),
				'status', COALESCE(discounts.status, 1), 
				'discount_max_price', shipping_discount.max_price)
				AS shipping_discount 
					FROM orders LEFT JOIN 
						(discounts LEFT JOIN 
							shipping_discount ON discounts.discount_id = shipping_discount.discount_id) 
						ON orders.shipping_discount_id = discounts.discount_id
				WHERE orders.user_id = $1 AND orders.cart_id = $2 AND orders.store_id = stores.store_id),
			(SELECT jsonb_build_object(
				'discount_id', COALESCE(discounts.discount_id, 1),
				'discount_name', COALESCE(discounts.name, 'default'), 
				'discount_description', COALESCE(discounts.description, 'no discount'),
				'status', COALESCE(discounts.status, 1), 
				'discount_start_date', seasoning_discount.start_date,
				'discount_end_date', seasoning_discount.end_date,
				'discount_percentage', seasoning_discount.discount_percentage)
				AS seasoning_discount
					FROM orders LEFT JOIN 
						(discounts LEFT JOIN 
							seasoning_discount ON discounts.discount_id = seasoning_discount.discount_id) 
						ON orders.seasoning_discount_id = seasoning_discount.discount_id
				WHERE orders.user_id = $1 AND orders.cart_id = $2 AND orders.store_id = stores.store_id),
			(SELECT COALESCE(jsonb_agg(jsonb_build_object(
				'event_discount_max_quantity', event_discount.max_quantity, 
				'event_discount_id', event_discount.discount_id, 
				'product_id', carts.product_id, 
				'product_price', products.price, 
				'product_name', products.name, 
				'product_quantity', carts.product_quantity, 
				'product_picture', products.picture)), '[]')
				AS product_order
					FROM carts 
					LEFT JOIN event_discount ON carts.event_discount_id = event_discount.discount_id 
					LEFT JOIN products ON carts.product_id = products.product_id
				WHERE carts.customer_id = $1 AND carts.cart_id = $2 AND carts.store_id = stores.store_id)
	FROM stores WHERE store_id = $3;
	`

	row := p.db.QueryRow(sqlStatement, customerId, cartId, storeId)

	storeOrder := &swagger.StoreOrderInfo{}
	var shippingDiscountString sql.NullString
	var seasoningDiscountString sql.NullString
	var productOrderString sql.NullString
	err := row.Scan(&storeOrder.StoreId, &storeOrder.StoreName, &storeOrder.StoreShippingFee,
		&shippingDiscountString, &seasoningDiscountString, &productOrderString)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// change string to json
	if shippingDiscountString.Valid {
		var shippingDiscount *swagger.ShippingDiscount
		err := json.Unmarshal([]byte(shippingDiscountString.String), &shippingDiscount)
		if err != nil {
			logrus.Error(err)
		}
		storeOrder.ShippingDiscount = shippingDiscount
	}
	// change string to json
	if seasoningDiscountString.Valid {
		var seasoningDiscount *swagger.SeasoningDiscount
		err := json.Unmarshal([]byte(seasoningDiscountString.String), &seasoningDiscount)
		if err != nil {
			logrus.Error(err)
		}
		storeOrder.SeasoningDiscount = seasoningDiscount
	}
	// change string to json
	if productOrderString.Valid {
		var productOrderArray []swagger.ProductOrderInfo
		err := json.Unmarshal([]byte(productOrderString.String), &productOrderArray)
		if err != nil {
			logrus.Error(err)
		}
		storeOrder.ProductOrder = productOrderArray
	}

	return storeOrder, nil
}

func (p *postgresqlCartRepo) GetCurrentCartID(ctx context.Context, id int64) ([]domain.IDs, error) {
	sqlStatement := `
	SELECT customer_id, cart_id, store_id FROM carts WHERE customer_id = $1
	AND cart_id = (SELECT current_cart_id FROM user_data WHERE user_id = $1)
	GROUP BY carts.store_id, carts.customer_id, carts.cart_id
	`

	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	ids := []domain.IDs{}
	for rows.Next() {
		idInfo := domain.IDs{}
		if err := rows.Scan(&idInfo.UserID, &idInfo.CartID, &idInfo.StoreID); err != nil {
			logrus.Error(err)
			return nil, err
		}

		ids = append(ids, idInfo)
	}

	return ids, nil
}

func (p *postgresqlCartRepo) AddUserCurrentCart(ctx context.Context, id int64) error {
	sqlStatement := `
	UPDATE user_data SET 
	current_cart_id = current_cart_id +1
	WHERE user_id = $1;
	`

	_, err := p.db.Exec(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlCartRepo) UpdateOrderStatusByID(ctx context.Context, customerId int64, storeId int64, status int64) error {
	sqlStatement := `
	UPDATE orders SET status = $1	
	WHERE user_id = $2 AND 
	cart_id = (SELECT current_user_id FROM user_data WHERE user_id = $2) AND
	store_id = $3
	`

	_, err := p.db.Exec(sqlStatement, status, customerId, storeId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlCartRepo) CheckoutOrderInfo(ctx context.Context, customerId int64, storeId int64, totalPrice int64) error {
	sqlStatement := `
	UPDATE orders SET
	status = 1,
	total_price = $1,
	order_date = $2,
	taking_address = (SELECT address FROM user_data WHERE user_id = $3)
	WHERE user_id = $3 AND 
	cart_id = (SELECT current_cart_id FROM user_data WHERE user_id = $3) AND
	store_id = $4
	`
	dt := time.Now().Format("01-02-2006 15:04:05")

	_, err := p.db.Exec(sqlStatement, totalPrice, dt, customerId, storeId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
