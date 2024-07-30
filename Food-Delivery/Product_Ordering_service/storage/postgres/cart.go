package postgres

import (
	"context"
	"errors"
	"time"

	cart "product_ordering/genproto/product"

	"github.com/jackc/pgx/v5"
)

type Cart struct {
	Db *pgx.Conn
}

func NewCart(db *pgx.Conn) *Cart {
	return &Cart{Db: db}
}

func (c *Cart) Create(ctx context.Context, req *cart.CreateCartRequest) (*cart.CartResponse, error) {
	query := `INSERT INTO carts (user_id) VALUES ($1) RETURNING cart_id`
	var cartID string
	err := c.Db.QueryRow(ctx, query, req.UserId).Scan(&cartID)
	if err != nil {
		return nil, err
	}
	return &cart.CartResponse{
		Success: true,
		Message: "Cart created successfully",
		Cart: &cart.Cart{
			CartId:    cartID,
			UserId:    req.UserId,
			CreatedAt: time.Now().Format(time.RFC3339),
		},
	}, nil
}

func (c *Cart) Get(ctx context.Context, req *cart.CartRequest) (*cart.CartResponse, error) {
	query := `SELECT cart_id, user_id, created_at FROM carts WHERE cart_id=$1 AND deleted_at=0`
	var item cart.Cart
	err := c.Db.QueryRow(ctx, query, req.CartId).Scan(&item.CartId, &item.UserId, &item.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("cart not found")
		}
		return nil, err
	}
	return &cart.CartResponse{
		Success: true,
		Message: "Cart retrieved successfully",
		Cart:    &item,
	}, nil
}

func (c *Cart) Update(ctx context.Context, req *cart.UpdateCartRequest) (*cart.CartResponse, error) {
	// Mavjud qiymatlarni olish
	getQuery := `SELECT user_id FROM carts WHERE cart_id=$1 AND deleted_at=0`
	var currentUserId string
	err := c.Db.QueryRow(ctx, getQuery, req.CartId).Scan(&currentUserId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("cart not found")
		}
		return nil, err
	}

	// Agar yangi qiymatlar berilmagan bo'lsa, eski qiymatlarni saqlab qolish
	userId := req.UserId
	if userId == "" {
		userId = currentUserId
	}

	// Yangilash
	updateQuery := `UPDATE carts SET user_id=$1, updated_at=NOW(), deleted_at=0 WHERE cart_id=$2`
	_, err = c.Db.Exec(ctx, updateQuery, userId, req.CartId)
	if err != nil {
		return nil, err
	}

	return &cart.CartResponse{
		Success: true,
		Message: "Cart updated successfully",
		Cart: &cart.Cart{
			CartId: req.CartId,
			UserId: userId,
			// `created_at` qiymati o'zgarmaydi va asl qiymat saqlanadi
		},
	}, nil
}

func (c *Cart) Delete(ctx context.Context, req *cart.CartRequest) (*cart.CartResponse, error) {
	query := `UPDATE carts SET deleted_at=EXTRACT(EPOCH FROM NOW())::BIGINT WHERE cart_id=$1 AND deleted_at=0`
	result, err := c.Db.Exec(ctx, query, req.CartId)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, errors.New("cart not found or already deleted")
	}
	return &cart.CartResponse{
		Success: true,
		Message: "Cart deleted successfully",
	}, nil
}

func (c *Cart) List(ctx context.Context, req *cart.Empty) (*cart.CartListResponse, error) {
	query := `SELECT cart_id, user_id, created_at FROM carts WHERE deleted_at=0`
	rows, err := c.Db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []*cart.Cart
	for rows.Next() {
		var item cart.Cart
		err := rows.Scan(&item.CartId, &item.UserId, &item.CreatedAt)
		if err != nil {
			return nil, err
		}
		carts = append(carts, &item)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &cart.CartListResponse{
		Carts: carts,
	}, nil
}
