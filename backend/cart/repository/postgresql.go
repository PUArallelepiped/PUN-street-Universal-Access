package repository

import (
	"context"
	"database/sql"
	"strconv"

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
	row := p.db.QueryRow("SELECT customer_id, product_id, store_id, product_quantity FROM cart WHERE customer_id = $1 AND product_id = $2 AND store_id = $3", customerId, productId, storeId)
	cart := &domain.Cart{}
	if err := row.Scan(&cart.CustomerID, &cart.ProductID, &cart.StoreID, &cart.ProductQuantity); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return cart, nil
}

func (p *postgresqlCartRepo) PostCart(ctx context.Context, cart *domain.Cart) (*domain.Cart, error) {
	sqlStatement := `
	INSERT INTO cart (product_quantity, customer_id, product_id, store_id) VALUES
    ( $1, (SELECT id from UserData LIMIT 1 OFFSET $2), (SELECT product_id from product LIMIT 1 OFFSET $3), (SELECT store_id from stores LIMIT 1 OFFSET $4));
	`
	_, err := p.db.Exec(sqlStatement, cart.ProductQuantity, cart.CustomerID, cart.ProductID, cart.StoreID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	_cart := &domain.Cart{}
	row := p.db.QueryRow("SELECT customer_id, product_id, store_id, product_quantity FROM cart WHERE customer_id = $1 AND product_id = $2 AND store_id = $3", add(cart.CustomerID, 1), add(cart.ProductID, 1), add(cart.StoreID, 1))
	if err := row.Scan(&_cart.CustomerID, &_cart.ProductID, &_cart.StoreID, &_cart.ProductQuantity); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return _cart, nil
}

func add(s string, n int) string {
	i, err := strconv.Atoi(s)
	if err != nil {
		return s
	}
	return strconv.Itoa(i + n)
}
