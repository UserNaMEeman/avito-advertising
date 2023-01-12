package repository

import (
	"advertising/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

var tableProduct string = "products"

type GetList interface {
	GetProductsDB(ctx context.Context, sort entity.SortProductDB, offset int, limit int) ([]entity.ResponseList, error)
}

type GetAdt interface {
	GetProductDB(ctx context.Context, id int, fields entity.ResponseFileds) (entity.Product, error)
}

type CreateAdt interface {
	CreateProductDB(ctx context.Context, product entity.Product) (int, error)
}

type Repository struct {
	GetList
	GetAdt
	CreateAdt
}

func NewProductPostgres(db *sqlx.DB) *Repository {
	return &Repository{
		CreateAdt: NewCreateProduct(db),
		GetList:   NewProductListDB(db),
		GetAdt:    NewProductRespDB(db),
	}
}
