package product

import (
	"context"
	"product/internal/models/product"
	pb "product/pkg/proto"
)

type Usecase interface {
	GetProduct(product.GetProductRequest) (result product.GetProductResponse, err error)
	AddProductByGRPC(ctx context.Context, req *pb.AddProductRequest) (res *pb.AddProductResponse, err error)
	GetProductByGRPC(ctx context.Context, req *pb.GetProductRequest) (res *pb.GetProductResponse, err error)
	GetProductsByPrice(req *pb.GetProductsByPriceRequest, stream pb.Product_GetProductsByPriceServer) error
}
