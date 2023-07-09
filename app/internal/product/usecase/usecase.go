package usecase

import (
	"context"
	"fmt"
	"product/config"
	models "product/internal/models/product"
	"product/internal/product"
	pb "product/pkg/proto"
	"sync"
	"sync/atomic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	count int64
)

type ProductUsecase struct {
	cfg        *config.Config
	productMap sync.Map
}

func NewProducetUsecase(cfg *config.Config) product.Usecase {
	return &ProductUsecase{cfg: cfg}
}

func (u *ProductUsecase) GetProduct(request models.GetProductRequest) (result models.GetProductResponse, err error) {
	return models.GetProductResponse{ID: request.ID, Name: "Best of the best Product", Price: 10}, nil
}

func (u *ProductUsecase) AddProductByGRPC(
	ctx context.Context, req *pb.AddProductRequest) (res *pb.AddProductResponse, err error) {

	atomic.AddInt64(&count, 1)
	u.productMap.Store(count, req)
	return &pb.AddProductResponse{ID: int64(count)}, nil
}

func (u *ProductUsecase) GetProductByGRPC(
	ctx context.Context, req *pb.GetProductRequest) (res *pb.GetProductResponse, err error) {

	product, ok := u.productMap.Load(req.GetID())
	if !ok {
		return &pb.GetProductResponse{},
			status.Error(codes.NotFound, fmt.Sprintf("Product with ID: %d does not exists", req.GetID()))
	}

	result, ok := product.(*pb.AddProductRequest)
	if !ok {
		return &pb.GetProductResponse{},
			status.Error(codes.Internal, fmt.Sprintf("Unable to cast value(%+v) to *pb.GetProductResponse", product))
	}

	return &pb.GetProductResponse{Name: result.GetName(), Price: result.GetPrice()}, nil
}
