package service

import (
	"advertising/entity"
	"advertising/internal/repository"
	"context"
)

type GetList interface {
	GetProducts(context.Context, entity.SortProduct) ([][]entity.ResponseList, error)
}

type GetAdt interface{}

type CreateAdt interface {
	CreateProduct(context.Context, entity.Product) (int, error)
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
	}
}
