package courier

import (
	"context"
	us "courier_delivery/genproto/user"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthStruct struct {
	Db *pgx.Conn
}

func NewAuthStruct(db *pgx.Conn) *AuthStruct {
	return &AuthStruct{Db: db}
}

// HashPassword hashes a plain text password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// UserRegister registers a new user
func (u *AuthStruct) UserRegister(ctx context.Context, req *us.UserRegisterRequest) (*us.UserRegisterResponse, error) { 
	id := uuid.NewString()
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	query := `INSERT INTO users (user_id, email, password_hash, full_name) VALUES ($1, $2, $3, $4) RETURNING user_id, created_at`
	var userID string
	var createdAt time.Time
	err = u.Db.QueryRow(ctx, query, id, req.Email, hashedPassword, req.FullName).Scan(&userID, &createdAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &us.UserRegisterResponse{
		Success: true,
		Message: "User registered successfully",
		User: &us.User{
			UserId:    userID,
			Email:     req.Email,
			FullName:  req.FullName,
			CreatedAt: createdAt.Format(time.RFC3339),  // Formatlash
		},
	}, nil
}


// UserLogin authenticates a user
func (u *AuthStruct) UserLogin(ctx context.Context, req *us.UserLoginRequest) (*us.UserLoginResponse, error) {
	var storedPassword string
	query := `SELECT password_hash FROM users WHERE email=$1`
	err := u.Db.QueryRow(ctx, query, req.Email).Scan(&storedPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(req.Password))
	if err != nil {
		return nil, errors.New("incorrect password_hash")
	}

	return &us.UserLoginResponse{
		Success: true,
		Message: "User logged in successfully",
	}, nil
}

// CreateUser creates a new user record
func (u *AuthStruct) CreateUser(ctx context.Context, req *us.CreateUserRequest) (*us.UserResponse, error) {
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
func (u *AuthStruct) GetUser(ctx context.Context, req *us.UserRequest) (*us.UserResponse, error) {
	query := `SELECT user_id, email, role, created_at, updated_at, FROM users WHERE user_id=$1`
	var user us.User
	var Role string
	err := u.Db.QueryRow(ctx, query, req.UserId).Scan(
		&user.UserId,
		&user.Email,
		&user.FullName,
		&Role,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	fmt.Println(us.UserResponse{User:    &user})
	return &us.UserResponse{
		Success: true,
		Message: "User retrieved successfully",
		User:    &user,
		}, nil
}


// UpdateUser updates a user record
func (u *AuthStruct) UpdateUser(ctx context.Context, req *us.UpdateUserRequest) (*us.UserResponse, error) {
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
func (u *AuthStruct) DeleteUser(ctx context.Context, req *us.UserRequest) (*us.UserResponse, error) {
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
func (u *AuthStruct) GetAllUsers(ctx context.Context, req *us.GetAllUsersRequest) (*us.GetAllUsersResponse, error) {
	query := `SELECT user_id, email, role, created_at, updated_at FROM users`
	rows, err := u.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	var users []*us.User
	var Role string
	for rows.Next() {
		var user us.User
		var createdAt, updatedAt time.Time
		if err := rows.Scan(
			&user.UserId,
			&user.Email,
			&Role,       
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		user.CreatedAt = createdAt.Format(time.RFC3339) // Formatlash
		user.UpdatedAt = updatedAt.Format(time.RFC3339) // Formatlash
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

// CourierRegister registers a new courier
func (c *AuthStruct) CourierRegister(ctx context.Context, req *us.CourierRegisterRequest) (*us.CourierRegisterResponse, error) {
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	query := `INSERT INTO couriers (email, password, phone_number, status, created_at) VALUES ($1, $2, $3, $4, NOW()) RETURNING courier_id, created_at`
	var courierID, createdAt string
	err = c.Db.QueryRow(ctx, query, req.Email, hashedPassword, req.PhoneNumber).Scan(&courierID, &createdAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create courier: %w", err)
	}
	return &us.CourierRegisterResponse{
		Success: true,
		Message: "Courier registered successfully",
		Courier: &us.CourierAuth{
			CourierId:   courierID,
			Email:       req.Email,
			PhoneNumber: req.PhoneNumber,
			CreatedAt:   createdAt,
		},
	}, nil
}

// CourierLogin authenticates a courier
func (c *AuthStruct) CourierLogin(ctx context.Context, req *us.CourierLoginRequest) (*us.CourierLoginResponse, error) {
	var storedPassword string
	query := `SELECT password FROM couriers WHERE email=$1`
	err := c.Db.QueryRow(ctx, query, req.Email).Scan(&storedPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to find courier: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(req.Password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	return &us.CourierLoginResponse{
		Success: true,
		Message: "Courier logged in successfully",
	}, nil
}


// CreateCourier creates a new courier record
func (c *AuthStruct) CreateCourier(ctx context.Context, req *us.CreateCourierRequestAuth) (*us.CourierResponseAuth, error) {
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
	return &us.CourierResponseAuth{
		Success: true,
		Message: "Courier created successfully",
		Courier: &us.CourierAuth{
			CourierId:   courierID,
			Email:       req.Email,
			PhoneNumber: req.PhoneNumber,
			Status:      req.Status,
			CreatedAt:   createdAt,
		},
	}, nil
}

// GetCourier retrieves a courier record by courier_id
func (c *AuthStruct) GetCourier(ctx context.Context, req *us.CourierRequestAuth) (*us.CourierResponseAuth, error) {
	query := `SELECT courier_id, email, phone_number, status, created_at, updated_at FROM couriers WHERE courier_id=$1`
	var courier us.CourierAuth
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
	return &us.CourierResponseAuth{
		Success: true,
		Message: "Courier retrieved successfully",
		Courier: &courier,
	}, nil
}

// UpdateCourier updates a courier record
func (c *AuthStruct) UpdateCourier(ctx context.Context, req *us.UpdateCourierRequestAuth) (*us.CourierResponseAuth, error) {
	query := `UPDATE couriers SET email=$1, phone_number=$2, status=$3, updated_at=NOW() WHERE courier_id=$4`
	_, err := c.Db.Exec(ctx, query, req.Email, req.PhoneNumber, req.Status, req.CourierId)
	if err != nil {
		return nil, fmt.Errorf("failed to update courier: %w", err)
	}
	return &us.CourierResponseAuth{
		Success: true,
		Message: "Courier updated successfully",
		Courier: &us.CourierAuth{
			CourierId:   req.CourierId,
			Email:       req.Email,
			PhoneNumber: req.PhoneNumber,
			Status:      req.Status,
		},
	}, nil
}

// DeleteCourier deletes a courier record
func (c *AuthStruct) DeleteCourier(ctx context.Context, req *us.CourierRequestAuth) (*us.CourierResponseAuth, error) {
	query := `DELETE FROM couriers WHERE courier_id=$1`
	result, err := c.Db.Exec(ctx, query, req.CourierId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete courier: %w", err)
	}
	if result.RowsAffected() == 0 {
		return nil, errors.New("courier not found")
	}
	return &us.CourierResponseAuth{
		Success: true,
		Message: "Courier deleted successfully",
	}, nil
}

// GetAllCouriers retrieves all couriers
func (c *AuthStruct) GetAllCouriers(ctx context.Context, req *us.GetAllCouriersRequest) (*us.GetAllCouriersResponse, error) {
	query := `SELECT courier_id, email, phone_number, status, created_at, updated_at FROM couriers`
	rows, err := c.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get couriers: %w", err)
	}
	defer rows.Close()

	var couriers []*us.CourierAuth
	for rows.Next() {
		var courier us.CourierAuth
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
