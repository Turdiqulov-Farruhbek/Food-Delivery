package courier

import (
	"context"

	"courier_delivery/config/logger"
	"courier_delivery/genproto/courier"
	"courier_delivery/storage"
)

// CourierLocationService kuryer lokatsiyalari uchun gRPC xizmatini taqdim etadi
type CourierLocationService struct {
	store storage.StorageCourierInterface
	courier.UnimplementedCourierLocationServiceServer
	log logger.Logger
}

// NewCourierLocationService yangi CourierLocationService yaratadi
func NewCourierLocationService(store storage.StorageCourierInterface, log logger.Logger) *CourierLocationService {
	return &CourierLocationService{store: store, log: log}
}

// CreateCourierLocation RPC chaqiruvini bajaradi va yangi kuryer lokatsiyasini yaratadi
func (s *CourierLocationService) CreateCourierLocation(ctx context.Context, req *courier.CreateCourierLocationRequest) (*courier.CourierLocationResponse, error) {
	res, err := s.store.CourierLocation().CreateCourierLocation(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to create courier location: %v", err)
		return nil, err
	}
	s.log.INFO.Println("Successfully created courier location:", res.Id)
	return res, nil
}

// GetCourierLocation RPC chaqiruvini bajaradi va kuryer lokatsiyasini qaytaradi
func (s *CourierLocationService) GetCourierLocation(ctx context.Context, req *courier.GetCourierLocationRequest) (*courier.CourierLocationResponse, error) {
	res, err := s.store.CourierLocation().GetCourierLocation(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to get courier location with ID %s: %v", req.Id, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully retrieved courier location:", res.Id)
	return res, nil
}

// UpdateCourierLocation RPC chaqiruvini bajaradi va kuryer lokatsiyasini yangilaydi
func (s *CourierLocationService) UpdateCourierLocation(ctx context.Context, req *courier.UpdateCourierLocationRequest) (*courier.CourierLocationResponse, error) {
	res, err := s.store.CourierLocation().UpdateCourierLocation(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to update courier location with ID %s: %v", req.Id, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully updated courier location:", res.Id)
	return res, nil
}

// DeleteCourierLocation RPC chaqiruvini bajaradi va kuryer lokatsiyasini o'chiradi
func (s *CourierLocationService) DeleteCourierLocation(ctx context.Context, req *courier.DeleteCourierLocationRequest) (*courier.DeleteCourierLocationResponse, error) {
	res, err := s.store.CourierLocation().DeleteCourierLocation(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to delete courier location with ID %s: %v", req.Id, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully deleted courier location:", req.Id)
	return res, nil
}

// GetAllCourierLocations RPC chaqiruvini bajaradi va barcha kuryer lokatsiyalarini qaytaradi
func (s *CourierLocationService) GetAllCourierLocations(ctx context.Context, req *courier.GetAllCourierLocationsRequest) (*courier.GetAllCourierLocationsResponse, error) {
	res, err := s.store.CourierLocation().GetAllCourierLocations(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to get all courier locations: %v", err)
		return nil, err
	}
	s.log.INFO.Println("Successfully retrieved all courier locations")
	return res, nil
}
