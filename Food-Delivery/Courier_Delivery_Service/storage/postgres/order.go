package postgres

import (
	"context"
	"courier_delivery/genproto"
	"errors"
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
func (o *Order) CreateOrder(ctx context.Context, req *genproto.CreateOrderRequest) (*genproto.OrderResponse, error) {
	query := `INSERT INTO orders (customer_id, order_details, delivery_address, payment_status, order_date) VALUES ($1, $2, $3, $4, NOW()) RETURNING order_id`
	var orderID string
	err := o.Db.QueryRow(ctx, query, req.CustomerId, req.OrderDetails, req.DeliveryAddress, req.PaymentStatus).Scan(&orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}
	return &genproto.OrderResponse{
		Success: true,
		Message: "Order created successfully",
		Order: &genproto.Order{
			OrderId:         orderID,
			CustomerId:      req.CustomerId,
			OrderDetails:    req.OrderDetails,
			DeliveryAddress: req.DeliveryAddress,
			PaymentStatus:   req.PaymentStatus,
			OrderDate:       time.Now().Format(time.RFC3339), // Vaqtni RFC3339 formatida saqlash,
		},
	}, nil
}

// GetOrder buyurtma yozuvini qaytaradi
func (o *Order) GetOrder(ctx context.Context, req *genproto.OrderRequest) (*genproto.OrderResponse, error) {
	query := `SELECT order_id, customer_id, order_details, delivery_address, payment_status, order_date FROM orders WHERE order_id=$1 AND deleted_at=0`
	var order genproto.Order
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
	return &genproto.OrderResponse{
		Success: true,
		Message: "Order retrieved successfully",
		Order:   &order,
	}, nil
}

// UpdateOrder buyurtma yozuvini yangilaydi
func (o *Order) UpdateOrder(ctx context.Context, req *genproto.UpdateOrderRequest) (*genproto.OrderResponse, error) {
	// Avval eski qiymatlarni olish
	var oldCustomerId, oldOrderDetails, oldDeliveryAddress string
	var oldPaymentStatus genproto.PaymentStatus

	getQuery := `SELECT customer_id, order_details, delivery_address, payment_status FROM orders WHERE order_id=$1`
	err := o.Db.QueryRow(ctx, getQuery, req.OrderId).Scan(&oldCustomerId, &oldOrderDetails, &oldDeliveryAddress, &oldPaymentStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to get existing order: %w", err)
	}

	// Yangi qiymatlar bo'sh bo'lsa, eski qiymatlarni ishlatish
	customerId := req.CustomerId
	if customerId == "" {
		customerId = oldCustomerId
	}

	orderDetails := req.OrderDetails
	if orderDetails == "" {
		orderDetails = oldOrderDetails
	}

	deliveryAddress := req.DeliveryAddress
	if deliveryAddress == "" {
		deliveryAddress = oldDeliveryAddress
	}

	paymentStatus := req.PaymentStatus
	if req.PaymentStatus == genproto.PaymentStatus(0) { // 0 status PENDING sifatida ko'riladi
		paymentStatus = oldPaymentStatus
	}

	// Yangilanish so'rovi
	updateQuery := `UPDATE orders SET customer_id=$1, order_details=$2, delivery_address=$3, payment_status=$4, order_date=NOW() WHERE order_id=$5 AND deleted_at=0`
	_, err = o.Db.Exec(ctx, updateQuery, customerId, orderDetails, deliveryAddress, paymentStatus, req.OrderId)
	if err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	return &genproto.OrderResponse{
		Success: true,
		Message: "Order updated successfully",
		Order: &genproto.Order{
			OrderId:         req.OrderId,
			CustomerId:      customerId,
			OrderDetails:    orderDetails,
			DeliveryAddress: deliveryAddress,
			PaymentStatus:   paymentStatus,
			OrderDate:       time.Now().Format(time.RFC3339), // Vaqtni RFC3339 formatida saqlash,
		},
	}, nil
}

// DeleteOrder buyurtma yozuvini mantiqiy o'chiradi
func (o *Order) DeleteOrder(ctx context.Context, req *genproto.OrderRequest) (*genproto.OrderResponse, error) {
	query := `UPDATE orders SET deleted_at=EXTRACT(EPOCH FROM NOW())::BIGINT WHERE order_id=$1 AND deleted_at=0`
	result, err := o.Db.Exec(ctx, query, req.OrderId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete order: %w", err)
	}

	if result.RowsAffected() == 0 {
		return nil, errors.New("order not found or already deleted")
	}

	return &genproto.OrderResponse{
		Success: true,
		Message: "Order deleted successfully",
	}, nil
}

// ListOrders barcha buyurtma yozuvlarini qaytaradi
func (o *Order) ListOrders(ctx context.Context, req *genproto.Empty) (*genproto.OrderListResponse, error) {
	query := `SELECT order_id, customer_id, order_details, delivery_address, payment_status, order_date FROM orders WHERE deleted_at=0`
	rows, err := o.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list orders: %w", err)
	}
	defer rows.Close()

	var orders []*genproto.Order
	for rows.Next() {
		var order genproto.Order
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

	return &genproto.OrderListResponse{
		Orders: orders,
	}, nil
}
