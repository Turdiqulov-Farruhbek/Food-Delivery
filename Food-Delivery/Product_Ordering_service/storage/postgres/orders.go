package postgres

import (
	"context"
	"errors"
	order "product_ordering/genproto"
	"github.com/jackc/pgx/v5"
)

type Order struct {
	Db *pgx.Conn
}

func NewOrder(db *pgx.Conn) *Order {
	return &Order{Db: db}
}

func (o *Order) Create(ctx context.Context, req *order.CreateOrderRequest) (*order.OrderResponse, error) {
	query := `INSERT INTO orders (user_id, status, total_price) VALUES ($1, $2, $3) RETURNING order_id`
	var orderID string
	var createdAt string
	err := o.Db.QueryRow(ctx, query, req.UserId, req.Status, req.TotalPrice).Scan(&orderID, &createdAt)
	if err != nil {
		return nil, err
	}
	return &order.OrderResponse{
		Success: true,
		Message: "Order created successfully",
		Order: &order.Order{
			OrderId:    orderID,
			UserId:     req.UserId,
			CreatedAt:  createdAt,
			Status:     req.Status,
			TotalPrice: req.TotalPrice,
		},
	}, nil
}

func (o *Order) Get(ctx context.Context, req *order.OrderRequest) (*order.OrderResponse, error) {
	query := `SELECT order_id, user_id, created_at, status, total_price FROM orders WHERE order_id=$1 AND deleted_at=0`
	var ord order.Order
	err := o.Db.QueryRow(ctx, query, req.OrderId).Scan(&ord.OrderId, &ord.UserId, &ord.CreatedAt, &ord.Status, &ord.TotalPrice)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &order.OrderResponse{
		Success: true,
		Message: "Order retrieved successfully",
		Order:   &ord,
	}, nil
}

func (o *Order) Update(ctx context.Context, req *order.UpdateOrderRequest) (*order.OrderResponse, error) {
	query := `UPDATE orders SET status=$1, total_price=$2 WHERE order_id=$3 AND deleted_at=0`
	_, err := o.Db.Exec(ctx, query, req.Status, req.TotalPrice, req.OrderId)
	if err != nil {
		return nil, err
	}
	return &order.OrderResponse{
		Success: true,
		Message: "Order updated successfully",
		Order: &order.Order{
			OrderId:    req.OrderId,
			Status:     req.Status,
			TotalPrice: req.TotalPrice,
		},
	}, nil
}

func (o *Order) Delete(ctx context.Context, req *order.OrderRequest) (*order.OrderResponse, error) {
	query :=  `UPDATE orders SET deleted_at=EXTRACT(EPOCH FROM NOW())::BIGINT  WHERE deleted_at=0`
	result, err := o.Db.Exec(ctx, query, req.OrderId)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, errors.New("order not found or already deleted")
	}
	return &order.OrderResponse{
		Success: true,
		Message: "Order deleted successfully",
	}, nil
}

func (o *Order) List(ctx context.Context, req *order.OrderListRequest) (*order.OrderListResponse, error) {
	query := `SELECT order_id, user_id, created_at, status, total_price FROM orders deleted_at=0`
	var args []interface{}
	var conditions []string

	// Filter by user_id if provided
	if req.UserId != "" {
		conditions = append(conditions, "user_id=$1")
		args = append(args, req.UserId)
	}
	// Filter by status if provided
	if req.Status != order.CardStatus_OR_PENDING {
		conditions = append(conditions, "status=$2")
		args = append(args, req.Status)
	}
	if len(conditions) > 0 {
		query += " WHERE " + conditions[0]
		if len(conditions) > 1 {
			for i := 1; i < len(conditions); i++ {
				query += " AND " + conditions[i]
			}
		}
	}

	rows, err := o.Db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*order.Order
	for rows.Next() {
		var ord order.Order
		err := rows.Scan(&ord.OrderId, &ord.UserId, &ord.CreatedAt, &ord.Status, &ord.TotalPrice)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &ord)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &order.OrderListResponse{
		Orders: orders,
	}, nil
}
