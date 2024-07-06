package service

import (
	"context"
	order "product_ordering/genproto"
	"product_ordering/storage"
)

// OrderService RPC xizmatini taqdim etadi
type OrderService struct {
	store storage.StorageInterface
	order.UnimplementedOrderServiceServer
}

// NewOrderService yangi OrderService ni yaratadi
func NewOrderService(store storage.StorageInterface) *OrderService {
	return &OrderService{store: store}
}

// CreateOrder RPC chaqiruvini bajaradi va buyurtma yaratadi
func (s *OrderService) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.OrderResponse, error) {
	return s.store.Order().Create(ctx, req)
}

// GetOrder RPC chaqiruvini bajaradi va buyurtma ma'lumotlarini qaytaradi
func (s *OrderService) GetOrder(ctx context.Context, req *order.OrderRequest) (*order.OrderResponse, error) {
	return s.store.Order().Get(ctx, req)
}

// UpdateOrder RPC chaqiruvini bajaradi va buyurtmani yangilaydi
func (s *OrderService) UpdateOrder(ctx context.Context, req *order.UpdateOrderRequest) (*order.OrderResponse, error) {
	return s.store.Order().Update(ctx, req)
}

// DeleteOrder RPC chaqiruvini bajaradi va buyurtmani o'chiradi
func (s *OrderService) DeleteOrder(ctx context.Context, req *order.OrderRequest) (*order.OrderResponse, error) {
	return s.store.Order().Delete(ctx, req)
}

// ListOrders RPC chaqiruvini bajaradi va buyurtmalar ro'yxatini qaytaradi
func (s *OrderService) ListOrders(ctx context.Context, req *order.OrderListRequest) (*order.OrderListResponse, error) {
	return s.store.Order().List(ctx, req)
}
