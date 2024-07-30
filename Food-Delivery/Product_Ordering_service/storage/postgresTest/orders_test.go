package postgrestest

import (
	"context"
	"fmt"
	"testing"

	gen "product_ordering/genproto/product"
	"product_ordering/storage/postgres"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

// Helper function to initialize PostgreSQL connection
func setupDBConOrder(t *testing.T) *postgres.Order {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"postgres",
		"root",
		"localhost",
		5432,
		"food_delivery_db",
	)

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return postgres.NewOrder(db)
}

// Helper function to create a new order for testing
func newOrderTest() *gen.CreateOrderProductRequest {
	return &gen.CreateOrderProductRequest{
		UserId:     "test-user-id",
		Status:     gen.CardStatus_OR_PENDING,
		TotalPrice: 100.0,
	}
}

// Test function to create an order
func TestCreateOrder(t *testing.T) {
	orderDB := setupDBConOrder(t)
	orderTest := newOrderTest()

	orderRes, err := orderDB.Create(context.Background(), orderTest)
	if err != nil {
		t.Fatalf("Error creating order: %v", err)
	}

	assert.True(t, orderRes.Success)
	assert.Equal(t, orderTest.UserId, orderRes.Order.UserId)
	assert.Equal(t, orderTest.Status, orderRes.Order.Status)
	assert.Equal(t, orderTest.TotalPrice, orderRes.Order.TotalPrice)
	assert.NotEmpty(t, orderRes.Order.OrderId)
}

// Test function to get an order
func TestGetOrder(t *testing.T) {
	orderDB := setupDBConOrder(t)
	orderTest := newOrderTest()

	createdOrderRes, err := orderDB.Create(context.Background(), orderTest)
	if err != nil {
		t.Fatalf("Error creating order: %v", err)
	}

	req := &gen.OrderProductRequest{OrderId: createdOrderRes.Order.OrderId}
	orderRes, err := orderDB.Get(context.Background(), req)
	if err != nil {
		t.Fatalf("Error getting order: %v", err)
	}

	assert.True(t, orderRes.Success)
	assert.Equal(t, createdOrderRes.Order.OrderId, orderRes.Order.OrderId)
	assert.Equal(t, createdOrderRes.Order.UserId, orderRes.Order.UserId)
	assert.Equal(t, createdOrderRes.Order.Status, orderRes.Order.Status)
	assert.Equal(t, createdOrderRes.Order.TotalPrice, orderRes.Order.TotalPrice)
}

// Test function to update an order
func TestUpdateOrder(t *testing.T) {
	orderDB := setupDBConOrder(t)
	orderTest := newOrderTest()

	createdOrderRes, err := orderDB.Create(context.Background(), orderTest)
	if err != nil {
		t.Fatalf("Error creating order: %v", err)
	}

	// Update fields
	updateReq := &gen.UpdateOrderProductRequest{
		OrderId:    createdOrderRes.Order.OrderId,
		Status:     gen.CardStatus_COMPLETED,
		TotalPrice: 150.0,
	}
	updatedOrderRes, err := orderDB.Update(context.Background(), updateReq)
	if err != nil {
		t.Fatalf("Error updating order: %v", err)
	}

	assert.True(t, updatedOrderRes.Success)
	assert.Equal(t, updateReq.OrderId, updatedOrderRes.Order.OrderId)
	assert.Equal(t, updateReq.Status, updatedOrderRes.Order.Status)
	assert.Equal(t, updateReq.TotalPrice, updatedOrderRes.Order.TotalPrice)
}

// Test function to delete an order
func TestDeleteOrder(t *testing.T) {
	orderDB := setupDBConOrder(t)
	orderTest := newOrderTest()

	createdOrderRes, err := orderDB.Create(context.Background(), orderTest)
	if err != nil {
		t.Fatalf("Error creating order: %v", err)
	}

	req := &gen.OrderProductRequest{OrderId: createdOrderRes.Order.OrderId}
	_, err = orderDB.Delete(context.Background(), req)
	if err != nil {
		t.Fatalf("Error deleting order: %v", err)
	}

	// Try to get the deleted order
	getReq := &gen.OrderProductRequest{OrderId: createdOrderRes.Order.OrderId}
	_, err = orderDB.Get(context.Background(), getReq)
	if err == nil {
		t.Fatalf("Expected error for deleted order, but got nil")
	}
}

// Test function to list orders
func TestListOrders(t *testing.T) {
	orderDB := setupDBConOrder(t)

	// Create a few orders for testing
	_, err := orderDB.Create(context.Background(), newOrderTest())
	if err != nil {
		t.Fatalf("Error creating order: %v", err)
	}

	resp, err := orderDB.List(context.Background(), &gen.OrderProductListRequest{})
	if err != nil {
		t.Fatalf("Error listing orders: %v", err)
	}

	assert.GreaterOrEqual(t, len(resp.Orders), 1)
}
