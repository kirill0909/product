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

func (h *productHandler) AddProduct(
	ctx context.Context, req *pb.AddProductRequest) (res *pb.AddProductResponse, err error) {
	return h.productUC.AddProductByGRPC(ctx, req)
}

func (h *productHandler) GetProduct(
	ctx context.Context, req *pb.GetProductRequest) (res *pb.GetProductResponse, err error) {
	return h.productUC.GetProductByGRPC(ctx, req)
}

func (h *productHandler) GetProductsByPrice(
	req *pb.GetProductsByPriceRequest, stream pb.Product_GetProductsByPriceServer) error {
	return h.productUC.GetProductsByPrice(req, stream)
}
