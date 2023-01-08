package repository

import (
	"github.com/jmoiron/sqlx"
)

type ProductList struct {
	db *sqlx.DB
}

func NewProductListDB(db *sqlx.DB) *ProductList {
	return &ProductList{db: db}
}

// func (prod *ProductList) GetProducts(ctx context.Context, sort entity.SortProductDB) ([]entity.Product, error) {
// 	query := fmt.Sprintf(`SELECT title, url1, price FROM %s order by %s %s
// 	offset %d limit %d`,
// 		tableProduct, sort.Field, sort.Type, 5, 5)
// 	prod.db.GetContext()
// }
