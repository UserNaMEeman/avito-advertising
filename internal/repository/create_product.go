package repository

import (
	"advertising/entity"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CreateProduct struct {
	db *sqlx.DB
}

func NewCreateProduct(db *sqlx.DB) *CreateProduct {
	return &CreateProduct{db: db}
}

func (prod *CreateProduct) CreateProductDB(ctx context.Context, product entity.Product) (int, error) {
	var n int
	query := fmt.Sprintf(
		`INSERT INTO %s (title, description, url1, url2, url3, price)
		 values ($1, $2, $3, $4, $5, $6) RETURNING id`,
		tableProduct)
	// id, err := prod.db.QueryxContext()

	id, err := prod.db.QueryxContext(
		ctx, query, product.Title, product.Description,
		product.URL1, product.URL2, product.URL3, product.Price)
	if err != nil {
		return 0, err
	}
	for id.Next() {
		err = id.Scan(&n)
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}
