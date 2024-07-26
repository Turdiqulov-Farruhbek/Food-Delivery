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

// setupTestDB tizim test ma'lumotlar bazasini tayyorlaydi.
func setupTestDBOrder(t *testing.T) *pgx.Conn {
	connString := "postgres://postgres:root@localhost:5432/testdb?sslmode=disable"
	db, err := pgx.Connect(context.Background(), connString)
	require.NoError(t, err)

	// Test uchun zarur jadval va ma'lumotlarni tayyorlash
	_, err = db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS orders (
		order_id SERIAL PRIMARY KEY,
		customer_id VARCHAR(255),
		order_details TEXT,
		delivery_address VARCHAR(255),
		payment_status VARCHAR(50),
		order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at BIGINT DEFAULT 0
	)`)
	require.NoError(t, err)

	return db
}

// teardownTestDB test tugagach ma'lumotlar bazasini o'chiradi.
func teardownTestDBOrder(t *testing.T, db *pgx.Conn) {
	_, err := db.Exec(context.Background(), `DROP TABLE IF EXISTS orders`)
	require.NoError(t, err)
}

// TestCreateOrder yangi buyurtma yozuvini yaratishni sinovdan o'tkazadi.
func TestCreateOrder(t *testing.T) {
	db := setupTestDBOrder(t)
	defer teardownTestDBOrder(t, db)

	orderStorage := cr.NewOrder(db)

	req := &genproto.CreateOrderRequest{
		CustomerId:     "customer-123",
		OrderDetails:   "Test order details",
		DeliveryAddress: "123 Test St",
		PaymentStatus:  genproto.PaymentStatus_PAID,
	}

	resp, err := orderStorage.CreateOrder(context.Background(), req)
	require.NoError(t, err)
	assert.True(t, resp.Success)
	assert.NotEmpty(t, resp.Order.OrderId)
	assert.Equal(t, req.CustomerId, resp.Order.CustomerId)
	assert.Equal(t, req.OrderDetails, resp.Order.OrderDetails)
	assert.Equal(t, req.DeliveryAddress, resp.Order.DeliveryAddress)
	assert.Equal(t, req.PaymentStatus, resp.Order.PaymentStatus)
}

// TestGetOrder buyurtma yozuvini olishni sinovdan o'tkazadi.
func TestGetOrder(t *testing.T) {
	db := setupTestDBOrder(t)
	defer teardownTestDBOrder(t, db)

	orderStorage := cr.NewOrder(db)

	// Avval yangi buyurtma yozuvini yaratish
	req := &genproto.CreateOrderRequest{
		CustomerId:     "customer-123",
		OrderDetails:   "Test order details",
		DeliveryAddress: "123 Test St",
		PaymentStatus:  genproto.PaymentStatus_PAID,
	}

	createResp, err := orderStorage.CreateOrder(context.Background(), req)
	require.NoError(t, err)

	// Yaratilgan buyurtma yozuvini olish
	getReq := &genproto.OrderRequest{OrderId: createResp.Order.OrderId}
	getResp, err := orderStorage.GetOrder(context.Background(), getReq)
	require.NoError(t, err)
	assert.True(t, getResp.Success)
	assert.Equal(t, req.CustomerId, getResp.Order.CustomerId)
	assert.Equal(t, req.OrderDetails, getResp.Order.OrderDetails)
	assert.Equal(t, req.DeliveryAddress, getResp.Order.DeliveryAddress)
	assert.Equal(t, req.PaymentStatus, getResp.Order.PaymentStatus)
}

// TestUpdateOrder buyurtma yozuvini yangilashni sinovdan o'tkazadi.
func TestUpdateOrder(t *testing.T) {
	db := setupTestDBOrder(t)
	defer teardownTestDBOrder(t, db)

	orderStorage := cr.NewOrder(db)

	// Avval yangi buyurtma yozuvini yaratish
	req := &genproto.CreateOrderRequest{
		CustomerId:     "customer-123",
		OrderDetails:   "Test order details",
		DeliveryAddress: "123 Test St",
		PaymentStatus:  genproto.PaymentStatus_PAID,
	}

	createResp, err := orderStorage.CreateOrder(context.Background(), req)
	require.NoError(t, err)

	// Yaratilgan buyurtma yozuvini yangilash
	updateReq := &genproto.UpdateOrderRequest{
		OrderId:        createResp.Order.OrderId,
		CustomerId:     "customer-456",
		OrderDetails:   "Updated order details",
		DeliveryAddress: "456 Updated St",
		PaymentStatus:  genproto.PaymentStatus_PENDING,
	}

	updateResp, err := orderStorage.UpdateOrder(context.Background(), updateReq)
	require.NoError(t, err)
	assert.True(t, updateResp.Success)
	assert.Equal(t, updateReq.CustomerId, updateResp.Order.CustomerId)
	assert.Equal(t, updateReq.OrderDetails, updateResp.Order.OrderDetails)
	assert.Equal(t, updateReq.DeliveryAddress, updateResp.Order.DeliveryAddress)
	assert.Equal(t, updateReq.PaymentStatus, updateResp.Order.PaymentStatus)
}

// TestDeleteOrder buyurtma yozuvini o'chirishni sinovdan o'tkazadi.
func TestDeleteOrder(t *testing.T) {
	db := setupTestDBOrder(t)
	defer teardownTestDBOrder(t, db)

	orderStorage := cr.NewOrder(db)

	// Avval yangi buyurtma yozuvini yaratish
	req := &genproto.CreateOrderRequest{
		CustomerId:     "customer-123",
		OrderDetails:   "Test order details",
		DeliveryAddress: "123 Test St",
		PaymentStatus:  genproto.PaymentStatus_PAID,
	}

	createResp, err := orderStorage.CreateOrder(context.Background(), req)
	require.NoError(t, err)

	// Yaratilgan buyurtma yozuvini o'chirish
	deleteReq := &genproto.OrderRequest{OrderId: createResp.Order.OrderId}
	deleteResp, err := orderStorage.DeleteOrder(context.Background(), deleteReq)
	require.NoError(t, err)
	assert.True(t, deleteResp.Success)

	// O'chirilgan buyurtma yozuvini qayta olishga urinish
	getReq := &genproto.OrderRequest{OrderId: createResp.Order.OrderId}
	_, err = orderStorage.GetOrder(context.Background(), getReq)
	assert.Error(t, err) // Buyurtma topilmadi xatolik kelishi kerak
}

// TestListOrders barcha buyurtma yozuvlarini olishni sinovdan o'tkazadi.
func TestListOrders(t *testing.T) {
	db := setupTestDBOrder(t)
	defer teardownTestDBOrder(t, db)

	orderStorage := cr.NewOrder(db)

	// Bir nechta buyurtma yozuvlarini yaratish
	for i := 0; i < 3; i++ {
		req := &genproto.CreateOrderRequest{
			CustomerId:     fmt.Sprintf("customer-%d", i),
			OrderDetails:   fmt.Sprintf("Test order details %d", i),
			DeliveryAddress: fmt.Sprintf("123 Test St %d", i),
			PaymentStatus:  genproto.PaymentStatus_PAID,
		}
		_, err := orderStorage.CreateOrder(context.Background(), req)
		require.NoError(t, err)
	}

	// Barcha buyurtma yozuvlarini olish
	listResp, err := orderStorage.ListOrders(context.Background(), &genproto.GetRecommendedOrdersRequest{})
	require.NoError(t, err)
	assert.Equal(t, 3, len(listResp.Orders))
}
