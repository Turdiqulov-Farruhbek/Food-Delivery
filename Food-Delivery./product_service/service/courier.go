package service

import (
	"context"

	pb "delivery/product_service/genproto"
	"delivery/product_service/storage"
)

type CourierService struct {
	stg storage.StorageI
	pb.UnimplementedCourierServiceServer
}

func NewCourierService(stg storage.StorageI) *CourierService {
	return &CourierService{stg: stg}
}

func (c *CourierService) AcceptOrder(ctx context.Context, req *pb.OrderAcept) (*pb.Void, error) {
	_, err := c.stg.Courier().AcceptOrder(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
