package usecase

import (
	"product/config"
	models "product/internal/models/product"
	"product/internal/product"
)

type ProductUsecase struct {
	cfg *config.Config
}

func NewProducetUsecase(cfg *config.Config) product.Usecase {
	return &ProductUsecase{cfg: cfg}
}

func (u *ProductUsecase) GetProduct(request models.GetProductRequest) (result models.GetProductResponse, err error) {
	return models.GetProductResponse{ID: request.ID, Name: "Best Product", Price: 10}, nil
}
