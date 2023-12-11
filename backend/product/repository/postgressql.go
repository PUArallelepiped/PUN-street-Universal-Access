package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/sirupsen/logrus"
)

type postgresqlProductRepo struct {
	db *sql.DB
}

func NewPostgressqlProductRepo(db *sql.DB) domain.ProductRepo {
	return &postgresqlProductRepo{db}
}

func (p *postgresqlProductRepo) GetProductByID(ctx context.Context, id int64) (*swagger.ProductInfoWithLabelAndDiscount, error) {
	sqlStatement := `
	SELECT product_id, store_id, name, description, picture, stock, price, status,
		(SELECT 
			coalesce(jsonb_agg(jsonb_build_object(
			'discount_max_quantity', max_quantity,
			'product_id', product_id,
			'discount_name', name,
			'discount_description', description,
			'discount_id', discount_id,
			'status', status
			)) , '[]')
			AS event_discount
			FROM discounts NATURAL JOIN 
			(SELECT event_discount.discount_id, event_discount.max_quantity, event_discount.product_id FROM event_discount WHERE event_discount.product_id = products.product_id)
		) AS event_discount_array,
		(SELECT 
			coalesce(jsonb_agg(jsonb_build_object(
			'product_id', product_id, 
			'label_name', label_name, 
			'required', required, 
			'item_array', 
			(SELECT (jsonb_agg(jsonb_build_object(
					'name', label_item.item_name))) AS item_array
			FROM label_item
			WHERE label_item.label_name = product_label.label_name AND label_item.product_id = product_label.product_id)
		)), '[]') AS product_label_array
		FROM product_label WHERE product_label.product_id = products.product_id)
	FROM products WHERE product_id = $1;
	`
	row := p.db.QueryRow(sqlStatement, id)

	product := &swagger.ProductInfoWithLabelAndDiscount{}
	var productLabelArrayString sql.NullString
	var eventDiscountArrayString sql.NullString
	if err := row.Scan(
		&product.ProductId, &product.StoreId, &product.Name,
		&product.Description, &product.Picture, &product.Stock,
		&product.Price, &product.Status,
		&eventDiscountArrayString, &productLabelArrayString); err != nil {
		logrus.Error(err)
		return nil, err
	}
	// change string to json
	if eventDiscountArrayString.Valid {
		var eventDiscountArray []swagger.EventDiscount
		err := json.Unmarshal([]byte(eventDiscountArrayString.String), &eventDiscountArray)
		if err != nil {
			logrus.Error(err)
		}
		product.EventDiscountArray = eventDiscountArray
	}

	// change string to json
	if productLabelArrayString.Valid {
		var productLabelArray []swagger.ProductLabelInfo
		err := json.Unmarshal([]byte(productLabelArrayString.String), &productLabelArray)
		if err != nil {
			logrus.Error(err)
		}
		product.ProductLabelArray = productLabelArray
	}

	return product, nil
}

