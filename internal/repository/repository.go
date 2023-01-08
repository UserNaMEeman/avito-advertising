package repository

import (
	"advertising/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

var tableProduct string = "products"

type GetList interface {
	GetProducts(context.Context, entity.SortProductDB) ([]entity.Product, error)
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
		// GetList:   NewProductListDB(db),
	}
}
