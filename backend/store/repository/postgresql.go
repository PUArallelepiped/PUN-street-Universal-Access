package repository

import (
	"context"
	"database/sql"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"

	"github.com/sirupsen/logrus"
)

type postgresqlStoreRepo struct {
	db *sql.DB
}

func NewPostgressqlStoreRepo(db *sql.DB) domain.StoreRepo {
	return &postgresqlStoreRepo{db}
}

func (p *postgresqlStoreRepo) GetByID(ctx context.Context, id string) (*domain.Store, error) {
	row := p.db.QueryRow("SELECT store_id, name, address, email, phone FROM stores WHERE store_id = $1", id)
	s := &domain.Store{}
	if err := row.Scan(&s.ID, &s.Name, &s.Address, &s.Email, &s.Phone); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}

func (p *postgresqlStoreRepo) GetAll(ctx context.Context) ([]domain.Store, error) {
	rows, err := p.db.Query("SELECT store_id, name, address, email, phone FROM stores")
	if err != nil {
		logrus.Error(err)
	}
	l := []domain.Store{}
	for rows.Next() {
		s := domain.Store{}
		err := rows.Scan(&s.ID, &s.Name, &s.Address, &s.Email, &s.Phone)
		if err != nil {
			logrus.Error(err)
		}
		l = append(l, s)
	}
	return l, nil
}
