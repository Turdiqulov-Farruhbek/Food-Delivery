package services

import (
	"context"

	ntf "Notification/genproto/notification"
	stor "Notification/storage"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserNotificationService struct {
	Store stor.StorageInterface
	ntf.UnimplementedUserNotificationServiceServer
}

func NewUserNotificationService(stor stor.StorageInterface) *UserNotificationService {
	return &UserNotificationService{Store: stor}
}

func (s *UserNotificationService) CreateUserNotification(ctx context.Context, req *ntf.CreateUserNotificationRequest) (*ntf.UserNotificationResponse, error) {
	resNtf, err := s.Store.UserNtf().CreateUserNtf(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user notification: %v", err)
	}

	return &ntf.UserNotificationResponse{
		Success: true,
		Message: "User notification created successfully",
		Notification: &ntf.UserNotification{
			NotificationId: resNtf.Notification.NotificationId,
			UserId:         req.UserId,
			OrderId:        req.OrderId,
			Type:           req.Type,
			Message:        req.Message,
		},
	}, nil
}

func (s *UserNotificationService) GetUserNotification(ctx context.Context, req *ntf.UserNotificationRequest) (*ntf.UserNotificationResponse, error) {
	resNtf, err := s.Store.UserNtf().GetUserNtf(ctx, req)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user notification not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get user notification: %v", err)
	}

	return &ntf.UserNotificationResponse{
		Success: true,
		Message: "User notification retrieved successfully",
		Notification: &ntf.UserNotification{
			NotificationId: resNtf.Notification.NotificationId,
			UserId:         resNtf.Notification.UserId,
			OrderId:        resNtf.Notification.OrderId,
			Type:           resNtf.Notification.Type,
			Message:        resNtf.Notification.Message,
			CreatedAt:      resNtf.Notification.CreatedAt,
			IsRead:         resNtf.Notification.IsRead,
		},
	}, nil
}

func (s *UserNotificationService) DeleteUserNotification(ctx context.Context, req *ntf.UserNotificationRequest) (*ntf.UserNotificationResponse, error) {
	resp, err := s.Store.UserNtf().DeleteUserNtf(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user notification: %v", err)
	}

	return &ntf.UserNotificationResponse{
		Success: true,
		Message: "User notification deleted successfully",
		Notification: &ntf.UserNotification{
			NotificationId: resp.Notification.NotificationId,
			UserId:         resp.Notification.UserId,
			OrderId:        resp.Notification.OrderId,
			Type:           resp.Notification.Type,
			Message:        resp.Notification.Message,
		},
	}, nil
}

func (s *UserNotificationService) ListUserNotifications(ctx context.Context, req *ntf.Empty) (*ntf.UserNotificationListResponse, error) {
	resNtf, err := s.Store.UserNtf().ListUserNtf(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list user notifications: %v", err)
	}

	return &ntf.UserNotificationListResponse{
		Notifications: resNtf.Notifications,
	}, nil
}
