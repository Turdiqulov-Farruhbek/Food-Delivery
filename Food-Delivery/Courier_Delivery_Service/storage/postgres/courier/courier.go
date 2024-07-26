package courier

import (
	"context"
	"courier_delivery/genproto"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Courier struct {
	Db *pgx.Conn
}

func NewCourier(db *pgx.Conn) *Courier{
	return &Courier{Db: db}
}

// CreateCourier yangi kuryer yozuvini yaratadi
func (c *Courier) CreateCourier(ctx context.Context, req *genproto.CreateCourierRequest) (*genproto.CourierResponse, error) {
	query := `INSERT INTO couriers (name, phone_number, email, status) VALUES ($1, $2, $3, $4) RETURNING courier_id`
	var courierID string
	err := c.Db.QueryRow(ctx, query, req.Name, req.PhoneNumber, req.Email, req.Status).Scan(&courierID)
	if err != nil {
		return nil, fmt.Errorf("failed to create courier: %w", err)
	}
	return &genproto.CourierResponse{
		Success: true,
		Message: "Courier created successfully",
		Courier: &genproto.Courier{
			CourierId:   courierID,
			Name:        req.Name,
			PhoneNumber: req.PhoneNumber,
			Email:       req.Email,
			Status:      req.Status,
		},
	}, nil
}

// GetCourier kuryer yozuvini kuryer_id orqali qaytaradi
func (c *Courier) GetCourier(ctx context.Context, req *genproto.CourierRequest) (*genproto.CourierResponse, error) {
	query := `SELECT courier_id, name, phone_number, email, status FROM couriers WHERE courier_id=$1 AND deleted_at=0`
	var courier genproto.Courier
	err := c.Db.QueryRow(ctx, query, req.CourierId).Scan(
		&courier.CourierId,
		&courier.Name,
		&courier.PhoneNumber,
		&courier.Email,
		&courier.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to get courier: %w", err)
	}
	return &genproto.CourierResponse{
		Success: true,
		Message: "Courier retrieved successfully",
		Courier: &courier,
	}, nil
}

// UpdateCourier kuryer yozuvini yangilaydi
func (c *Courier) UpdateCourier(ctx context.Context, req *genproto.UpdateCourierRequest) (*genproto.CourierResponse, error) {
	// Eski qiymatlarni olish
	var oldName, oldPhoneNumber, oldEmail string
	var oldStatus genproto.CourierStatus

	getQuery := `SELECT name, phone_number, email, status FROM couriers WHERE courier_id=$1`
	err := c.Db.QueryRow(ctx, getQuery, req.CourierId).Scan(&oldName, &oldPhoneNumber, &oldEmail, &oldStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to get existing courier: %w", err)
	}

	// Yangi qiymatlar bo'sh bo'lsa, eski qiymatlarni ishlatish
	name := req.Name
	if name == "" {
		name = oldName
	}

	phoneNumber := req.PhoneNumber
	if phoneNumber == "" {
		phoneNumber = oldPhoneNumber
	}

	email := req.Email
	if email == "" {
		email = oldEmail
	}

	status := req.Status
	if req.Status == genproto.CourierStatus(0) {
		status = oldStatus
	}

	// Yangilanish so'rovi
	updateQuery := `UPDATE couriers SET name=$1, phone_number=$2, email=$3, status=$4 WHERE courier_id=$5 AND deleted_at=0`
	_, err = c.Db.Exec(ctx, updateQuery, name, phoneNumber, email, status, req.CourierId)
	if err != nil {
		return nil, fmt.Errorf("failed to update courier: %w", err)
	}

	return &genproto.CourierResponse{
		Success: true,
		Message: "Courier updated successfully",
		Courier: &genproto.Courier{
			CourierId:   req.CourierId,
			Name:        name,
			PhoneNumber: phoneNumber,
			Email:       email,
			Status:      status,
		},
	}, nil
}

// DeleteCourier kuryer yozuvini mantiqiy o'chiradi (deleted_at ni yangilash)
func (c *Courier) DeleteCourier(ctx context.Context, req *genproto.CourierRequest) (*genproto.CourierResponse, error) {
	query := `UPDATE couriers SET deleted_at=EXTRACT(EPOCH FROM NOW())::BIGINT WHERE courier_id=$1 AND deleted_at=0`
	result, err := c.Db.Exec(ctx, query, req.CourierId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete courier: %w", err)
	}

	// Agar hech qator yangilanmagan bo'lsa, yozuv topilmagan
	if result.RowsAffected() == 0 {
		return nil, errors.New("courier not found or already deleted")
	}

	return &genproto.CourierResponse{
		Success: true,
		Message: "Courier deleted successfully",
	}, nil
}

// ListCouriers barcha kuryer yozuvlarini qaytaradi
func (c *Courier) ListCouriers(ctx context.Context, req *genproto.Empty) (*genproto.CourierListResponse, error) {
	query := `SELECT courier_id, name, phone_number, email, status FROM couriers WHERE deleted_at=0`
	rows, err := c.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list couriers: %w", err)
	}
	defer rows.Close()

	var couriers []*genproto.Courier
	for rows.Next() {
		var courier genproto.Courier
		err := rows.Scan(
			&courier.CourierId,
			&courier.Name,
			&courier.PhoneNumber,
			&courier.Email,
			&courier.Status)
		if err != nil {
			return nil, fmt.Errorf("failed to scan courier: %w", err)
		}
		couriers = append(couriers, &courier)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to iterate over couriers: %w", rows.Err())
	}

	return &genproto.CourierListResponse{
		Couriers: couriers,
	}, nil
}
