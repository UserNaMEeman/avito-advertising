package repository

import (
	"advertising/entity"
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductResp struct {
	db *sqlx.DB
}

func NewProductRespDB(db *sqlx.DB) *ProductResp {
	return &ProductResp{db: db}
}

func (prod *ProductResp) GetProductDB(ctx context.Context, id int, fields entity.ResponseFileds) (entity.Product, error) {

	var query string
	var row *sql.Row
	product := entity.Product{}
	// fmt.Println(fields.Descr, fields.Photo)
	// // switch fields.Descr && fields.Photo {
	if !fields.Descr && !fields.Photo {
		query = fmt.Sprintf(`SELECT title, url1, price FROM %s where id = %d`,
			tableProduct, id)
		row = prod.db.QueryRowContext(ctx, query)
		if err := row.Scan(&product.Title, &product.URL1, &product.Price); err != nil {
			return product, err
		}
	}
	if fields.Descr && !fields.Photo {
		query = fmt.Sprintf(`SELECT title, description, url1, price FROM %s where id = %d`,
			tableProduct, id)
		row = prod.db.QueryRowContext(ctx, query)
		if err := row.Scan(&product.Title, &product.Description, &product.URL1, &product.Price); err != nil {
			return product, err
		}
	}
	if !fields.Descr && fields.Photo {
		query = fmt.Sprintf(`SELECT title, url1, url2, url3, price FROM %s where id = %d`,
			tableProduct, id)
		row = prod.db.QueryRowContext(ctx, query)
		if err := row.Scan(&product.Title, &product.URL1, &product.URL2, &product.URL3, &product.Price); err != nil {
			return product, err
		}
	}
	if fields.Descr && fields.Photo {
		query = fmt.Sprintf(`SELECT title, description, url1, url2, url3, price FROM %s where id = %d`,
			tableProduct, id)
		row = prod.db.QueryRowContext(ctx, query)
		if err := row.Scan(&product.Title, &product.Description, &product.URL1, &product.URL2, &product.URL3,
			&product.Price); err != nil {
			return product, err
		}
	}
	return product, nil
}
