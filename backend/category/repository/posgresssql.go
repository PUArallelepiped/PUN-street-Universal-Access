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
	FROM categorys;
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
