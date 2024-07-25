package postgres

import (
	"context"
	"errors"

	ntf "Notification/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserNtf struct {
	Db *pgx.Conn
}

func NewUserNtf(db *pgx.Conn) *UserNtf {
	return &UserNtf{Db: db}
}

func (u *UserNtf) CreateUserNtf(ctx context.Context, req *ntf.CreateUserNotificationRequest) (*ntf.UserNotificationResponse, error) {
	notificationID := uuid.New().String()

	query := `
		INSERT INTO usernotifications (notification_id, user_id, order_id, type, message)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING notification_id, user_id, order_id, type, message, created_at
	`

	row := u.Db.QueryRow(ctx, query, notificationID, req.UserId, req.OrderId, req.Type, req.Message)
	var notification ntf.UserNotification
	err := row.Scan(&notification.NotificationId, &notification.UserId, &notification.OrderId, &notification.Type, &notification.Message, &notification.CreatedAt, &notification.IsRead)
	if err != nil {
		return nil, err
	}

	return &ntf.UserNotificationResponse{
		Success: true,
		Message: "User notification marked as deleted successfully",
	}, nil
}

func (u *UserNtf) GetUserNtf(ctx context.Context, req *ntf.UserNotificationRequest) (*ntf.UserNotificationResponse, error) {
	query := `
		SELECT notification_id, user_id, order_id, type, message, created_at
		FROM usernotifications
		WHERE notification_id = $1
	`

	row := u.Db.QueryRow(ctx, query, req.NotificationId)
	var notification ntf.UserNotification
	err := row.Scan(&notification.NotificationId, &notification.UserId, &notification.OrderId, &notification.Type, &notification.Message, &notification.CreatedAt, &notification.IsRead)
	if err != nil {
		return nil, err
	}

	return &ntf.UserNotificationResponse{
		Success: true,
		Message: "User notification marked as deleted successfully",
	}, nil
}

func (u *UserNtf) DeleteUserNtf(ctx context.Context, req *ntf.UserNotificationRequest) (*ntf.UserNotificationResponse, error) {
	query := `UPDATE usernotifications SET deleted_at=EXTRACT(EPOCH FROM NOW())::BIGINT WHERE notification_id=$1 AND deleted_at = 0`
	result, err := u.Db.Exec(ctx, query, req.NotificationId)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, errors.New("user notification not found or already deleted")
	}
	return &ntf.UserNotificationResponse{
		Success: true,
		Message: "User notification marked as deleted successfully",
	}, nil
}

func (u *UserNtf) ListUserNtf(ctx context.Context, req *ntf.Empty) (*ntf.UserNotificationListResponse, error) {
	query := `
		SELECT notification_id, user_id, order_id, type, message, created_at, is_read
		FROM usernotifications
	`

	rows, err := u.Db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*ntf.UserNotification
	for rows.Next() {
		var notification ntf.UserNotification
		err := rows.Scan(&notification.NotificationId, &notification.UserId, &notification.OrderId, &notification.Type, &notification.Message, &notification.CreatedAt, &notification.IsRead)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &ntf.UserNotificationListResponse{
		Notifications: notifications,
	}, nil
}

