package service

import (
	"context"

	"courier_delivery/config/logger"
	"courier_delivery/genproto"
	"courier_delivery/storage"
)

// CourierService kuryerlar uchun gRPC xizmatini taqdim etadi
type CourierService struct {
	store storage.StorageInterface
	genproto.UnimplementedCourierServiceServer
	log logger.Logger
}

// NewCourierService yangi CourierService yaratadi
func NewCourierService(store storage.StorageInterface, log logger.Logger) *CourierService {
	return &CourierService{store: store, log: log}
}

// CreateCourier RPC chaqiruvini bajaradi va yangi kuryer yozuvini yaratadi
func (s *CourierService) CreateCourier(ctx context.Context, req *genproto.CreateCourierRequest) (*genproto.CourierResponse, error) {
	res, err := s.store.Courier().CreateCourier(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to create courier: %v", err)
		return nil, err
	}
	s.log.INFO.Println("Successfully created courier:", res.Courier.CourierId)
	return res, nil
}

// GetCourier RPC chaqiruvini bajaradi va kuryer yozuvini kuryer_id orqali qaytaradi
func (s *CourierService) GetCourier(ctx context.Context, req *genproto.CourierRequest) (*genproto.CourierResponse, error) {
	res, err := s.store.Courier().GetCourier(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to get courier with ID %s: %v", req.CourierId, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully retrieved courier:", res.Courier.CourierId)
	return res, nil
}

// UpdateCourier RPC chaqiruvini bajaradi va kuryer yozuvini yangilaydi
func (s *CourierService) UpdateCourier(ctx context.Context, req *genproto.UpdateCourierRequest) (*genproto.CourierResponse, error) {
	res, err := s.store.Courier().UpdateCourier(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to update courier with ID %s: %v", req.CourierId, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully updated courier:", res.Courier.CourierId)
	return res, nil
}

// DeleteCourier RPC chaqiruvini bajaradi va kuryer yozuvini mantiqiy o'chiradi
func (s *CourierService) DeleteCourier(ctx context.Context, req *genproto.CourierRequest) (*genproto.CourierResponse, error) {
	res, err := s.store.Courier().DeleteCourier(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to delete courier with ID %s: %v", req.CourierId, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully deleted courier:", req.CourierId)
	return res, nil
}

// ListCouriers RPC chaqiruvini bajaradi va barcha kuryer yozuvlarini qaytaradi
func (s *CourierService) ListCouriers(ctx context.Context, req *genproto.Empty) (*genproto.CourierListResponse, error) {
	res, err := s.store.Courier().ListCouriers(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to list couriers: %v", err)
		return nil, err
	}
	s.log.INFO.Println("Successfully listed couriers:", len(res.Couriers))
	return res, nil
}
