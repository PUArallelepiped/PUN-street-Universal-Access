package repository

import (
	"context"
	"database/sql"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/sirupsen/logrus"
)

type postgresqlDiscountRepo struct {
	db *sql.DB
}

func NewPostgressqlDiscountRepo(db *sql.DB) domain.DiscountRepo {
	return &postgresqlDiscountRepo{db}
}

func (p *postgresqlDiscountRepo) GetShippingByStoreID(ctx context.Context, id int64) ([]swagger.ShippingDiscount, error) {
	sqlStatement := `
	SELECT discount_id, max_price FROM shipping_discount WHERE store_id = $1
	`
	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	shipping_discounts := []swagger.ShippingDiscount{}
	for rows.Next() {
		shipping_discount := swagger.ShippingDiscount{}
		err := rows.Scan(&shipping_discount.DiscountId, &shipping_discount.DiscountMaxPrice)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		shipping_discounts = append(shipping_discounts, shipping_discount)
	}
	return shipping_discounts, nil
}

func (p *postgresqlDiscountRepo) GetByDiscountID(ctx context.Context, id int64) (*swagger.DiscountInfo, error) {
	sqlStatement := `
	SELECT discount_id, discount_type, status FROM discounts WHERE discount_id = $1
	`

	row := p.db.QueryRow(sqlStatement, id)
	d := &swagger.DiscountInfo{}
	if err := row.Scan(&d.DiscountId, &d.DiscountType, &d.Status); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return d, nil
}

func (p *postgresqlDiscountRepo) AddSeasoning(ctx context.Context, seasoning *swagger.SeasoningDiscount) error {
	sqlStatement := `
	WITH ins1 AS (
	INSERT INTO discounts(discount_type, status, description, name)
	VALUES (1, 1, $1, $2)
	RETURNING discount_id)
	INSERT INTO seasoning_discount (discount_id, discount_percentage, start_date, end_date)
	select discount_id, $3, $4, $5 FROM ins1;
	`

	_, err := p.db.Exec(sqlStatement, seasoning.DiscountDescription, seasoning.DiscountName,
		seasoning.DiscountPercentage, seasoning.DiscountStartDate, seasoning.DiscountEndDate)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlDiscountRepo) AddShipping(ctx context.Context, shipping *swagger.ShippingDiscount, id int64) error {
	sqlStatement := `
	WITH ins1 AS (
	INSERT INTO discounts(discount_type, status, description, name)
	VALUES (2, 1, $1, $2)
	RETURNING discount_id)
	INSERT INTO shipping_discount (discount_id, max_price, store_id)
	select discount_id, $3, $4 FROM ins1;
	`
	_, err := p.db.Exec(sqlStatement, shipping.DiscountDescription, shipping.DiscountName,
		shipping.DiscountMaxPrice, id)

	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
