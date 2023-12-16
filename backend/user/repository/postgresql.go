package repository

import (
	"context"
	"database/sql"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type postgresqlUserRepo struct {
	db *sql.DB
}

func NewPostgressqlUserRepo(db *sql.DB) domain.UserRepo {
	return &postgresqlUserRepo{db}
}

func (p *postgresqlUserRepo) GetByID(ctx context.Context, id int64) (*swagger.UserData, error) {
	row := p.db.QueryRow("SELECT * FROM user_data WHERE user_id = $1", id)
	s := &swagger.UserData{}
	if err := row.Scan(&s.UserId, &s.UserName, &s.Password, &s.UserEmail, &s.Address, &s.Phone, &s.Birthday, &s.Authority, &s.CartId, &s.Status); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}

func (p *postgresqlUserRepo) GetAllUser(ctx context.Context) ([]swagger.UserDataShort, error) {
	sqlStatement := `
		SELECT 	email, user_id, name, authority, status
		FROM user_data;
	`

	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	l := []swagger.UserDataShort{}
	for rows.Next() {
		s := swagger.UserDataShort{}
		err := rows.Scan(&s.UserEmail, &s.UserId, &s.UserName, &s.Authority, &s.Status)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		l = append(l, s)
	}

	return l, nil
}

func (p *postgresqlUserRepo) GetAllOrder(ctx context.Context) ([]swagger.OrderInfoShort, error) {
	sqlStatement := `
		SELECT orders.cart_id, orders.store_id, orders.order_date, orders.user_id, user_data.name  
		FROM orders LEFT JOIN user_data 
		ON orders.user_id = user_data.user_id
	`

	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	orders := []swagger.OrderInfoShort{}
	for rows.Next() {
		order := swagger.OrderInfoShort{}
		err := rows.Scan(&order.CartId, &order.StoreId, &order.OrderDate, &order.UserId, &order.UserName)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (p *postgresqlUserRepo) BanUser(ctx context.Context, id int64) error {
	sqlStatement := `
	UPDATE user_data SET status = 0 WHERE user_id = $1
	`

	_, err := p.db.Exec(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlUserRepo) UnBanUser(ctx context.Context, id int64) error {
	sqlStatement := `
	UPDATE user_data SET status = 1 WHERE user_id = $1
	`

	_, err := p.db.Exec(sqlStatement, id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
