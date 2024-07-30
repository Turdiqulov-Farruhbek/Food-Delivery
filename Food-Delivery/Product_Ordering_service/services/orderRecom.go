package service

import (
	"context"

	orderRecommendation "product_ordering/genproto/product"
	"product_ordering/storage"
)

// OrderRecommendationService RPC xizmatini taqdim etadi
type OrderRecommendationService struct {
	store storage.StorageInterface
	orderRecommendation.UnimplementedOrderRecommendationServiceServer
}

// NewOrderRecommendationService yangi OrderRecommendationService ni yaratadi
func NewOrderRecommendationService(store storage.StorageInterface) *OrderRecommendationService {
	return &OrderRecommendationService{store: store}
}

// CreateOrderRecommendation RPC chaqiruvini bajaradi va buyurtma tavsiyasini yaratadi
func (s *OrderRecommendationService) CreateOrderRecommendation(ctx context.Context, req *orderRecommendation.CreateOrderRecommendationRequest) (*orderRecommendation.OrderRecommendationResponse, error) {
	return s.store.OrderRecomend().Create(ctx, req)
}

// GetOrderRecommendation RPC chaqiruvini bajaradi va buyurtma tavsiyasi ma'lumotlarini qaytaradi
func (s *OrderRecommendationService) GetOrderRecommendation(ctx context.Context, req *orderRecommendation.OrderRecommendationRequest) (*orderRecommendation.OrderRecommendationResponse, error) {
	return s.store.OrderRecomend().Get(ctx, req)
}

// UpdateOrderRecommendation RPC chaqiruvini bajaradi va buyurtma tavsiyasini yangilaydi
func (s *OrderRecommendationService) UpdateOrderRecommendation(ctx context.Context, req *orderRecommendation.UpdateOrderRecommendationRequest) (*orderRecommendation.OrderRecommendationResponse, error) {
	return s.store.OrderRecomend().Update(ctx, req)
}

// DeleteOrderRecommendation RPC chaqiruvini bajaradi va buyurtma tavsiyasini o'chiradi
func (s *OrderRecommendationService) DeleteOrderRecommendation(ctx context.Context, req *orderRecommendation.OrderRecommendationRequest) (*orderRecommendation.OrderRecommendationResponse, error) {
	return s.store.OrderRecomend().Delete(ctx, req)
}

// ListOrderRecommendations RPC chaqiruvini bajaradi va buyurtma tavsiyalari ro'yxatini qaytaradi
func (s *OrderRecommendationService) ListOrderRecommendations(ctx context.Context, req *orderRecommendation.Empty) (*orderRecommendation.OrderRecommendationListResponse, error) {
	return s.store.OrderRecomend().List(ctx, req)
}
