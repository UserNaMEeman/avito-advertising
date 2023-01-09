package repository

import (
	"advertising/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

var tableProduct string = "products"

type GetList interface {
	GetProductsDB(context.Context, entity.SortProductDB, int, int) ([]entity.ResponseList, error)
}

type GetAdt interface{}

type CreateAdt interface {
	CreateProductDB(context.Context, entity.Product) (int, error)
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
	}
}
