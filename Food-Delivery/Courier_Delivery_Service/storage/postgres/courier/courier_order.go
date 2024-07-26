package courier

import (
	"context"
	"courier_delivery/genproto"
	"errors"
	"fmt"
	"time" // time paketini import qilish

	"github.com/jackc/pgx/v5"
)

type CourierOrder struct {
	Db *pgx.Conn
}

func NewCourierOrder(db *pgx.Conn) *CourierOrder {
	return &CourierOrder{Db: db}
}

// Create kuryer-buyurtma munosabatini yaratadi
func (c *CourierOrder) CreateCourierOrder(ctx context.Context, req *genproto.CreateCourierOrderRequest) (*genproto.CourierOrderResponse, error) {
	query := `INSERT INTO courierorders (courier_id, order_id, status, assigned_time, last_updated) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING courier_order_id`
	var courierOrderID string
	err := c.Db.QueryRow(ctx, query, req.CourierId, req.OrderId, req.Status).Scan(&courierOrderID)
	if err != nil {
		return nil, fmt.Errorf("failed to create courier order: %w", err)
	}
	return &genproto.CourierOrderResponse{
		Success: true,
		Message: "Courier order created successfully",
		CourierOrder: &genproto.CourierOrder{
			CourierOrderId: courierOrderID,
			CourierId:      req.CourierId,
			OrderId:        req.OrderId,
			Status:         req.Status,
			AssignedTime:   time.Now().Format(time.RFC3339), // Vaqtni RFC3339 formatida saqlash
			LastUpdated:    time.Now().Format(time.RFC3339), // Vaqtni RFC3339 formatida saqlash
		},
	}, nil
}

// Get kuryer-buyurtma munosabati ma'lumotlarini qaytaradi
func (c *CourierOrder) GetCourierOrder(ctx context.Context, req *genproto.CourierOrderRequest) (*genproto.CourierOrderResponse, error) {
	query := `SELECT courier_order_id, courier_id, order_id, status, assigned_time, last_updated FROM courierorders WHERE courier_order_id=$1 AND deleted_at=0`
	var courierOrder genproto.CourierOrder
	err := c.Db.QueryRow(ctx, query, req.CourierOrderId).Scan(
		&courierOrder.CourierOrderId,
		&courierOrder.CourierId,
		&courierOrder.OrderId,
		&courierOrder.Status,
		&courierOrder.AssignedTime,
		&courierOrder.LastUpdated)
	if err != nil {
		return nil, fmt.Errorf("failed to get courier order: %w", err)
	}
	return &genproto.CourierOrderResponse{
		Success:      true,
		Message:      "Courier order retrieved successfully",
		CourierOrder: &courierOrder,
	}, nil
}

// Update kuryer-buyurtma munosabatini yangilaydi
func (c *CourierOrder) UpdateCourierOrder(ctx context.Context, req *genproto.UpdateCourierOrderRequest) (*genproto.CourierOrderResponse, error) {
	// Avval eski qiymatlarni olish
	var oldCourierId, oldOrderId, oldAssignedTime string
	var oldStatus genproto.CourierOrderStatus

	getQuery := `SELECT courier_id, order_id, status, assigned_time FROM courierorders WHERE courier_order_id=$1`
	err := c.Db.QueryRow(ctx, getQuery, req.CourierOrderId).Scan(&oldCourierId, &oldOrderId, &oldStatus, &oldAssignedTime)
	if err != nil {
		return nil, fmt.Errorf("failed to get existing courier order: %w", err)
	}

	// Yangi qiymatlar bo'sh bo'lsa, eski qiymatlarni ishlatish
	courierId := req.CourierId
	if courierId == "" {
		courierId = oldCourierId
	}

	orderId := req.OrderId
	if orderId == "" {
		orderId = oldOrderId
	}

	status := req.Status
	if req.Status == genproto.CourierOrderStatus(0) { // 0 status qiymati ASSIGNED sifatida ko'riladi
		status = oldStatus
	}

	assignedTime := req.AssignedTime
	if assignedTime == "" {
		assignedTime = oldAssignedTime
	}

	// Yangilanish so'rovi
	updateQuery := `UPDATE courierorders SET courier_id=$1, order_id=$2, status=$3, assigned_time=$4, last_updated=NOW() WHERE courier_order_id=$5 AND deleted_at=0`
	_, err = c.Db.Exec(ctx, updateQuery, courierId, orderId, status, assignedTime, req.CourierOrderId)
	if err != nil {
		return nil, fmt.Errorf("failed to update courier order: %w", err)
	}

	return &genproto.CourierOrderResponse{
		Success: true,
		Message: "Courier order updated successfully",
		CourierOrder: &genproto.CourierOrder{
			CourierOrderId: req.CourierOrderId,
			CourierId:      courierId,
			OrderId:        orderId,
			Status:         status,
			AssignedTime:   assignedTime,
			LastUpdated:    time.Now().Format(time.RFC3339), // Vaqtni RFC3339 formatida saqlash
		},
	}, nil
}

// Delete kuryer-buyurtma munosabatini mantiqiy o'chiradi
func (c *CourierOrder) DeleteCourierOrder(ctx context.Context, req *genproto.CourierOrderRequest) (*genproto.CourierOrderResponse, error) {
	// Mantiqiy o'chirish: deleted_at ustunini yangilash
	query := `UPDATE courierorders SET deleted_at=EXTRACT(EPOCH FROM NOW())::BIGINT WHERE courier_order_id=$1 AND deleted_at=0`
	result, err := c.Db.Exec(ctx, query, req.CourierOrderId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete courier order: %w", err)
	}

	// Agar hech qator yangilanmagan bo'lsa, yozuv topilmagan
	if result.RowsAffected() == 0 {
		return nil, errors.New("courier order not found or already deleted")
	}

	return &genproto.CourierOrderResponse{
		Success: true,
		Message: "Courier order deleted successfully",
	}, nil
}

// List barcha kuryer-buyurtma munosabatlari ro'yxatini qaytaradi
func (c *CourierOrder) ListCourierOrders(ctx context.Context, req *genproto.Empty) (*genproto.CourierOrderListResponse, error) {
	query := `SELECT courier_order_id, courier_id, order_id, status, assigned_time, last_updated FROM courierorders WHERE deleted_at=0`
	rows, err := c.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list courier orders: %w", err)
	}
	defer rows.Close()

	var courierOrders []*genproto.CourierOrder
	for rows.Next() {
		var courierOrder genproto.CourierOrder
		err := rows.Scan(
			&courierOrder.CourierOrderId,
			&courierOrder.CourierId,
			&courierOrder.OrderId,
			&courierOrder.Status,
			&courierOrder.AssignedTime,
			&courierOrder.LastUpdated)
		if err != nil {
			return nil, fmt.Errorf("failed to scan courier order: %w", err)
		}
		courierOrders = append(courierOrders, &courierOrder)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to iterate over courier orders: %w", rows.Err())
	}

	return &genproto.CourierOrderListResponse{
		CourierOrders: courierOrders,
	}, nil
}
