package service

import (
	"advertising/internal/repository"
)

type ProductList struct {
	repo repository.GetList
}

func NewProductList(repo repository.GetList) *ProductList {
	return &ProductList{repo: repo}
}

// func (*ProductList) GetProducts(ctx context.Context, sort entity.SortProduct) ([]entity.Product, error) {
// 	var fieldSort string
// 	var typeSort string
// 	if sort.Up {
// 		typeSort = "DESC"
// 	} else {
// 		typeSort = "ASC"
// 	}
// 	if sort.PriceSort {
// 		fieldSort = "price"
// 	} else {
// 		fieldSort = "date_create"
// 	}
// 	return func() ([]entity.Product, error) {
// 		b := make([]entity.Product, 0)
// 		return b, nil
// 	}()
// }
