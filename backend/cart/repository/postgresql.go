package repository

import (
	"context"
	"database/sql"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"

	"github.com/sirupsen/logrus"
)

type postgresqlCartRepo struct {
	db *sql.DB
}

func NewPostgressqlCartRepo(db *sql.DB) domain.CartRepo {
	return &postgresqlCartRepo{db}
}

func (p *postgresqlCartRepo) GetByID(ctx context.Context, customerId string, productId string, storeId string) (*domain.Cart, error) {
	row := p.db.QueryRow("SELECT customer_id, product_id, store_id, product_quantity FROM carts WHERE customer_id = $1 AND product_id = $2 AND store_id = $3", customerId, productId, storeId)
	cart := &domain.Cart{}
	if err := row.Scan(&cart.CustomerID, &cart.ProductID, &cart.StoreID, &cart.ProductQuantity); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return cart, nil
}

func (p *postgresqlCartRepo) PostCart(ctx context.Context, cart *domain.Cart) (*domain.Cart, error) {
	sqlStatement := `
	INSERT INTO carts (product_quantity, customer_id, product_id, store_id) VALUES
    ( $1, (SELECT user_id FROM user_datas WHERE user_id=$2 ), (SELECT product_id FROM products WHERE product_id=$3), (SELECT store_id FROM stores WHERE store_id=$4) );
	`
	_, err := p.db.Exec(sqlStatement, cart.ProductQuantity, cart.CustomerID, cart.ProductID, cart.StoreID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	sqlStatement = `
	SELECT customer_id, product_id, store_id, product_quantity FROM carts WHERE customer_id = $1 AND product_id = $2 AND store_id = $3;
	`
	_cart := &domain.Cart{}
	row := p.db.QueryRow(sqlStatement, cart.CustomerID, cart.ProductID, cart.StoreID)
	if err := row.Scan(&_cart.CustomerID, &_cart.ProductID, &_cart.StoreID, &_cart.ProductQuantity); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return _cart, nil
}
