package storage

import (
	"context"
	ntf "Notification/genproto"
)

type StorageInterface interface {
	AdminAlirt() AdminAlirtInterface
	CourierNtf() CourierNtfInterface
	UserNtf() UserNtfInterface
}

type AdminAlirtInterface interface {
	CreateAdminAlert(ctx context.Context, req *ntf.CreateAdminAlertRequest) (*ntf.AdminAlertResponse, error)
	GetAdminAlert(ctx context.Context, req *ntf.AdminAlertRequest) (*ntf.AdminAlertResponse, error)
	DeleteAdminAlert(ctx context.Context, req *ntf.AdminAlertRequest) (*ntf.AdminAlertResponse, error)
	ListAdminAlerts(ctx context.Context, req *ntf.Empty) (*ntf.AdminAlertListResponse, error)
}


type CourierNtfInterface interface {
    CreateCourierNtf(ctx context.Context, req *ntf.CreateCourierNotificationRequest) (*ntf.CourierNotificationResponse, error)
    GetCourierNtf(ctx context.Context, req *ntf.CourierNotificationRequest) (*ntf.CourierNotificationResponse, error)
    DeleteCourierNtf(ctx context.Context, req *ntf.CourierNotificationRequest) (*ntf.CourierNotificationResponse, error)
	ListCourierNtf(ctx context.Context, req *ntf.Empty) (*ntf.CourierNotificationListResponse, error)
}


type UserNtfInterface interface {
    CreateUserNtf(ctx context.Context, req *ntf.CreateUserNotificationRequest) (*ntf.UserNotificationResponse, error)
    GetUserNtf(ctx context.Context, req *ntf.UserNotificationRequest) (*ntf.UserNotificationResponse, error)
    DeleteUserNtf(ctx context.Context, req *ntf.UserNotificationRequest) (*ntf.UserNotificationResponse, error)
	ListUserNtf(ctx context.Context, req *ntf.Empty) (*ntf.UserNotificationListResponse, error)
}
