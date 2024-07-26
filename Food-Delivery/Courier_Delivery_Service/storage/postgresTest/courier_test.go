package postgrestest

import (
	"context"
	"courier_delivery/genproto/courier"
	cr "courier_delivery/storage/postgres/courier"
	"fmt"
	"testing"

	pgx "github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestDBCourer test ma'lumotlar bazasini tayyorlaydi.
func setupTestDBCourer(t *testing.T) *pgx.Conn {
	// Test ma'lumotlar bazasiga ulanish (bu yerda o'zingizning test ma'lumotlar bazasi parametrlarini kiriting)
	connString := "postgres://postgres:root@localhost:5432/testdb?sslmode=disable"
	db, err := pgx.Connect(context.Background(), connString)
	require.NoError(t, err)

	// Test uchun zarur bo'lgan jadval va ma'lumotlarni tayyorlash
	_, err = db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS couriers (
		courier_id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		phone_number VARCHAR(50),
		email VARCHAR(255),
		status VARCHAR(50),
		deleted_at BIGINT DEFAULT 0
	)`)
	require.NoError(t, err)

	return db
}

// teardownTestDB test tugagach ma'lumotlar bazasini tozalaydi.
func teardownTestDBCourier(t *testing.T, db *pgx.Conn) {
	_, err := db.Exec(context.Background(), `DROP TABLE IF EXISTS couriers`)
	require.NoError(t, err)
}

// TestCreateCourier yangi kuryer yozuvini yaratishni sinovdan o'tkazadi.
func TestCreateCourier(t *testing.T) {
	db := setupTestDBCourer(t)
	defer teardownTestDBCourier(t, db)

	storage := cr.NewCourier(db)

	req := &courier.CreateCourierRequest{
		Name:        "John Doe",
		PhoneNumber: "1234567890",
		Email:       "john@example.com",
		Status:      courier.CourierStatus_UNAVAILABLE,
	}

	resp, err := storage.CreateCourier(context.Background(), req)
	require.NoError(t, err)
	assert.True(t, resp.Success)
	assert.NotEmpty(t, resp.Courier.CourierId)
	assert.Equal(t, req.Name, resp.Courier.Name)
	assert.Equal(t, req.PhoneNumber, resp.Courier.PhoneNumber)
	assert.Equal(t, req.Email, resp.Courier.Email)
	assert.Equal(t, req.Status, resp.Courier.Status)
}

// TestGetCourier kuryer yozuvini kuryer_id orqali olishni sinovdan o'tkazadi.
func TestGetCourier(t *testing.T) {
	db := setupTestDBCourer(t)
	defer teardownTestDBCourier(t, db)

	storage := cr.NewCourier(db)

	// Avval yangi kuryer yozuvini yaratish
	req := &courier.CreateCourierRequest{
		Name:        "John Doe",
		PhoneNumber: "1234567890",
		Email:       "john@example.com",
		Status:      courier.CourierStatus_AVAILABLE,
	}

	createResp, err := storage.CreateCourier(context.Background(), req)
	require.NoError(t, err)

	// Yaratilgan kuryer yozuvini olish
	getReq := &courier.CourierRequest{CourierId: createResp.Courier.CourierId}
	getResp, err := storage.GetCourier(context.Background(), getReq)
	require.NoError(t, err)
	assert.True(t, getResp.Success)
	assert.Equal(t, req.Name, getResp.Courier.Name)
	assert.Equal(t, req.PhoneNumber, getResp.Courier.PhoneNumber)
	assert.Equal(t, req.Email, getResp.Courier.Email)
	assert.Equal(t, req.Status, getResp.Courier.Status)
}

// TestUpdateCourier kuryer yozuvini yangilashni sinovdan o'tkazadi.
func TestUpdateCourier(t *testing.T) {
	db := setupTestDBCourer(t)
	defer teardownTestDBCourier(t, db)

	storage := cr.NewCourier(db)

	// Avval yangi kuryer yozuvini yaratish
	req := &courier.CreateCourierRequest{
		Name:        "John Doe",
		PhoneNumber: "1234567890",
		Email:       "john@example.com",
		Status:      courier.CourierStatus_UNAVAILABLE,
	}

	createResp, err := storage.CreateCourier(context.Background(), req)
	require.NoError(t, err)

	// Yaratilgan kuryer yozuvini yangilash
	updateReq := &courier.UpdateCourierRequest{
		CourierId:   createResp.Courier.CourierId,
		Name:        "Jane Doe",
		PhoneNumber: "0987654321",
		Email:       "jane@example.com",
		Status:      courier.CourierStatus_AVAILABLE,
	}

	updateResp, err := storage.UpdateCourier(context.Background(), updateReq)
	require.NoError(t, err)
	assert.True(t, updateResp.Success)
	assert.Equal(t, updateReq.Name, updateResp.Courier.Name)
	assert.Equal(t, updateReq.PhoneNumber, updateResp.Courier.PhoneNumber)
	assert.Equal(t, updateReq.Email, updateResp.Courier.Email)
	assert.Equal(t, updateReq.Status, updateResp.Courier.Status)
}

// TestDeleteCourier kuryer yozuvini mantiqiy o'chirishni sinovdan o'tkazadi.
func TestDeleteCourier(t *testing.T) {
	db := setupTestDBCourer(t)
	defer teardownTestDBCourier(t, db)

	storage := cr.NewCourier(db)

	// Avval yangi kuryer yozuvini yaratish
	req := &courier.CreateCourierRequest{
		Name:        "John Doe",
		PhoneNumber: "1234567890",
		Email:       "john@example.com",
		Status:      courier.CourierStatus_UNAVAILABLE,
	}

	createResp, err := storage.CreateCourier(context.Background(), req)
	require.NoError(t, err)

	// Yaratilgan kuryer yozuvini o'chirish
	deleteReq := &courier.CourierRequest{CourierId: createResp.Courier.CourierId}
	deleteResp, err := storage.DeleteCourier(context.Background(), deleteReq)
	require.NoError(t, err)
	assert.True(t, deleteResp.Success)

	// O'chirilgan kuryer yozuvini qayta olishga urinish
	getReq := &courier.CourierRequest{CourierId: createResp.Courier.CourierId}
	_, err = storage.GetCourier(context.Background(), getReq)
	assert.Error(t, err) // Kuryer topilmasligi kerak
}

// TestListCouriers barcha kuryer yozuvlarini olishni sinovdan o'tkazadi.
func TestListCouriers(t *testing.T) {
	db := setupTestDBCourer(t)
	defer teardownTestDBCourier(t, db)

	storage := cr.NewCourier(db)

	// Bir nechta kuryer yozuvlarini yaratish
	for i := 0; i < 3; i++ {
		req := &courier.CreateCourierRequest{
			Name:        fmt.Sprintf("Courier %d", i),
			PhoneNumber: fmt.Sprintf("12345678%d", i),
			Email:       fmt.Sprintf("courier%d@example.com", i),
			Status:      courier.CourierStatus_UNAVAILABLE,
		}
		_, err := storage.CreateCourier(context.Background(), req)
		require.NoError(t, err)
	}

	// Barcha kuryer yozuvlarini olish
	listResp, err := storage.ListCouriers(context.Background(), &courier.Empty{})
	require.NoError(t, err)
	assert.Equal(t, 3, len(listResp.Couriers))
}
