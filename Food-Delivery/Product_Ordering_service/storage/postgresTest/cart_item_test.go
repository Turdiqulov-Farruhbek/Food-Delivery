package postgrestest

import (
	"context"
	"fmt"
	"testing"

	cartItem "product_ordering/genproto"
	"product_ordering/storage/postgres"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

// Helper function to initialize PostgreSQL connection
func setupDBConnCartItem(t *testing.T) *postgres.CartItem {
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
	return postgres.NewCartItem(db)
}

// Helper function to create a new cart item for testing
func newCartItemTest() *cartItem.CreateCartItemRequest {
	return &cartItem.CreateCartItemRequest{
		CartId:   "fdcf6b6f-e9db-49e5-be7f-cf429c53b462",
		ProductId: "176dab25-69a9-4195-9699-3fa52232c8d6",
		Quantity:  1,
		Options:   `{"sss": "medium", "aacc": "red"}`,
	}
}

// Test function to create a cart item
func TestCreateCartItem(t *testing.T) {
	cartItemDB := setupDBConnCartItem(t)
	cartItemTest := newCartItemTest()

	cartItemRes, err := cartItemDB.Create(context.Background(), cartItemTest)
	if err != nil {
		t.Fatalf("Error creating cart item: %v", err)
	}

	assert.True(t, cartItemRes.Success)
	assert.Equal(t, cartItemTest.CartId, cartItemRes.CartItem.CartId)
	assert.Equal(t, cartItemTest.ProductId, cartItemRes.CartItem.ProductId)
	assert.Equal(t, cartItemTest.Quantity, cartItemRes.CartItem.Quantity)
	assert.JSONEq(t, cartItemTest.Options, cartItemRes.CartItem.Options) // JSON equality check
	assert.NotEmpty(t, cartItemRes.CartItem.CartItemId)
}

// Test function to get a cart item
func TestGetCartItem(t *testing.T) {
	cartItemDB := setupDBConnCartItem(t)
	cartItemTest := newCartItemTest()

	createdCartItemRes, err := cartItemDB.Create(context.Background(), cartItemTest)
	if err != nil {
		t.Fatalf("Error creating cart item: %v", err)
	}

	req := &cartItem.CartItemRequest{CartItemId: createdCartItemRes.CartItem.CartItemId}
	cartItemRes, err := cartItemDB.Get(context.Background(), req)
	if err != nil {
		t.Fatalf("Error getting cart item: %v", err)
	}

	assert.True(t, cartItemRes.Success)
	assert.Equal(t, createdCartItemRes.CartItem.CartItemId, cartItemRes.CartItem.CartItemId)
}

// Test function to update a cart item
func TestUpdateCartItem(t *testing.T) {
	cartItemDB := setupDBConnCartItem(t)
	cartItemTest := newCartItemTest()

	createdCartItemRes, err := cartItemDB.Create(context.Background(), cartItemTest)
	if err != nil {
		t.Fatalf("Error creating cart item: %v", err)
	}

	var CartItemId string
	// Update fields
	cartItemTest.CartId = "d613380d-65cf-40be-b48e-b386a7fd9514"
	cartItemTest.ProductId = "176dab25-69a9-4195-9699-3fa52232c8d6"
	cartItemTest.Quantity = 5
	cartItemTest.Options = `{"size": "large", "color": "blue"}` // JSON format
	CartItemId = createdCartItemRes.CartItem.CartItemId
	
	updateReq := &cartItem.UpdateCartItemRequest{
		CartItemId: CartItemId,
		CartId:     cartItemTest.CartId,
		ProductId:  cartItemTest.ProductId,
		Quantity:   cartItemTest.Quantity,
		Options:    cartItemTest.Options,
	}
	updatedCartItemRes, err := cartItemDB.Update(context.Background(), updateReq)
	if err != nil {
		t.Fatalf("Error updating cart item: %v", err)
	}

	assert.True(t, updatedCartItemRes.Success)
	assert.Equal(t, cartItemTest.CartId, updatedCartItemRes.CartItem.CartId)
	assert.Equal(t, cartItemTest.ProductId, updatedCartItemRes.CartItem.ProductId)
	assert.Equal(t, cartItemTest.Quantity, updatedCartItemRes.CartItem.Quantity)
	assert.JSONEq(t, cartItemTest.Options, updatedCartItemRes.CartItem.Options) // JSON equality check
}

// Test function to delete a cart item
func TestDeleteCartItem(t *testing.T) {
	cartItemDB := setupDBConnCartItem(t)
	cartItemTest := newCartItemTest()

	createdCartItemRes, err := cartItemDB.Create(context.Background(), cartItemTest)
	if err != nil {
		t.Fatalf("Error creating cart item: %v", err)
	}

	req := &cartItem.CartItemRequest{CartItemId: createdCartItemRes.CartItem.CartItemId}
	_,err = cartItemDB.Delete(context.Background(), req)
	if err != nil {
		t.Fatalf("Error deleting cart item: %v", err)
	}

	// O'chirilgan mahsulotni olishga harakat qilish
	getReq := &cartItem.CartItemRequest{
		CartItemId: createdCartItemRes.CartItem.CartItemId,
	}
	_, err = cartItemDB.Get(context.Background(), getReq)
	if err == nil {
		t.Fatalf("Expected error for deleted cart item, but got nil")
	}
}


// Test function to list cart items
func TestListCartItems(t *testing.T) {
	cartItemDB := setupDBConnCartItem(t)

	// Create a few cart items for testing
	_, err := cartItemDB.Create(context.Background(), newCartItemTest())
	if err != nil {
		t.Fatalf("Error creating cart item: %v", err)
	}

	_, err = cartItemDB.Create(context.Background(), newCartItemTest())
	if err != nil {
		t.Fatalf("Error creating cart item: %v", err)
	}

	resp, err := cartItemDB.List(context.Background(), &cartItem.Empty{})
	if err != nil {
		t.Fatalf("Error listing cart items: %v", err)
	}

	assert.GreaterOrEqual(t, len(resp.CartItems), 2)
}