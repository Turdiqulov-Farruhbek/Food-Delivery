package postgrestest


import (
	"context"
	"fmt"
	"testing"

	gen "product_ordering/genproto/product"
	"product_ordering/storage/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/jackc/pgx/v5"
)

// Yordamchi funksiya PostgreSQL ulanishini o'rnatish uchun
func setupDBConProduct(t *testing.T) *postgres.Product {
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
	return postgres.NewProduct(db)
}

// Test funksiyasi Create uchun
func TestCreateProduct(t *testing.T) {
	product := setupDBConProduct(t)
	req := &gen.CreateProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       100.0,
		ImageUrl:    "http://example.com/image.jpg",
	}
	resp, err := product.Create(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
}

// Test funksiyasi Get uchun
func TestGetProduct(t *testing.T) {
	product := setupDBConProduct(t)
	req := &gen.ProductRequest{ProductId: "1"}
	resp, err := product.Get(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
}

// Test funksiyasi Update uchun
func TestUpdateProduct(t *testing.T) {
	product := setupDBConProduct(t)
	req := &gen.UpdateProductRequest{
		ProductId:   "1",
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       150.0,
		ImageUrl:    "http://example.com/updated_image.jpg",
	}
	resp, err := product.Update(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
}

// Test funksiyasi Delete uchun
func TestDeleteProduct(t *testing.T) {
	product := setupDBConProduct(t)
	req := &gen.ProductRequest{ProductId: "1"}
	resp, err := product.Delete(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
}

// Test funksiyasi List uchun
func TestListProducts(t *testing.T) {
	product := setupDBConProduct(t)
	resp, err := product.List(context.Background(), &gen.Empty{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Greater(t, len(resp.Products), 0)
}
