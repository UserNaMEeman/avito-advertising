package service

import (
	"advertising/entity"
	"advertising/internal/repository"
	"context"
)

type ProductList struct {
	repo repository.GetList
}

func NewProductList(repo repository.GetList) *ProductList {
	return &ProductList{repo: repo}
}

func (prod *ProductList) GetProducts(ctx context.Context, sort entity.SortProduct) ([][]entity.ResponseList, error) {
	var fieldSort string
	var typeSort string
	productPages := make([][]entity.ResponseList, 0)
	// offset := 5
	limit := 5
	if sort.Up {
		typeSort = "DESC"
	} else {
		typeSort = "ASC"
	}
	if sort.PriceSort {
		fieldSort = "price"
	} else {
		fieldSort = "date_create"
	}
	sortDB := entity.SortProductDB{
		Field: fieldSort,
		Type:  typeSort,
	}
	// fmt.Println(sortDB)
	for offset := 0; ; offset += 5 {
		products, err := prod.repo.GetProductsDB(ctx, sortDB, offset, limit)
		if err != nil {
			return nil, err
		}
		productPages = append(productPages, products)
		if len(products) < limit {
			break
		}
	}
	return productPages, nil
}
