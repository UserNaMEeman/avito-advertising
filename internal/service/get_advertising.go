package service

import (
	"advertising/entity"
	"advertising/internal/repository"
	"context"
)

type ProductOne struct {
	repo repository.GetAdt
}

func NewProduct(repo repository.GetAdt) *ProductOne {
	return &ProductOne{repo: repo}
}

// (ctx context.Context, id int, fields entity.ResponseFileds) (entity.Product, error)
func (prod *ProductOne) GetProduct(ctx context.Context, id int,
	fields entity.ResponseFileds) (entity.Product, error) {
	product, err := prod.repo.GetProductDB(ctx, id, fields)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}
