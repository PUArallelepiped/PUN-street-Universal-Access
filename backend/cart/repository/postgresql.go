package repository

import (
	"context"
	"database/sql"
	"fmt"

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

func (p *postgresqlCartRepo) PostCart(ctx context.Context, cart *swagger.CartInfo) error {
	sqlStatement := `
	INSERT INTO carts (customer_id, product_id, store_id, cart_id, product_quantity, event_discount_id) VALUES 
	($1,$2,$3,$4,$5,$6) 
	ON CONFLICT (customer_id, product_id, store_id) DO UPDATE 
	SET product_quantity = excluded.product_quantity;
	`

	_, err := p.db.Exec(sqlStatement, cart.CustomerId, cart.ProductId, cart.StoreId, cart.CartId, cart.ProductQuantity, cart.DiscountId)
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
func (p *postgresqlCartRepo) AddOrder(ctx context.Context, customerId int64, cartId int64, storeId int64, order *swagger.OrderInfo) error {
	sqlStatement := `INSERT INTO orders 
	(cart_id, store_id, user_id, seasoning_discount_id, shipping_discount_id, status, total_price, Order_date, taking_address, taking_method) VALUES 
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	fmt.Println(order)

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

func (p *postgresqlCartRepo) GetCartArrayByCustomerID(ctx context.Context, id int64) (*[]swagger.CartInfo, error) {
	sqlStatement := `SELECT 
	customer_id, product_id, store_id, cart_id, product_quantity, event_discount_id 
	FROM carts 
	WHERE customer_id = $1;
	`

	rows, err := p.db.Query(sqlStatement, id)
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
