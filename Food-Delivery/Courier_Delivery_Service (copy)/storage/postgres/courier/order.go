package courier

import (
	"context"
	"courier_delivery/genproto/courier"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Order struct {
	Db *pgx.Conn
}

func NewOrder(db *pgx.Conn) *Order {
	return &Order{Db: db}
}

// CreateOrder yangi buyurtma yozuvini yaratadi
func (o *Order) CreateOrder(ctx context.Context, req *courier.CreateOrderRequest) (*courier.OrderResponse, error) {
	query := `INSERT INTO orders (customer_id, order_details, delivery_address, payment_status, order_date) VALUES ($1, $2, $3, $4, NOW()) RETURNING order_id, order_date`
	var orderID string
	var orderDate time.Time
	err := o.Db.QueryRow(ctx, query, req.CustomerId, req.OrderDetails, req.DeliveryAddress, req.PaymentStatus).Scan(&orderID, &orderDate)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}
	return &courier.OrderResponse{
		Success: true,
		Message: "Order created successfully",
		Order: &courier.Order{
			OrderId:         orderID,
			CustomerId:      req.CustomerId,
			OrderDetails:    req.OrderDetails,
			DeliveryAddress: req.DeliveryAddress,
			PaymentStatus:   req.PaymentStatus,
			OrderDate:       orderDate.Format(time.RFC3339),
		},
	}, nil
}

// GetOrder buyurtma yozuvini qaytaradi
func (o *Order) GetOrder(ctx context.Context, req *courier.OrderRequest) (*courier.OrderResponse, error) {
	query := `SELECT order_id, customer_id, order_details, delivery_address, payment_status, order_date FROM orders WHERE order_id=$1`
	var order courier.Order
	err := o.Db.QueryRow(ctx, query, req.OrderId).Scan(
		&order.OrderId,
		&order.CustomerId,
		&order.OrderDetails,
		&order.DeliveryAddress,
		&order.PaymentStatus,
		&order.OrderDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}
	return &courier.OrderResponse{
		Success: true,
		Message: "Order retrieved successfully",
		Order:   &order,
	}, nil
}

// UpdateOrder buyurtma yozuvini yangilaydi
func (o *Order) UpdateOrder(ctx context.Context, req *courier.UpdateOrderRequest) (*courier.OrderResponse, error) {
	query := `UPDATE orders SET customer_id=$1, order_details=$2, delivery_address=$3, payment_status=$4, updated_at=NOW() WHERE order_id=$5 RETURNING order_id, customer_id, order_details, delivery_address, payment_status, order_date`
	var order courier.Order
	err := o.Db.QueryRow(ctx, query, req.CustomerId, req.OrderDetails, req.DeliveryAddress, req.PaymentStatus, req.OrderId).Scan(
		&order.OrderId,
		&order.CustomerId,
		&order.OrderDetails,
		&order.DeliveryAddress,
		&order.PaymentStatus,
		&order.OrderDate)
	if err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}
	return &courier.OrderResponse{
		Success: true,
		Message: "Order updated successfully",
		Order:   &order,
	}, nil
}

// DeleteOrder buyurtma yozuvini o'chiradi
func (o *Order) DeleteOrder(ctx context.Context, req *courier.OrderRequest) (*courier.OrderResponse, error) {
	query := `DELETE FROM orders WHERE order_id=$1`
	result, err := o.Db.Exec(ctx, query, req.OrderId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete order: %w", err)
	}

	if result.RowsAffected() == 0 {
		return nil, fmt.Errorf("order not found")
	}

	return &courier.OrderResponse{
		Success: true,
		Message: "Order deleted successfully",
	}, nil
}

// ListOrders barcha buyurtma yozuvlarini qaytaradi
func (o *Order) ListOrders(ctx context.Context, req *courier.GetRecommendedOrdersRequest) (*courier.OrderListResponse, error) {
	query := `SELECT order_id, customer_id, order_details, delivery_address, payment_status, order_date FROM orders`
	rows, err := o.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list orders: %w", err)
	}
	defer rows.Close()

	var orders []*courier.Order
	for rows.Next() {
		var order courier.Order
		err := rows.Scan(
			&order.OrderId,
			&order.CustomerId,
			&order.OrderDetails,
			&order.DeliveryAddress,
			&order.PaymentStatus,
			&order.OrderDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}
		orders = append(orders, &order)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to iterate over orders: %w", rows.Err())
	}

	return &courier.OrderListResponse{
		Orders: orders,
	}, nil
}

// UpdateOrderStatus buyurtma holatini yangilaydi
func (o *Order) UpdateOrderStatus(ctx context.Context, req *courier.UpdateOrderStatusRequest) (*courier.UpdateOrderStatusResponse, error) {
	query := `UPDATE orders SET status=$1 WHERE order_id=$2`
	_, err := o.Db.Exec(ctx, query, req.Status, req.OrderId)
	if err != nil {
		return nil, fmt.Errorf("failed to update order status: %w", err)
	}

	return &courier.UpdateOrderStatusResponse{
		Success: true,
		Message: "Order status updated successfully",
	}, nil
}

// GetCourierOrderHistory kuryer buyurtma tarixini qaytaradi
func (o *Order) GetCourierOrderHistory(ctx context.Context, req *courier.GetCourierOrderHistoryRequest) (*courier.GetCourierOrderHistoryResponse, error) {
	query := `SELECT order_id, customer_id, order_details, delivery_address, payment_status, order_date FROM orders WHERE courier_id=$1`
	rows, err := o.Db.Query(ctx, query, req.CourierId)
	if err != nil {
		return nil, fmt.Errorf("failed to get courier order history: %w", err)
	}
	defer rows.Close()

	var orders []*courier.Order
	for rows.Next() {
		var order courier.Order
		err := rows.Scan(
			&order.OrderId,
			&order.CustomerId,
			&order.OrderDetails,
			&order.DeliveryAddress,
			&order.PaymentStatus,
			&order.OrderDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}
		orders = append(orders, &order)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to iterate over orders: %w", rows.Err())
	}

	return &courier.GetCourierOrderHistoryResponse{
		Orders: orders,
	}, nil
}
