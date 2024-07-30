package postgres

import (
	"context"
	"errors"

	ntf "Notification/genproto/notification"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CourierNtf struct {
	Db *pgx.Conn
}

func NewCourierNtf(db *pgx.Conn) *CourierNtf {
	return &CourierNtf{Db: db}
}

func (c *CourierNtf) CreateCourierNtf(ctx context.Context, req *ntf.CreateCourierNotificationRequest) (*ntf.CourierNotificationResponse, error) {
	notificationID := uuid.New().String()

	query := `
		INSERT INTO couriernotifications (notification_id, courier_id, order_id, type, message)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING notification_id, courier_id, order_id, type, message
	`

	row := c.Db.QueryRow(ctx, query, notificationID, req.CourierId, req.OrderId, req.Type, req.Message)
	var notification ntf.CourierNotification
	err := row.Scan(&notification.NotificationId, &notification.CourierId, &notification.OrderId, &notification.Type, &notification.Message, &notification.CreatedAt, &notification.IsRead)
	if err != nil {
		return nil, err
	}

	return &ntf.CourierNotificationResponse{
		Success: true,
		Message: "Courier notification marked as deleted successfully",
	}, nil
}

func (c *CourierNtf) GetCourierNtf(ctx context.Context, req *ntf.CourierNotificationRequest) (*ntf.CourierNotificationResponse, error) {
	query := `
		SELECT notification_id, courier_id, order_id, type, message
		FROM couriernotifications
		WHERE notification_id = $1
	`

	row := c.Db.QueryRow(ctx, query, req.NotificationId)
	var notification ntf.CourierNotification
	err := row.Scan(&notification.NotificationId, &notification.CourierId, &notification.OrderId, &notification.Type, &notification.Message, &notification.CreatedAt, &notification.IsRead)
	if err != nil {
		return nil, err
	}

	return &ntf.CourierNotificationResponse{
		Success: true,
		Message: "Courier notification marked as deleted successfully",
	}, nil
}

func (c *CourierNtf) DeleteCourierNtf(ctx context.Context, req *ntf.CourierNotificationRequest) (*ntf.CourierNotificationResponse, error) {
	query := `UPDATE couriernotifications SET deleted_at=EXTRACT(EPOCH FROM NOW())::BIGINT WHERE notification_id=$1 AND deleted_at = 0`
	result, err := c.Db.Exec(ctx, query, req.NotificationId)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, errors.New("courier notification not found or already deleted")
	}
	return &ntf.CourierNotificationResponse{
		Success: true,
		Message: "Courier notification marked as deleted successfully",
	}, nil
}

func (c *CourierNtf) ListCourierNtf(ctx context.Context, req *ntf.Empty) (*ntf.CourierNotificationListResponse, error) {
	query := `
		SELECT notification_id, courier_id, order_id, type, message
		FROM couriernotifications
	`

	rows, err := c.Db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*ntf.CourierNotification
	for rows.Next() {
		var notification ntf.CourierNotification
		err := rows.Scan(&notification.NotificationId, &notification.CourierId, &notification.OrderId, &notification.Type, &notification.Message, &notification.CreatedAt, &notification.IsRead)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &ntf.CourierNotificationListResponse{
		Notifications: notifications,
	}, nil
}
