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
func (p *postgresqlHelloRepo) SayHello(ctx context.Context) (*domain.HelloMsg, error) {
	// TODO: need sqlmock to test
	// row := p.db.QueryRow("SHOW TABLES")
	name := &domain.HelloMsg{Name: "Gopher"}
	// if err := row.Scan(); err != nil {
	// 	return nil, err
	// }
	return name, nil
}