func (p *postgresqlProductRepo) GetProductsByStoreID(ctx context.Context, id int64) (*[]swagger.ProductInfoWithLabelAndDiscount, error) {
	// same and GetProductByID but the products where is store_id
	sqlStatement := `
	SELECT product_id, store_id, name, description, picture, stock, price, status,
		(SELECT 
			coalesce(jsonb_agg(jsonb_build_object(
			'discount_max_quantity', max_quantity,
			'product_id', product_id,
			'discount_name', name,
			'discount_description', description,
			'discount_id', discount_id,
			'status', status
			)) , '[]')
			AS event_discount
			FROM discounts NATURAL JOIN 
			(SELECT event_discount.discount_id, event_discount.max_quantity, event_discount.product_id FROM event_discount WHERE event_discount.product_id = products.product_id)
		) AS event_discount_array,
		(SELECT 
			coalesce(jsonb_agg(jsonb_build_object(
			'product_id', product_id, 
			'label_name', label_name, 
			'required', required, 
			'item_array', 
			(SELECT (jsonb_agg(jsonb_build_object(
					'name', label_item.item_name))) AS item_array
			FROM label_item
			WHERE label_item.label_name = product_label.label_name AND label_item.product_id = product_label.product_id)
		)), '[]') AS product_label_array
		FROM product_label WHERE product_label.product_id = products.product_id)
	FROM products WHERE store_id = $1 AND status != 0;
	`
	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	products := &[]swagger.ProductInfoWithLabelAndDiscount{}
	for rows.Next() {
		product := &swagger.ProductInfoWithLabelAndDiscount{}
		var productLabelArrayString sql.NullString
		var eventDiscountArrayString sql.NullString

		if err := rows.Scan(
			&product.ProductId, &product.StoreId, &product.Name,
			&product.Description, &product.Picture, &product.Stock,
			&product.Price, &product.Status,
			&eventDiscountArrayString, &productLabelArrayString); err != nil {
			logrus.Error(err)
			return nil, err
		}
		// change string to json
		if eventDiscountArrayString.Valid {
			var eventDiscountArray []swagger.EventDiscount
			err := json.Unmarshal([]byte(eventDiscountArrayString.String), &eventDiscountArray)
			if err != nil {
				logrus.Error(err)
			}
			product.EventDiscountArray = eventDiscountArray
		}

		// change string to json
		if productLabelArrayString.Valid {
			var productLabelArray []swagger.ProductLabelInfo
			err := json.Unmarshal([]byte(productLabelArrayString.String), &productLabelArray)
			if err != nil {
				logrus.Error(err)
			}
			product.ProductLabelArray = productLabelArray
		}
		*products = append(*products, *product)
	}

	return products, nil
}

func (p *postgresqlProductRepo) ChangeProductStatusByProductID(ctx context.Context, id int64, status int64) error {
	sqlStatement := `
	UPDATE products SET status = $2 WHERE product_id = $1
	`

	_, err := p.db.Exec(sqlStatement, id, status)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlProductRepo) AddProductByStoreID(ctx context.Context, id int64, product *swagger.ProductInfoWithLabelAndDiscount) (int64, error) {
	sqlStatement := `
	INSERT INTO products (store_id, name, description, picture, price, stock, status) VALUES
    ($1, $2, $3, $4, $5, $6, $7)
    RETURNING product_id;
	`

	row := p.db.QueryRow(sqlStatement, product.StoreId, product.Name,
		product.Description, product.Picture, product.Price,
		product.Stock, product.Status)

	var productID int64
	if err := row.Scan(&productID); err != nil {
		logrus.Error(err)
		return 0, err
	}

	return productID, nil
}

func (p *postgresqlProductRepo) AddDiscountByProductID(ctx context.Context, event *swagger.EventDiscount) error {
	sqlStatement := `
	WITH ins1 AS (
	INSERT INTO discounts(discount_type, status, description, name)
	VALUES (3, 1, $1, $2)
	RETURNING discount_id)
	INSERT INTO event_discount (discount_id, max_quantity, product_id)
	select discount_id, $3, $4 FROM ins1;
	`
	_, err := p.db.Exec(sqlStatement, event.DiscountDescription, event.DiscountName,
		event.DiscountMaxQuantity, event.ProductId)

	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlProductRepo) AddProductLabel(ctx context.Context, productId int64, label_name string, required bool) error {
	sqlStatement := `
	INSERT INTO product_label (product_id, label_name, required) VALUES
    ($1, $2, $3)
	`
	_, err := p.db.Exec(sqlStatement, productId, label_name, required)

	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlProductRepo) AddProductLabelItem(ctx context.Context, productId int64, labelName string, name string) error {
	sqlStatement := `
	INSERT INTO label_item (product_id, label_name, item_name) VALUES
    ($1, $2, $3)
	`
	_, err := p.db.Exec(sqlStatement, productId, labelName, name)

	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
