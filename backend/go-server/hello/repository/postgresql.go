package repository

import (
	"context"
	"database/sql"
	"main/domain"
)

type postgresqlHelloRepo struct {
	db *sql.DB
}

func NewPostgressqlHelloRepo(db *sql.DB) domain.HelloRepo {
	return &postgresqlHelloRepo{db}
}

// sayHello implements domain.HelloRepo.
func (p *postgresqlHelloRepo) SayHello(ctx context.Context) error {
	row := p.db.QueryRow("")
	if err := row.Scan(); err != nil {
		return err
	}
	return nil
}
