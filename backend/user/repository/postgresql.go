package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
)

type postgresqlUserRepo struct {
	db *sql.DB
}

func NewPostgressqlUserRepo(db *sql.DB) domain.UserRepo {
	return &postgresqlUserRepo{db}
}

func (p *postgresqlUserRepo) Login(ctx context.Context, email string, password string) (int, error) {
	sqlStatement := `
	SELECT password, authority FROM users WHERE email = $1;
	`
	var hashedPassword string
	var authority int
	err := p.db.QueryRow(sqlStatement, email).Scan(&authority, &hashedPassword)
	if err != nil {
		return 0, err
	}
	if hashedPassword != password {
		return 0, errors.New("wrong password")
	}

	return authority, nil
}
