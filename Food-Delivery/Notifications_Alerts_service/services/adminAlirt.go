package services

import (
	"context"

	ntf "Notification/genproto"
	stor "Notification/storage"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AdminAlertService struct {
	Store stor.StorageInterface
	ntf.UnimplementedAdminAlertServiceServer
}

func NewAdminAlertService(stor stor.StorageInterface) *AdminAlertService {
	return &AdminAlertService{Store: stor}
}

func (s *AdminAlertService) CreateAdminAlert(ctx context.Context, req *ntf.CreateAdminAlertRequest) (*ntf.AdminAlertResponse, error) {
	alert, err := s.Store.AdminAlirt().CreateAdminAlert(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create admin alert: %v", err)
	}

	return &ntf.AdminAlertResponse{
		Success: true,
		Message: "Admin alert created successfully",
		Alert: &ntf.AdminAlert{
			AlertId:   alert.Alert.AdminId,
			AdminId:   alert.Alert.AdminId,
			Type:      alert.Alert.Type,
			Message:   alert.Alert.Message,
			CreatedAt: alert.Alert.CreatedAt,
			IsRead:    alert.Alert.IsRead,
		},
	}, nil
}

func (s *AdminAlertService) GetAdminAlert(ctx context.Context, req *ntf.AdminAlertRequest) (*ntf.AdminAlertResponse, error) {
	alert, err := s.Store.AdminAlirt().GetAdminAlert(ctx, req)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "admin alert not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get admin alert: %v", err)
	}

	return &ntf.AdminAlertResponse{
		Success: true,
		Message: "Admin alert retrieved successfully",
		Alert: &ntf.AdminAlert{
			AlertId:   alert.Alert.AdminId,
			AdminId:   alert.Alert.AdminId,
			Type:      alert.Alert.Type,
			Message:   alert.Alert.Message,
			CreatedAt: alert.Alert.CreatedAt,
			IsRead:    alert.Alert.IsRead,
		},
	}, nil
}

func (s *AdminAlertService) DeleteAdminAlert(ctx context.Context, req *ntf.AdminAlertRequest) (*ntf.AdminAlertResponse, error) {
	res, err := s.Store.AdminAlirt().DeleteAdminAlert(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete admin alert: %v", err)
	}

	return &ntf.AdminAlertResponse{
		Success: true,
		Message: "Admin alert retrieved successfully",
		Alert: &ntf.AdminAlert{
			AlertId:   res.Alert.AdminId,
			AdminId:   res.Alert.AdminId,
			Type:      res.Alert.Type,
			Message:   res.Alert.Message,
			CreatedAt: res.Alert.CreatedAt,
			IsRead:    res.Alert.IsRead,
		},
	}, nil
}

func (s *AdminAlertService) ListAdminAlerts(ctx context.Context, req *ntf.Empty) (*ntf.AdminAlertListResponse, error) {
	alerts, err := s.Store.AdminAlirt().ListAdminAlerts(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list admin alerts: %v", err)
	}

	return &ntf.AdminAlertListResponse{
		Alerts: alerts.Alerts,
	}, nil
}
