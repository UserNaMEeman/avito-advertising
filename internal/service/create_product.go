package service

import (
	"advertising/entity"
	"advertising/internal/repository"
	"context"
)

type CreateProduct struct {
	repo repository.CreateAdt
}

func NewCreateProduct(repo repository.CreateAdt) *CreateProduct {
	return &CreateProduct{repo: repo}
}

func (prod *CreateProduct) CreateProduct(ctx context.Context, product entity.Product) (int, error) {
	if len([]byte(product.Title)) > 200 || len([]byte(product.Description)) > 1000 {
		return 0, entity.ErrLongField
	}
	id, err := prod.repo.CreateProductDB(ctx, product)
	return id, err
}
