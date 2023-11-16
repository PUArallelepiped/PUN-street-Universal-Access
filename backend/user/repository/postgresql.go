package repository

import (
	"context"
	"database/sql"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
)

type postgresqlUserRepo struct {
	db *sql.DB
}

func NewPostgressqlUserRepo(db *sql.DB) domain.UserRepo {
	return &postgresqlUserRepo{db}
}

func (p *postgresqlUserRepo) Login(ctx context.Context, email string, password string) error {
	sqlStatement := `
	SELECT password FROM users WHERE email = $1;
	`
	var hashedPassword string
	err := p.db.QueryRow(sqlStatement, email).Scan(&hashedPassword)
	if err != nil {
		return err
	}
	if hashedPassword != password {
		return err
	}

	return nil
}
