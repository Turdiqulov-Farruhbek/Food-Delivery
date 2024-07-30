package services

import (
	"context"

	ntf "Notification/genproto/notification"
	stor "Notification/storage"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CourierNotificationService struct {
	Store stor.StorageInterface
	ntf.UnimplementedCourierNotificationServiceServer
}

func NewCourierNotificationService(stor stor.StorageInterface) *CourierNotificationService {
	return &CourierNotificationService{Store: stor}
}

func (s *CourierNotificationService) CreateCourierNotification(ctx context.Context, req *ntf.CreateCourierNotificationRequest) (*ntf.CourierNotificationResponse, error) {
	resNtf, err := s.Store.CourierNtf().CreateCourierNtf(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create courier notification: %v", err)
	}

	return &ntf.CourierNotificationResponse{
		Success: true,
		Message: "Courier notification created successfully",
		Notification: &ntf.CourierNotification{
			NotificationId: resNtf.Notification.NotificationId,
			CourierId:      resNtf.Notification.CourierId,
			OrderId:        resNtf.Notification.OrderId,
			Type:           resNtf.Notification.Type,
			Message:        resNtf.Notification.Message,
			CreatedAt:      resNtf.Notification.CreatedAt,
			IsRead:         resNtf.Notification.IsRead,
		},
	}, nil
}

func (s *CourierNotificationService) GetCourierNotification(ctx context.Context, req *ntf.CourierNotificationRequest) (*ntf.CourierNotificationResponse, error) {
	resNtf, err := s.Store.CourierNtf().GetCourierNtf(ctx, req)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "courier notification not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get courier notification: %v", err)
	}

	return &ntf.CourierNotificationResponse{
		Success: true,
		Message: "Courier notification retrieved successfully",
		Notification: &ntf.CourierNotification{
			NotificationId: resNtf.Notification.NotificationId,
			CourierId:      resNtf.Notification.CourierId,
			OrderId:        resNtf.Notification.OrderId,
			Type:           resNtf.Notification.Type,
			Message:        resNtf.Notification.Message,
			CreatedAt:      resNtf.Notification.CreatedAt,
			IsRead:         resNtf.Notification.IsRead,
		},
	}, nil
}

func (s *CourierNotificationService) DeleteCourierNotification(ctx context.Context, req *ntf.CourierNotificationRequest) (*ntf.CourierNotificationResponse, error) {
	resNtf, err := s.Store.CourierNtf().DeleteCourierNtf(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete courier notification: %v", err)
	}

	return &ntf.CourierNotificationResponse{
		Success: true,
		Message: "Courier notification retrieved successfully",
		Notification: &ntf.CourierNotification{
			NotificationId: resNtf.Notification.NotificationId,
			CourierId:      resNtf.Notification.CourierId,
			OrderId:        resNtf.Notification.OrderId,
			Type:           resNtf.Notification.Type,
			Message:        resNtf.Notification.Message,
			CreatedAt:      resNtf.Notification.CreatedAt,
			IsRead:         resNtf.Notification.IsRead,
		},
	}, nil
}

func (s *CourierNotificationService) ListCourierNotifications(ctx context.Context, req *ntf.Empty) (*ntf.CourierNotificationListResponse, error) {
	resNtf, err := s.Store.CourierNtf().ListCourierNtf(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list courier notifications: %v", err)
	}

	return &ntf.CourierNotificationListResponse{
		Notifications: resNtf.Notifications,
	}, nil
}
