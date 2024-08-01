package courier

import (
	"context"

	"courier_delivery/config/logger"
	"courier_delivery/genproto/courier"
	"courier_delivery/storage"
)

// CourierOrderService kuryer-buyurtma operatsiyalari uchun gRPC xizmatini taqdim etadi
type CourierOrderService struct {
	store storage.StorageCourierInterface
	courier.UnimplementedCourierOrderServiceServer
	log logger.Logger
}

// NewCourierOrderService yangi CourierOrderService yaratadi
func NewCourierOrderService(store storage.StorageCourierInterface, log logger.Logger) *CourierOrderService {
	return &CourierOrderService{store: store, log: log}
}

// CreateCourierOrder RPC chaqiruvini bajaradi va yangi kuryer-buyurtma munosabatini yaratadi
func (s *CourierOrderService) CreateCourierOrder(ctx context.Context, req *courier.CreateCourierOrderRequest) (*courier.CourierOrderResponse, error) {
	res, err := s.store.CourierOrder().CreateCourierOrder(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to create courier order: %v", err)
		return nil, err
	}
	s.log.INFO.Println("Successfully created courier order:", res.CourierOrder.CourierOrderId)
	return res, nil
}

// GetCourierOrder RPC chaqiruvini bajaradi va kuryer-buyurtma munosabati ma'lumotlarini qaytaradi
func (s *CourierOrderService) GetCourierOrder(ctx context.Context, req *courier.CourierOrderRequest) (*courier.CourierOrderResponse, error) {
	res, err := s.store.CourierOrder().GetCourierOrder(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to get courier order with ID %s: %v", req.CourierOrderId, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully retrieved courier order:", res.CourierOrder.CourierOrderId)
	return res, nil
}

// UpdateCourierOrder RPC chaqiruvini bajaradi va kuryer-buyurtma munosabatini yangilaydi
func (s *CourierOrderService) UpdateCourierOrder(ctx context.Context, req *courier.UpdateCourierOrderRequest) (*courier.CourierOrderResponse, error) {
	res, err := s.store.CourierOrder().UpdateCourierOrder(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to update courier order with ID %s: %v", req.CourierOrderId, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully updated courier order:", res.CourierOrder.CourierOrderId)
	return res, nil
}

// DeleteCourierOrder RPC chaqiruvini bajaradi va kuryer-buyurtma munosabatini o'chiradi (mantiqiy o'chirish)
func (s *CourierOrderService) DeleteCourierOrder(ctx context.Context, req *courier.CourierOrderRequest) (*courier.CourierOrderResponse, error) {
	res, err := s.store.CourierOrder().DeleteCourierOrder(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to delete courier order with ID %s: %v", req.CourierOrderId, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully deleted courier order:", req.CourierOrderId)
	return res, nil
}

// ListCourierOrders RPC chaqiruvini bajaradi va barcha kuryer-buyurtma munosabatlarini qaytaradi
func (s *CourierOrderService) ListCourierOrders(ctx context.Context, req *courier.EmptyCourierOrder) (*courier.CourierOrderListResponse, error) {
	res, err := s.store.CourierOrder().ListCourierOrders(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to list courier orders: %v", err)
		return nil, err
	}
	s.log.INFO.Println("Successfully listed courier orders:", len(res.CourierOrders))
	return res, nil
}
