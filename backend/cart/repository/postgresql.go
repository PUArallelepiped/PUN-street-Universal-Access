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
