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
func setupDBConCart(t *testing.T) *postgres.Cart {
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
	return postgres.NewCart(db)
}

// Helper function to create a new cart for testing
func newCartTest() *gen.CreateCartRequest {
	return &gen.CreateCartRequest{
		UserId: "test-user-id",
	}
}

// Test function to create a cart
func TestCreateCart(t *testing.T) {
	cartDB := setupDBConCart(t)
	cartTest := newCartTest()

	cartRes, err := cartDB.Create(context.Background(), cartTest)
	if err != nil {
		t.Fatalf("Error creating cart: %v", err)
	}

	assert.True(t, cartRes.Success)
	assert.Equal(t, cartTest.UserId, cartRes.Cart.UserId)
	assert.NotEmpty(t, cartRes.Cart.CartId)
}

// Test function to get a cart
func TestGetCart(t *testing.T) {
	cartDB := setupDBConCart(t)
	cartTest := newCartTest()

	createdCartRes, err := cartDB.Create(context.Background(), cartTest)
	if err != nil {
		t.Fatalf("Error creating cart: %v", err)
	}

	req := &gen.CartRequest{CartId: createdCartRes.Cart.CartId}
	cartRes, err := cartDB.Get(context.Background(), req)
	if err != nil {
		t.Fatalf("Error getting cart: %v", err)
	}

	assert.True(t, cartRes.Success)
	assert.Equal(t, createdCartRes.Cart.CartId, cartRes.Cart.CartId)
	assert.Equal(t, createdCartRes.Cart.UserId, cartRes.Cart.UserId)
}

// Test function to update a cart
func TestUpdateCart(t *testing.T) {
	cartDB := setupDBConCart(t)
	cartTest := newCartTest()

	createdCartRes, err := cartDB.Create(context.Background(), cartTest)
	if err != nil {
		t.Fatalf("Error creating cart: %v", err)
	}

	// Update fields
	updateReq := &gen.UpdateCartRequest{
		CartId: createdCartRes.Cart.CartId,
		UserId: "updated-test-user-id",
	}
	updatedCartRes, err := cartDB.Update(context.Background(), updateReq)
	if err != nil {
		t.Fatalf("Error updating cart: %v", err)
	}

	assert.True(t, updatedCartRes.Success)
	assert.Equal(t, updateReq.CartId, updatedCartRes.Cart.CartId)
	assert.Equal(t, updateReq.UserId, updatedCartRes.Cart.UserId)
}

// Test function to delete a cart
func TestDeleteCart(t *testing.T) {
	cartDB := setupDBConCart(t)
	cartTest := newCartTest()

	createdCartRes, err := cartDB.Create(context.Background(), cartTest)
	if err != nil {
		t.Fatalf("Error creating cart: %v", err)
	}

	req := &gen.CartRequest{CartId: createdCartRes.Cart.CartId}
	_, err = cartDB.Delete(context.Background(), req)
	if err != nil {
		t.Fatalf("Error deleting cart: %v", err)
	}

	// Try to get the deleted cart
	getReq := &gen.CartRequest{CartId: createdCartRes.Cart.CartId}
	_, err = cartDB.Get(context.Background(), getReq)
	if err == nil {
		t.Fatalf("Expected error for deleted cart, but got nil")
	}
}

// Test function to list carts
func TestListCarts(t *testing.T) {
	cartDB := setupDBConCart(t)

	// Create a few carts for testing
	_, err := cartDB.Create(context.Background(), newCartTest())
	if err != nil {
		t.Fatalf("Error creating cart: %v", err)
	}

	resp, err := cartDB.List(context.Background(), &gen.Empty{})
	if err != nil {
		t.Fatalf("Error listing carts: %v", err)
	}

	assert.GreaterOrEqual(t, len(resp.Carts), 1)
}
