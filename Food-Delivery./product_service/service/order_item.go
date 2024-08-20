package service

import (
	"context"

	pb "delivery/product_service/genproto"
	"delivery/product_service/storage"
)

type OrderItemService struct {
	stg storage.StorageI
	pb.UnimplementedItemServiceServer
}

func NewOrderItemService(stg storage.StorageI) *OrderItemService {
	return &OrderItemService{stg: stg}
}

func (s *OrderItemService) AddItem(ctx context.Context, req *pb.ItemCreateReq) (*pb.Void, error) {
	_, err := s.stg.OrderItem().AddItem(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *OrderItemService) GetItem(ctx context.Context, req *pb.ById) (*pb.ItemGet, error) {
	return s.stg.OrderItem().GetItem(req)
}

func (s *OrderItemService) UpdateItem(ctx context.Context, req *pb.ItemCreateReq) (*pb.Void, error) {
	_, err := s.stg.OrderItem().UpdateItem(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *OrderItemService) DeleteItem(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	_, err := s.stg.OrderItem().DeleteItem(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *OrderItemService) ListItems(ctx context.Context, req *pb.ItemFilter) (*pb.ItemList, error) {
	return s.stg.OrderItem().ListItems(req)
}
