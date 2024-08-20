package service

import (
	"context"
	"fmt"

	pb "delivery/product_service/genproto"
	"delivery/product_service/storage"
)

type OrderService struct {
	stg storage.StorageI
	pb.UnimplementedOrderServiceServer
}

func NewOrderService(stg storage.StorageI) *OrderService {
	return &OrderService{stg: stg}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.OrderCreateReq) (*pb.Void, error) {
	_, err := s.stg.Order().CreateOrder(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *OrderService) GetOrder(ctx context.Context, req *pb.ById) (*pb.OrderGet, error) {
	return s.stg.Order().GetOrder(req)
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *pb.OrderUpdate) (*pb.Void, error) {
	_, err := s.stg.Order().UpdateOrder(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	_, err := s.stg.Order().DeleteOrder(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, req *pb.OrderFilter) (*pb.OrderList, error) {
	return s.stg.Order().ListOrders(req)
}

func (s *OrderService) AssignOrder(ctx context.Context, req *pb.OrderAssign) (*pb.Void, error) {
	_, err := s.stg.Order().AssignOrder(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (s *OrderService) ReviewOrder(ctx context.Context, id *pb.ById) (*pb.OrderReview, error) {
	order, err := s.stg.Order().GetOrder(id)
	if err != nil {
		return nil, fmt.Errorf("service -- order -- review order- geteing the order error: %w", err)
	}
	items, err := s.stg.OrderItem().ListItems(&pb.ItemFilter{
		OrderId: order.Id,
	})
	if err != nil {
		return nil, fmt.Errorf("service -- order -- review order- getting the items error: %w", err)
	}

	return &pb.OrderReview{
		Id:           order.Id,
		UserId:       order.UserId,
		CourierId:    order.CourierId,
		Status:       order.Status,
		Total:        order.Total,
		Address:      order.Address,
		CreatedAt:    order.CreatedAt,
		DeliveryTime: order.DeliveryTime,
		IsPaid:       order.IsPaid,
		Items:        items,
	}, nil
}
