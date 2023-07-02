package product

import (
	"product/internal/models/product"
)

type Usecase interface {
	GetProduct(product.GetProductRequest) (result product.GetProductResponse, err error)
}
