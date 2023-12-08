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

func (p *postgresqlCartRepo) GetAllHistoryById(ctx context.Context, id int64) (*[]swagger.HistoryInfo, error) {
	sqlStatement := `
	SELECT orders.store_id, orders.cart_id, orders.order_date, orders.total_price, orders.user_id AS customer_id, orders.status, 
	stores.name AS store_name, stores.picture AS store_picture, stores.rate AS store_rate
	FROM orders LEFT JOIN stores ON orders.store_id = stores.store_id WHERE orders.user_id = $1 AND orders.status = 6;
	`

	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	historyArray := &[]swagger.HistoryInfo{}
	for rows.Next() {
		history := &swagger.HistoryInfo{}
		err := rows.Scan(&history.StoreId, &history.CartId, &history.OrderDate,
			&history.TotalPrice, &history.CustomerId, &history.Status,
			&history.StoreName, &history.StorePicture, &history.StoreRate)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		*historyArray = append(*historyArray, *history)
	}

	return historyArray, nil
}
