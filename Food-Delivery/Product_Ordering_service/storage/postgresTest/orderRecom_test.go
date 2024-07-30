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
func setupDBConOrderRecommendation(t *testing.T) *postgres.OrderRecommendation {
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
	return postgres.NewOrderRecomend(db)
}

// Helper function to create a new order recommendation for testing
func newOrderRecommendationTest() *gen.CreateOrderRecommendationRequest {
	return &gen.CreateOrderRecommendationRequest{
		OrderId:       "test-order-id",
		CourierId:     "test-courier-id",
		RecommendedAt: "2024-07-05T00:00:00Z",
		Status:        gen.RecommendationStatus_PENDING,
	}
}


// Test function to create an order recommendation
func TestCreateOrderRecommendation(t *testing.T) {
	orderRecommendationDB := setupDBConOrderRecommendation(t)
	orderRecommendationTest := newOrderRecommendationTest()

	orderRecommendationRes, err := orderRecommendationDB.Create(context.Background(), orderRecommendationTest)
	if err != nil {
		t.Fatalf("Error creating order recommendation: %v", err)
	}

	assert.True(t, orderRecommendationRes.Success)
	assert.Equal(t, orderRecommendationTest.OrderId, orderRecommendationRes.OrderRecommendation.OrderId)
	assert.Equal(t, orderRecommendationTest.CourierId, orderRecommendationRes.OrderRecommendation.CourierId)
	assert.Equal(t, orderRecommendationTest.RecommendedAt, orderRecommendationRes.OrderRecommendation.RecommendedAt)
	assert.Equal(t, orderRecommendationTest.Status, orderRecommendationRes.OrderRecommendation.Status)
	assert.NotEmpty(t, orderRecommendationRes.OrderRecommendation.RecommendationId)
}

// Test function to get an order recommendation
func TestGetOrderRecommendation(t *testing.T) {
	orderRecommendationDB := setupDBConOrderRecommendation(t)
	orderRecommendationTest := newOrderRecommendationTest()

	createdOrderRecommendationRes, err := orderRecommendationDB.Create(context.Background(), orderRecommendationTest)
	if err != nil {
		t.Fatalf("Error creating order recommendation: %v", err)
	}

	req := &gen.OrderRecommendationRequest{RecommendationId: createdOrderRecommendationRes.OrderRecommendation.RecommendationId}
	orderRecommendationRes, err := orderRecommendationDB.Get(context.Background(), req)
	if err != nil {
		t.Fatalf("Error getting order recommendation: %v", err)
	}

	assert.True(t, orderRecommendationRes.Success)
	assert.Equal(t, createdOrderRecommendationRes.OrderRecommendation.RecommendationId, orderRecommendationRes.OrderRecommendation.RecommendationId)
	assert.Equal(t, createdOrderRecommendationRes.OrderRecommendation.OrderId, orderRecommendationRes.OrderRecommendation.OrderId)
	assert.Equal(t, createdOrderRecommendationRes.OrderRecommendation.CourierId, orderRecommendationRes.OrderRecommendation.CourierId)
	assert.Equal(t, createdOrderRecommendationRes.OrderRecommendation.RecommendedAt, orderRecommendationRes.OrderRecommendation.RecommendedAt)
	assert.Equal(t, createdOrderRecommendationRes.OrderRecommendation.Status, orderRecommendationRes.OrderRecommendation.Status)
}

// Test function to update an order recommendation
func TestUpdateOrderRecommendation(t *testing.T) {
	orderRecommendationDB := setupDBConOrderRecommendation(t)
	orderRecommendationTest := newOrderRecommendationTest()

	createdOrderRecommendationRes, err := orderRecommendationDB.Create(context.Background(), orderRecommendationTest)
	if err != nil {
		t.Fatalf("Error creating order recommendation: %v", err)
	}

	// Update fields
	updateReq := &gen.UpdateOrderRecommendationRequest{
		RecommendationId: createdOrderRecommendationRes.OrderRecommendation.RecommendationId,
		OrderId:          "updated-order-id",
		CourierId:        "updated-courier-id",
		RecommendedAt:    "2024-07-06T00:00:00Z",
		Status:           gen.RecommendationStatus_ACCEPTED,
	}
	updatedOrderRecommendationRes, err := orderRecommendationDB.Update(context.Background(), updateReq)
	if err != nil {
		t.Fatalf("Error updating order recommendation: %v", err)
	}

	assert.True(t, updatedOrderRecommendationRes.Success)
	assert.Equal(t, updateReq.RecommendationId, updatedOrderRecommendationRes.OrderRecommendation.RecommendationId)
	assert.Equal(t, updateReq.OrderId, updatedOrderRecommendationRes.OrderRecommendation.OrderId)
	assert.Equal(t, updateReq.CourierId, updatedOrderRecommendationRes.OrderRecommendation.CourierId)
	assert.Equal(t, updateReq.RecommendedAt, updatedOrderRecommendationRes.OrderRecommendation.RecommendedAt)
	assert.Equal(t, updateReq.Status, updatedOrderRecommendationRes.OrderRecommendation.Status)
}

// Test function to delete an order recommendation
func TestDeleteOrderRecommendation(t *testing.T) {
	orderRecommendationDB := setupDBConOrderRecommendation(t)
	orderRecommendationTest := newOrderRecommendationTest()

	createdOrderRecommendationRes, err := orderRecommendationDB.Create(context.Background(), orderRecommendationTest)
	if err != nil {
		t.Fatalf("Error creating order recommendation: %v", err)
	}

	req := &gen.OrderRecommendationRequest{RecommendationId: createdOrderRecommendationRes.OrderRecommendation.RecommendationId}
	_, err = orderRecommendationDB.Delete(context.Background(), req)
	if err != nil {
		t.Fatalf("Error deleting order recommendation: %v", err)
	}

	// Try to get the deleted order recommendation
	getReq := &gen.OrderRecommendationRequest{RecommendationId: createdOrderRecommendationRes.OrderRecommendation.RecommendationId}
	_, err = orderRecommendationDB.Get(context.Background(), getReq)
	if err == nil {
		t.Fatalf("Expected error for deleted order recommendation, but got nil")
	}
}

// Test function to list order recommendations
func TestListOrderRecommendations(t *testing.T) {
	orderRecommendationDB := setupDBConOrderRecommendation(t)

	// Create a few order recommendations for testing
	_, err := orderRecommendationDB.Create(context.Background(), newOrderRecommendationTest())
	if err != nil {
		t.Fatalf("Error creating order recommendation: %v", err)
	}

	resp, err := orderRecommendationDB.List(context.Background(), &gen.Empty{})
	if err != nil {
		t.Fatalf("Error listing order recommendations: %v", err)
	}

	assert.GreaterOrEqual(t, len(resp.OrderRecommendations), 1)
}
