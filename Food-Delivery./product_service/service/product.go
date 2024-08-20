package service

import (
	"context"

	pb "delivery/product_service/genproto"
	"delivery/product_service/storage"
)

type ProductService struct {
	stg storage.StorageI
	pb.UnimplementedProductServiceServer
}

func NewProductService(stg storage.StorageI) *ProductService {
	return &ProductService{stg: stg}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.ProductCreateReq) (*pb.Void, error) {
	_, err := s.stg.Product().CreateProduct(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.ById) (*pb.ProductGet, error) {
	return s.stg.Product().GetProduct(req)
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.ProductCreateReq) (*pb.Void, error) {
	_, err := s.stg.Product().UpdateProduct(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	_, err := s.stg.Product().DeleteProduct(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *pb.ProductFilter) (*pb.ProductList, error) {
	return s.stg.Product().ListProducts(req)
}
