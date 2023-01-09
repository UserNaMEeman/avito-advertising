package repository

import (
	"advertising/entity"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductList struct {
	db *sqlx.DB
}

func NewProductListDB(db *sqlx.DB) *ProductList {
	return &ProductList{db: db}
}

func (prod *ProductList) GetProductsDB(ctx context.Context, sort entity.SortProductDB, offset, limit int) ([]entity.ResponseList, error) {
	query := fmt.Sprintf(`SELECT title, url1, price FROM %s order by %s %s offset %d limit %d`,
		tableProduct, sort.Field, sort.Type, offset, limit)
	// fmt.Println(query)
	products := make([]entity.ResponseList, 0, limit)
	// rows, err := prod.db.QueryContext(ctx, `SELECT title, url1, price FROM %s order by %s %s
	// offset %s limit %s`, tableProduct, sort.Field, sort.Type, offset, limit)
	rows, err := prod.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product entity.ResponseList
		err = rows.Scan(&product.Title, &product.URL1, &product.Price)
		if err != nil {
			return nil, err
		}
		// fmt.Println(product)
		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return products, nil
}
