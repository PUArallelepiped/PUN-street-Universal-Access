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
	SELECT discounts.discount_id, status, description, name, max_price 
	FROM discounts NATURAL JOIN shipping_discount
	WHERE shipping_discount.store_id = $1 AND discounts.status = 1;
	`
	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	shipping_discounts := []swagger.ShippingDiscount{}
	for rows.Next() {
		shipping_discount := swagger.ShippingDiscount{}
		err := rows.Scan(&shipping_discount.DiscountId, &shipping_discount.Status, &shipping_discount.DiscountDescription, &shipping_discount.DiscountName, &shipping_discount.DiscountMaxPrice)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		shipping_discounts = append(shipping_discounts, shipping_discount)
	}
	return shipping_discounts, nil
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

func (p *postgresqlDiscountRepo) DisableDiscountByDiscountID(ctx context.Context, id int64) error {
	sqlStatement := `
	UPDATE discounts
	SET status = 0
	WHERE discount_id = $1;
	`
	_, err := p.db.Exec(sqlStatement, id)

	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlDiscountRepo) IsExistShippingDiscountByStoreID(ctx context.Context, id int64) (bool, error) {
	sqlStatement := `
	SELECT COUNT(*) > 0 FROM shipping_discount WHERE store_id = $1
	`
	row := p.db.QueryRow(sqlStatement, id)

	exist := false
	err := row.Scan(&exist)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	return exist, nil
}

func (p *postgresqlDiscountRepo) AddEvent(ctx context.Context, event *swagger.EventDiscount, id int64) error {
	sqlStatement := `
	WITH ins1 AS (
	INSERT INTO discounts(discount_type, status, description, name)
	VALUES (3, 1, $1, $2)
	RETURNING discount_id)
	INSERT INTO event_discount (discount_id, max_quantity, product_id)
	select discount_id, $3, $4 FROM ins1;
	`
	_, err := p.db.Exec(sqlStatement, event.DiscountDescription, event.DiscountName,
		event.DiscountMaxQuantity, id)

	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlDiscountRepo) GetAllSeasoning(ctx context.Context) ([]swagger.SeasoningDiscount, error) {
	sqlStatement := `
	SELECT discounts.discount_id, name, description, start_date, end_date, discount_percentage, status 
	FROM discounts NATURAL JOIN seasoning_discount
	WHERE discounts.status = 1;
	`
	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	seasoningDiscounts := []swagger.SeasoningDiscount{}
	for rows.Next() {
		seasoningDiscount := swagger.SeasoningDiscount{}
		err := rows.Scan(&seasoningDiscount.DiscountId, &seasoningDiscount.DiscountName, &seasoningDiscount.DiscountDescription, &seasoningDiscount.DiscountStartDate, &seasoningDiscount.DiscountEndDate, &seasoningDiscount.DiscountPercentage, &seasoningDiscount.Status)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		seasoningDiscounts = append(seasoningDiscounts, seasoningDiscount)
	}
	return seasoningDiscounts, nil
}

func (p *postgresqlDiscountRepo) GetAllEventByProductID(ctx context.Context, id int64) ([]swagger.EventDiscount, error) {
	sqlStatement := `
	SELECT discounts.discount_id, name, description, max_quantity, product_id,  status 
	FROM discounts NATURAL JOIN event_discount
	WHERE product_id = $1 AND discounts.status = 1;
	`
	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	eventDiscounts := []swagger.EventDiscount{}
	for rows.Next() {
		eventDiscount := swagger.EventDiscount{}
		err := rows.Scan(&eventDiscount.DiscountId, &eventDiscount.DiscountName, &eventDiscount.DiscountDescription, &eventDiscount.DiscountMaxQuantity, &eventDiscount.ProductId, &eventDiscount.Status)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		eventDiscounts = append(eventDiscounts, eventDiscount)
	}
	return eventDiscounts, nil
}
