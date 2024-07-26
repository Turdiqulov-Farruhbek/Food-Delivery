package postgrestest

import (
	"context"
	"courier_delivery/genproto"
	cr "courier_delivery/storage/postgres/courier"
	"fmt"
	"testing"

	pgx "github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *pgx.Conn {
	// Test ma'lumotlar bazasiga ulanish (bu yerda o'zingizning test ma'lumotlar bazasi parametrlarini kiriting)
	connString := "postgres://postgres:root@localhost:5432/testdb?sslmode=disable"
	db, err := pgx.Connect(context.Background(), connString)
	require.NoError(t, err)

	// Test uchun zarur bo'lgan jadval va ma'lumotlarni tayyorlash
	_, err = db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS courierorders (
		courier_order_id SERIAL PRIMARY KEY,
		courier_id VARCHAR(50),
		order_id VARCHAR(50),
		status INT,
		assigned_time TIMESTAMP,
		last_updated TIMESTAMP,
		deleted_at BIGINT DEFAULT 0
	)`)
	require.NoError(t, err)

	return db
}

func teardownTestDB(t *testing.T, db *pgx.Conn) {
	// Test tugagach jadvalni tozalash
	_, err := db.Exec(context.Background(), `DROP TABLE IF EXISTS courierorders`)
	require.NoError(t, err)
}

func TestCourierOrder_CreateCourierOrder(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	storage := cr.NewCourierOrder(db)

	req := &genproto.CreateCourierOrderRequest{
		CourierId: "courier_123",
		OrderId:   "order_456",
		Status:    genproto.CourierOrderStatus_ASSIGNED,
	}

	resp, err := storage.CreateCourierOrder(context.Background(), req)
	require.NoError(t, err)
	assert.True(t, resp.Success)
	assert.NotEmpty(t, resp.CourierOrder.CourierOrderId)
}

func TestCourierOrder_GetCourierOrder(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	storage := cr.NewCourierOrder(db)

	// Avval yangi kuryer-buyurtma munosabatini yaratish
	req := &genproto.CreateCourierOrderRequest{
		CourierId: "courier_123",
		OrderId:   "order_456",
		Status:    genproto.CourierOrderStatus_ASSIGNED,
	}

	createResp, err := storage.CreateCourierOrder(context.Background(), req)
	require.NoError(t, err)

	// Yaratilgan buyurtmani olish
	getReq := &genproto.CourierOrderRequest{CourierOrderId: createResp.CourierOrder.CourierOrderId}
	getResp, err := storage.GetCourierOrder(context.Background(), getReq)
	require.NoError(t, err)
	assert.True(t, getResp.Success)
	assert.Equal(t, req.CourierId, getResp.CourierOrder.CourierId)
	assert.Equal(t, req.OrderId, getResp.CourierOrder.OrderId)
}

func TestCourierOrder_UpdateCourierOrder(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	storage := cr.NewCourierOrder(db)

	// Avval yangi kuryer-buyurtma munosabatini yaratish
	req := &genproto.CreateCourierOrderRequest{
		CourierId: "courier_123",
		OrderId:   "order_456",
		Status:    genproto.CourierOrderStatus_ASSIGNED,
	}

	createResp, err := storage.CreateCourierOrder(context.Background(), req)
	require.NoError(t, err)

	// Yaratilgan buyurtmani yangilash
	updateReq := &genproto.UpdateCourierOrderRequest{
		CourierOrderId: createResp.CourierOrder.CourierOrderId,
		CourierId:      "courier_789",
		OrderId:        "order_101",
		Status:         genproto.CourierOrderStatus_DELIVERED,
	}

	updateResp, err := storage.UpdateCourierOrder(context.Background(), updateReq)
	require.NoError(t, err)
	assert.True(t, updateResp.Success)
	assert.Equal(t, updateReq.CourierId, updateResp.CourierOrder.CourierId)
	assert.Equal(t, updateReq.OrderId, updateResp.CourierOrder.OrderId)
}

func TestCourierOrder_DeleteCourierOrder(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	storage := cr.NewCourierOrder(db)

	// Avval yangi kuryer-buyurtma munosabatini yaratish
	req := &genproto.CreateCourierOrderRequest{
		CourierId: "courier_123",
		OrderId:   "order_456",
		Status:    genproto.CourierOrderStatus_ASSIGNED,
	}

	createResp, err := storage.CreateCourierOrder(context.Background(), req)
	require.NoError(t, err)

	// Yaratilgan buyurtmani o'chirish
	deleteReq := &genproto.CourierOrderRequest{CourierOrderId: createResp.CourierOrder.CourierOrderId}
	deleteResp, err := storage.DeleteCourierOrder(context.Background(), deleteReq)
	require.NoError(t, err)
	assert.True(t, deleteResp.Success)

	// O'chirilgan buyurtmani qayta olishga urinish
	getReq := &genproto.CourierOrderRequest{CourierOrderId: createResp.CourierOrder.CourierOrderId}
	_, err = storage.GetCourierOrder(context.Background(), getReq)
	assert.Error(t, err) // Buyurtma topilmasligi kerak
}

func TestCourierOrder_ListCourierOrders(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	storage := cr.NewCourierOrder(db)

	// Bir nechta kuryer-buyurtma munosabatlarini yaratish
	for i := 0; i < 3; i++ {
		req := &genproto.CreateCourierOrderRequest{
			CourierId: fmt.Sprintf("courier_%d", i),
			OrderId:   fmt.Sprintf("order_%d", i),
			Status:    genproto.CourierOrderStatus_ASSIGNED,
		}
		_, err := storage.CreateCourierOrder(context.Background(), req)
		require.NoError(t, err)
	}

	// Barcha kuryer-buyurtma munosabatlarini olish
	listResp, err := storage.ListCourierOrders(context.Background(), &genproto.Empty{})
	require.NoError(t, err)
	assert.Equal(t, 3, len(listResp.CourierOrders))
}
