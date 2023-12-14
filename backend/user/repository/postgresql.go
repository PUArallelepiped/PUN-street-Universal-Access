package repository

import (
	"context"
	"database/sql"
	"errors"

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

func (p *postgresqlUserRepo) Login(ctx context.Context, email string, password string) (int, error) {
	sqlStatement := `
	SELECT password, authority FROM user_data WHERE email = $1;
	`
	var hashedPassword string
	var authority int
	err := p.db.QueryRow(sqlStatement, email).Scan(&hashedPassword, &authority)
	if err != nil {
		return 0, err
	}
	if hashedPassword != password {
		return 0, errors.New("wrong password")
	}

	return authority, nil
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

func (p *postgresqlUserRepo) RegisterUser(ctx context.Context, user *swagger.RegisterInfo, authority string) (int, error) {
	sqlStatement := `
		INSERT INTO user_data (name, password, email, address, phone_number, birthday, authority, current_cart_id, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING user_id;
	`
	id := 0
	err := p.db.QueryRow(sqlStatement, user.UserName, user.Password, user.UserEmail, user.Address, user.Phone, user.Birthday, authority, 1, 1).Scan(&id)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return id, nil
}

func (p *postgresqlUserRepo) RegisterStore(ctx context.Context, storeInfo swagger.StoreRegisterInfo, id int) error {
	sqlStatement := `
		INSERT INTO stores (store_id, name, rate, rate_count, address, picture, description, shipping_fee, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`
	_, err := p.db.Exec(sqlStatement, id, storeInfo.Name, 0, 0, storeInfo.Address, storeInfo.Picture, storeInfo.Description, storeInfo.ShippingFee, 1)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (p *postgresqlUserRepo) CheckEmail(ctx context.Context, email string) (bool, error) {
	sqlStatement := `
		SELECT EXISTS(SELECT 1 FROM user_data WHERE email = $1);
	`
	var exists bool
	err := p.db.QueryRow(sqlStatement, email).Scan(&exists)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	return exists, nil
}
