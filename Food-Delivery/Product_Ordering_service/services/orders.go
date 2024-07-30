package service

import (
	"context"
	order "product_ordering/genproto/product"
	"product_ordering/storage"
)

// OrderService RPC xizmatini taqdim etadi
type OrderService struct {
	store storage.StorageInterface
	order.UnimplementedOrderProductServiceServer
}

// NewOrderService yangi OrderService ni yaratadi
func NewOrderService(store storage.StorageInterface) *OrderService {
	return &OrderService{store: store}
}

// CreateOrder RPC chaqiruvini bajaradi va buyurtma yaratadi
func (s *OrderService) CreateOrder(ctx context.Context, req *order.CreateOrderProductRequest) (*order.OrderProductResponse, error) {
	return s.store.Order().Create(ctx, req)
}

// GetOrder RPC chaqiruvini bajaradi va buyurtma ma'lumotlarini qaytaradi
func (s *OrderService) GetOrder(ctx context.Context, req *order.OrderProductRequest) (*order.OrderProductResponse, error) {
	return s.store.Order().Get(ctx, req)
}

// UpdateOrder RPC chaqiruvini bajaradi va buyurtmani yangilaydi
func (s *OrderService) UpdateOrder(ctx context.Context, req *order.UpdateOrderProductRequest) (*order.OrderProductResponse, error) {
	return s.store.Order().Update(ctx, req)
}

// DeleteOrder RPC chaqiruvini bajaradi va buyurtmani o'chiradi
func (s *OrderService) DeleteOrder(ctx context.Context, req *order.OrderProductRequest) (*order.OrderProductResponse, error) {
	return s.store.Order().Delete(ctx, req)
}

// ListOrders RPC chaqiruvini bajaradi va buyurtmalar ro'yxatini qaytaradi
func (s *OrderService) ListOrders(ctx context.Context, req *order.OrderProductListRequest) (*order.OrderProductListResponse, error) {
	return s.store.Order().List(ctx, req)
}
