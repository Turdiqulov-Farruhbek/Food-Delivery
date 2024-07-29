package courier

import (
	"context"
	"errors"
	"fmt"
	us "courier_delivery/genproto/user"


	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Db *pgx.Conn
}

func NewUserService(db *pgx.Conn) *UserService {
	return &UserService{Db: db}
}

// HashPassword hashes a plain text password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CreateUser creates a new user record
func (u *UserService) CreateUser(ctx context.Context, req *us.CreateUserRequest) (*us.UserResponse, error) {
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	query := `INSERT INTO users (email, password, role, created_at) VALUES ($1, $2, $3, NOW()) RETURNING user_id, created_at`
	var userID, createdAt string
	err = u.Db.QueryRow(ctx, query, req.Email, hashedPassword).Scan(&userID, &createdAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &us.UserResponse{
		Success: true,
		Message: "User created successfully",
		User: &us.User{
			UserId:    userID,
			Email:     req.Email,
			CreatedAt: createdAt,
		},
	}, nil
}

// GetUser retrieves a user record by user_id
func (u *UserService) GetUser(ctx context.Context, req *us.UserRequest) (*us.UserResponse, error) {
	query := `SELECT user_id, email, role, created_at, updated_at FROM users WHERE user_id=$1`
	var user us.User
	err := u.Db.QueryRow(ctx, query, req.UserId).Scan(
		&user.UserId,
		&user.Email,
		&user.Role,
		&user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &us.UserResponse{
		Success: true,
		Message: "User retrieved successfully",
		User:    &user,
	}, nil
}

// UpdateUser updates a user record
func (u *UserService) UpdateUser(ctx context.Context, req *us.UpdateUserRequest) (*us.UserResponse, error) {
	query := `UPDATE users SET email=$1, role=$2, updated_at=NOW() WHERE user_id=$3`
	_, err := u.Db.Exec(ctx, query, req.Email, req.Role, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return &us.UserResponse{
		Success: true,
		Message: "User updated successfully",
		User: &us.User{
			UserId: req.UserId,
			Email:  req.Email,
			Role:   req.Role,
		},
	}, nil
}

// DeleteUser deletes a user record
func (u *UserService) DeleteUser(ctx context.Context, req *us.UserRequest) (*us.UserResponse, error) {
	query := `DELETE FROM users WHERE user_id=$1`
	result, err := u.Db.Exec(ctx, query, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %w", err)
	}
	if result.RowsAffected() == 0 {
		return nil, errors.New("user not found")
	}
	return &us.UserResponse{
		Success: true,
		Message: "User deleted successfully",
	}, nil
}


// GetAllUsers retrieves all users
func (u *UserService) GetAllUsers(ctx context.Context, req *us.GetAllUsersRequest) (*us.GetAllUsersResponse, error) {
	query := `SELECT user_id, email, role, created_at, updated_at FROM users`
	rows, err := u.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	var users []*us.User
	for rows.Next() {
		var user us.User
		if err := rows.Scan(
			&user.UserId,
			&user.Email,
			&user.Role,
			&user.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return &us.GetAllUsersResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Users:   users,
	}, nil
}

//==========================================================================================================================================



type CourierService struct {
	Db *pgx.Conn
}

func NewCourierService(db *pgx.Conn) *CourierService {
	return &CourierService{Db: db}
}

// CreateCourier creates a new courier record
func (c *CourierService) CreateCourier(ctx context.Context, req *us.CreateCourierRequest) (*us.CourierResponse, error) {
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	query := `INSERT INTO couriers (email, password, phone_number, status, created_at) VALUES ($1, $2, $3, $4, NOW()) RETURNING courier_id, created_at`
	var courierID, createdAt string
	err = c.Db.QueryRow(ctx, query, req.Email, hashedPassword, req.PhoneNumber, req.Status).Scan(&courierID, &createdAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create courier: %w", err)
	}
	return &us.CourierResponse{
		Success: true,
		Message: "Courier created successfully",
		Courier: &us.Courier{
			CourierId: courierID,
			Email:     req.Email,
			PhoneNumber: req.PhoneNumber,
			Status:    req.Status,
			CreatedAt: createdAt,
		},
	}, nil
}

// GetCourier retrieves a courier record by courier_id
func (c *CourierService) GetCourier(ctx context.Context, req *us.CourierRequest) (*us.CourierResponse, error) {
	query := `SELECT courier_id, email, phone_number, status, created_at, updated_at FROM couriers WHERE courier_id=$1`
	var courier us.Courier
	err := c.Db.QueryRow(ctx, query, req.CourierId).Scan(
		&courier.CourierId,
		&courier.Email,
		&courier.PhoneNumber,
		&courier.Status,
		&courier.CreatedAt,
		&courier.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get courier: %w", err)
	}
	return &us.CourierResponse{
		Success: true,
		Message: "Courier retrieved successfully",
		Courier: &courier,
	}, nil
}

// UpdateCourier updates a courier record
func (c *CourierService) UpdateCourier(ctx context.Context, req *us.UpdateCourierRequest) (*us.CourierResponse, error) {
	query := `UPDATE couriers SET email=$1, phone_number=$2, status=$3, updated_at=NOW() WHERE courier_id=$4`
	_, err := c.Db.Exec(ctx, query, req.Email, req.PhoneNumber, req.Status, req.CourierId)
	if err != nil {
		return nil, fmt.Errorf("failed to update courier: %w", err)
	}
	return &us.CourierResponse{
		Success: true,
		Message: "Courier updated successfully",
		Courier: &us.Courier{
			CourierId: req.CourierId,
			Email:     req.Email,
			PhoneNumber: req.PhoneNumber,
			Status:    req.Status,
		},
	}, nil
}

// DeleteCourier deletes a courier record
func (c *CourierService) DeleteCourier(ctx context.Context, req *us.CourierRequest) (*us.CourierResponse, error) {
	query := `DELETE FROM couriers WHERE courier_id=$1`
	result, err := c.Db.Exec(ctx, query, req.CourierId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete courier: %w", err)
	}
	if result.RowsAffected() == 0 {
		return nil, errors.New("courier not found")
	}
	return &us.CourierResponse{
		Success: true,
		Message: "Courier deleted successfully",
	}, nil
}

// GetAllCouriers retrieves all couriers
func (c *CourierService) GetAllCouriers(ctx context.Context, req *us.GetAllCouriersRequest) (*us.GetAllCouriersResponse, error) {
	query := `SELECT courier_id, email, phone_number, status, created_at, updated_at FROM couriers`
	rows, err := c.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get couriers: %w", err)
	}
	defer rows.Close()

	var couriers []*us.Courier
	for rows.Next() {
		var courier us.Courier
		if err := rows.Scan(
			&courier.CourierId,
			&courier.Email,
			&courier.PhoneNumber,
			&courier.Status,
			&courier.CreatedAt,
			&courier.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan courier: %w", err)
		}
		couriers = append(couriers, &courier)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return &us.GetAllCouriersResponse{
		Success:  true,
		Message:  "Couriers retrieved successfully",
		Couriers: couriers,
	}, nil
}
