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
			jsonb_agg(jsonb_build_object(
			'discount_max_quantity', max_quantity,
			'product_id', product_id,
			'discount_name', name,
			'discount_description', description,
			'discount_id', discount_id,
			'status', status
			)) 
			AS event_discount
			FROM discounts NATURAL JOIN 
			(SELECT event_discount.discount_id, event_discount.max_quantity, event_discount.product_id FROM event_discount WHERE event_discount.product_id = products.product_id)
		) AS event_discount_array,
		(SELECT (jsonb_agg(jsonb_build_object(
			'product_id', product_id, 
			'label_name', label_name, 
			'required', required, 
			'item_array', 
			(SELECT (jsonb_agg(jsonb_build_object(
					'name', label_item.item_name))) AS item_array
			FROM label_item
			WHERE label_item.label_name = product_label.label_name)
		))) AS product_label_array
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
			jsonb_agg(jsonb_build_object(
			'discount_max_quantity', max_quantity,
			'product_id', product_id,
			'discount_name', name,
			'discount_description', description,
			'discount_id', discount_id,
			'status', status
			)) 
			AS event_discount
			FROM discounts NATURAL JOIN 
			(SELECT event_discount.discount_id, event_discount.max_quantity, event_discount.product_id FROM event_discount WHERE event_discount.product_id = products.product_id)
		) AS event_discount_array,
		(SELECT (jsonb_agg(jsonb_build_object(
			'product_id', product_id, 
			'label_name', label_name, 
			'required', required, 
			'item_array', 
			(SELECT (jsonb_agg(jsonb_build_object(
					'name', label_item.item_name))) AS item_array
			FROM label_item
			WHERE label_item.label_name = product_label.label_name)
		))) AS product_label_array
		FROM product_label WHERE product_label.product_id = products.product_id)
	FROM products WHERE store_id = $1;
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
