package repository

import (
	"context"
	"database/sql"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/sirupsen/logrus"
)

type postgresqlCategoty struct {
	db *sql.DB
}

func NewPostgressqlCategoryRepo(db *sql.DB) domain.CategoryRepo {
	return &postgresqlCategoty{db}
}

func (p *postgresqlCategoty) GetAllCategory(ctx context.Context) (*[]swagger.Category, error) {
	sqlStatement := `SELECT 
	category_id, name
	FROM categories;
	`

	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	categorys := &[]swagger.Category{}
	for rows.Next() {
		category := &swagger.Category{}
		err = rows.Scan(&category.CategoryId, &category.CategoryName)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		*categorys = append(*categorys, *category)
	}

	return categorys, nil

}

func (p *postgresqlCategoty) AddCategoryToStore(ctx context.Context, store_id int64, category_id int64) error {
	sqlStatement := `INSERT INTO labels 
	(store_id, category_id) VALUES
    ($1, $2)
	`

	_, err := p.db.Exec(sqlStatement, store_id, category_id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *postgresqlCategoty) DeleteCategoryToStore(ctx context.Context, store_id int64, category_id int64) error {
	sqlStatement := `
	DELETE FROM labels 
	WHERE category_id = $1 AND store_id = $2;
	`

	_, err := p.db.Exec(sqlStatement, category_id, store_id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
