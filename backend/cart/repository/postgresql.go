package repository

import (
	"context"
	"database/sql"

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

func (p *postgresqlCartRepo) PostCart(ctx context.Context, cart *swagger.CartInfo, id int64) error {
	sqlStatement := `
	INSERT INTO carts (customer_id, product_id, store_id, cart_id, product_quantity, event_discount_id) VALUES 
	($1,$2,$3,$4,$5,$6) 
	`

	_, err := p.db.Exec(sqlStatement, id, cart.ProductId, cart.StoreId, cart.CartId, cart.ProductQuantity, cart.DiscountId)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (p *postgresqlCartRepo) GetCartByID(ctx context.Context, customerId int64, cartId int64, storeId int64) (*[]swagger.CartInfo, error) {
	sqlStatement := `SELECT 
	customer_id, product_id, store_id, cart_id, product_quantity, event_discount_id 
	FROM carts 
	WHERE customer_id = $1 AND cart_id = $2 AND store_id = $3;
	`

	rows, err := p.db.Query(sqlStatement, customerId, cartId, storeId)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	carts := &[]swagger.CartInfo{}
	for rows.Next() {
		cart := &swagger.CartInfo{}
		err = rows.Scan(&cart.CustomerId, &cart.ProductId, &cart.StoreId, &cart.CartId, &cart.ProductQuantity, &cart.DiscountId)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		*carts = append(*carts, *cart)
	}

	return carts, nil
}

func (p *postgresqlCartRepo) GetByProductID(ctx context.Context, id int64) (*swagger.ProductInfo, error) {
	row := p.db.QueryRow("SELECT product_id, store_id, name, description, picture, price, stock, status FROM products WHERE product_id = $1", id)

	product := &swagger.ProductInfo{}
	if err := row.Scan(&product.ProductId, &product.StoreId, &product.Name, &product.Description, &product.Picture, &product.Price, &product.Stock, &product.Status); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return product, nil
}

func (p *postgresqlCartRepo) DeleteProduct(ctx context.Context, customerId int64, cartId int64, productId int64) error {
	sqlStatement := `
	DELETE FROM carts
	WHERE customer_id = $1 AND cart_id = $2 AND product_id = $3;
	`
	_, err := p.db.Exec(sqlStatement, customerId, cartId, productId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlCartRepo) DeleteOrder(ctx context.Context, customerId int64, cartId int64, storeId int64) error {
	sqlStatement := `
	DELETE FROM orders
	WHERE user_id = $1 AND cart_id = $2 AND store_id = $3;
	`
	_, err := p.db.Exec(sqlStatement, customerId, cartId, storeId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlCartRepo) AddOrder(ctx context.Context, customerId int64, cartId int64, storeId int64, order *swagger.OrderInfo) error {
	sqlStatement := `INSERT INTO orders 
	(cart_id, store_id, user_id, seasoning_discount_id, shipping_discount_id, status, total_price, Order_date, taking_address, taking_method) VALUES 
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := p.db.Exec(sqlStatement, order.CartId, order.StoreId, order.CustomerId, order.SeasoningDiscountId, order.ShippingDiscountId, order.OrderStatus, order.TotalPrice, order.OrderDate, order.TakingAddress, order.TakingMethod)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (p *postgresqlCartRepo) GetUserAddressById(ctx context.Context, id int64) (string, error) {
	sqlStatement := `SELECT address FROM user_data WHERE user_id = $1;`

	row := p.db.QueryRow(sqlStatement, id)

	user_address := ""
	if err := row.Scan(&user_address); err != nil {
		logrus.Error(err)
		return "", err
	}
	return user_address, nil

}

func (p *postgresqlCartRepo) AddUserCartId(ctx context.Context, id int64) error {
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

func (p *postgresqlCartRepo) GetEventDiscountQuantity(ctx context.Context, id int64) (int64, error) {
	sqlStatement := `
	SELECT max_quantity FROM event_discount WHERE discount_id = $1;
	`

	row := p.db.QueryRow(sqlStatement, id)

	var quantity int64 = 0
	if err := row.Scan(&quantity); err != nil {
		logrus.Error(err)
		return 0, err
	}

	return quantity, nil
}

func (p *postgresqlCartRepo) GetOrderById(ctx context.Context, customerId int64, cartId int64, storeId int64) (*swagger.OrderInfo, error) {
	sqlStatement := `SELECT 
	cart_id, store_id, user_id, seasoning_discount_id, shipping_discount_id, status, total_price, Order_date, taking_address, taking_method
	FROM orders 
	WHERE user_id = $1 AND cart_id = $2 AND store_id = $3;
	`

	row := p.db.QueryRow(sqlStatement, customerId, cartId, storeId)

	order := &swagger.OrderInfo{}
	err := row.Scan(
		&order.CartId, &order.StoreId, &order.CustomerId,
		&order.SeasoningDiscountId, &order.ShippingDiscountId,
		&order.OrderStatus, &order.TotalPrice, &order.OrderDate,
		&order.TakingAddress, &order.TakingMethod)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return order, nil
}

func (p *postgresqlCartRepo) CheckoutOrder(ctx context.Context, customerId int64, cartId int64, storeId int64, totalPrice int64, orderDate string) error {
	sqlStatement := `
	UPDATE orders SET 
	total_price = $1, order_date = $2, status = 1
	WHERE user_id = $3 AND cart_id = $4 AND store_id = $5;
	`

	_, err := p.db.Exec(sqlStatement, totalPrice, orderDate,
		customerId, cartId, storeId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlCartRepo) GetStoreShippingFeeByID(ctx context.Context, id int64) (int64, error) {
	row := p.db.QueryRow("SELECT shipping_fee FROM stores WHERE store_id = $1", id)

	var shippingFee int64 = 0
	if err := row.Scan(&shippingFee); err != nil {
		logrus.Error(err)
		return 0, err
	}
	return shippingFee, nil
}

func (p *postgresqlCartRepo) GetMaxPriceByID(ctx context.Context, id int64) (int64, error) {
	sqlStatement := `
	SELECT max_price FROM shipping_discount
	WHERE discount_id = $1;
	`

	row := p.db.QueryRow(sqlStatement, id)

	var maxPrice int64 = 0
	if err := row.Scan(&maxPrice); err != nil {
		logrus.Error(err)
		return 0, err
	}

	return maxPrice, nil
}

func (p *postgresqlCartRepo) GetPercentageByID(ctx context.Context, id int64) (int64, error) {
	sqlStatement := `
	SELECT discount_percentage FROM seasoning_discount
	WHERE discount_id = $1;
	`

	row := p.db.QueryRow(sqlStatement, id)

	var percentage int64 = 0
	if err := row.Scan(&percentage); err != nil {
		logrus.Error(err)
		return 0, err
	}

	return percentage, nil
}

func (p *postgresqlCartRepo) GetOrderByCustomerID(ctx context.Context, id int64) (*[]swagger.OrderInfo, error) {
	sqlStatement := `SELECT 
	user_id, cart_id, store_id, seasoning_discount_id, shipping_discount_id, status, total_price, Order_date, taking_address, taking_method 
	FROM orders
	WHERE user_id = $1;
	`

	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	orders := &[]swagger.OrderInfo{}
	for rows.Next() {
		order := &swagger.OrderInfo{}
		err = rows.Scan(&order.CustomerId, &order.CartId, &order.StoreId, &order.SeasoningDiscountId, &order.ShippingDiscountId, &order.OrderStatus, &order.TotalPrice, &order.OrderDate, &order.TakingAddress, &order.TakingMethod)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		*orders = append(*orders, *order)
	}

	return orders, nil
}
