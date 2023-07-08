package grpc

import (
	"context"
	"product/internal/product"
	pb "product/pkg/proto"
)

type productHandler struct {
	pb.UnimplementedProductServer
	productUC product.Usecase
}

func NewProductHandler(productUC product.Usecase) pb.ProductServer {
	return &productHandler{productUC: productUC}
}

func (h *productHandler) AddProduct(ctx context.Context, req *pb.AddProductRequest) (res *pb.AddProductResponse, err error) {
	return &pb.AddProductResponse{ID: 10}, nil
}
