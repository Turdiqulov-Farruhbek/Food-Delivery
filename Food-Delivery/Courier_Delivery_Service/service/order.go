package service

import (
	"context"

	"courier_delivery/config/logger"
	"courier_delivery/genproto"
	"courier_delivery/storage"
)

// OrderService buyurtmalar uchun gRPC xizmatini taqdim etadi
type OrderService struct {
	store storage.StorageInterface
	genproto.UnimplementedOrderServiceServer
	log logger.Logger
}

// NewOrderService yangi OrderService yaratadi
func NewOrderService(store storage.StorageInterface, log logger.Logger) *OrderService {
	return &OrderService{store: store, log: log}
}

// CreateOrder RPC chaqiruvini bajaradi va yangi buyurtma yozuvini yaratadi
func (s *OrderService) CreateOrder(ctx context.Context, req *genproto.CreateOrderRequest) (*genproto.OrderResponse, error) {
	res, err := s.store.Order().CreateOrder(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to create order: %v", err)
		return nil, err
	}
	s.log.INFO.Println("Successfully created order:", res.Order.OrderId)
	return res, nil
}

// GetOrder RPC chaqiruvini bajaradi va buyurtma yozuvini qaytaradi
func (s *OrderService) GetOrder(ctx context.Context, req *genproto.OrderRequest) (*genproto.OrderResponse, error) {
	res, err := s.store.Order().GetOrder(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to get order with ID %s: %v", req.OrderId, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully retrieved order:", res.Order.OrderId)
	return res, nil
}

// UpdateOrder RPC chaqiruvini bajaradi va buyurtma yozuvini yangilaydi
func (s *OrderService) UpdateOrder(ctx context.Context, req *genproto.UpdateOrderRequest) (*genproto.OrderResponse, error) {
	res, err := s.store.Order().UpdateOrder(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to update order with ID %s: %v", req.OrderId, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully updated order:", res.Order.OrderId)
	return res, nil
}

// DeleteOrder RPC chaqiruvini bajaradi va buyurtma yozuvini mantiqiy o'chiradi
func (s *OrderService) DeleteOrder(ctx context.Context, req *genproto.OrderRequest) (*genproto.OrderResponse, error) {
	res, err := s.store.Order().DeleteOrder(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to delete order with ID %s: %v", req.OrderId, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully deleted order:", req.OrderId)
	return res, nil
}

// ListOrders RPC chaqiruvini bajaradi va barcha buyurtmalar ro'yxatini qaytaradi
func (s *OrderService) ListOrders(ctx context.Context, req *genproto.Empty) (*genproto.OrderListResponse, error) {
	res, err := s.store.Order().ListOrders(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to list orders: %v", err)
		return nil, err
	}
	s.log.INFO.Println("Successfully listed orders:", len(res.Orders))
	return res, nil
}
