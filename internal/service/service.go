package service

import (
	"advertising/entity"
	"advertising/internal/repository"
	"context"
)

type GetList interface {
	GetProducts(ctx context.Context, sortProd entity.SortProduct) ([][]entity.ResponseList, error)
}

type GetAdt interface {
	GetProduct(ctx context.Context, id int, fields entity.ResponseFileds) (entity.Product, error)
}

type CreateAdt interface {
	CreateProduct(ctx context.Context, product entity.Product) (int, error)
}

type Service struct {
	GetList
	GetAdt
	CreateAdt
}

func NewServices(repos *repository.Repository) *Service {
	return &Service{
		GetList:   NewProductList(repos.GetList),
		CreateAdt: NewCreateProduct(repos.CreateAdt),
		GetAdt:    NewProduct(repos.GetAdt),
	}
}
