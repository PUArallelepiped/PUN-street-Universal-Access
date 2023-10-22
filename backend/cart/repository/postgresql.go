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
	sqlStatement := `
	SELECT customer_id, product_id, store_id, product_quantity FROM carts WHERE customer_id = $1 AND product_id = $2 AND store_id = $3
	`
	row := p.db.QueryRow(sqlStatement, customerId, productId, storeId)
	cart := &domain.Cart{}
	if err := row.Scan(&cart.CustomerId, &cart.ProductId, &cart.StoreId, &cart.ProductQuantity); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return cart, nil
}

func (p *postgresqlCartRepo) PostCart(ctx context.Context, cart *domain.Cart) (*domain.Cart, error) {
	sqlStatement := `
	INSERT INTO carts VALUES ($1,$2,$3,$4) 
	ON CONFLICT (customer_id, product_id, store_id) DO UPDATE 
	SET product_quantity = excluded.product_quantity;
	`

	_, err := p.db.Exec(sqlStatement, cart.CustomerId, cart.ProductId, cart.StoreId, cart.ProductQuantity)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	sqlStatement = `
	SELECT customer_id, product_id, store_id, product_quantity FROM carts WHERE customer_id = $1 AND product_id = $2 AND store_id = $3;
	`
	_cart := &domain.Cart{}
	row := p.db.QueryRow(sqlStatement, cart.CustomerId, cart.ProductId, cart.StoreId)
	if err := row.Scan(&_cart.CustomerId, &_cart.ProductId, &_cart.StoreId, &_cart.ProductQuantity); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return _cart, nil
}
