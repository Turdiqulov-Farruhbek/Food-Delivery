package postgrestest

import (
	"context"
	"fmt"
	"testing"

	"product_ordering/storage/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/jackc/pgx/v5"
)

// Yordamchi funksiya PostgreSQL ulanishini o'rnatish uchun
func setupDBCon(t *testing.T) *postgres.StorageStruct {
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
	return &postgres.StorageStruct{DB: db}
}

// Test funksiyasi DbCon uchun
func TestDbCon(t *testing.T) {
	storage, err := postgres.DbCon()
	assert.NoError(t, err)
	assert.NotNil(t, storage)
}

// Test funksiyasi CartItem uchun
func TestCartItem(t *testing.T) {
	storage := setupDBCon(t)
	cartItem := storage.CartItem()
	assert.NotNil(t, cartItem)
}

// Test funksiyasi Cart uchun
func TestCart(t *testing.T) {
	storage := setupDBCon(t)
	cart := storage.Cart()
	assert.NotNil(t, cart)
}

// Test funksiyasi OrderRecomend uchun
func TestOrderRecomend(t *testing.T) {
	storage := setupDBCon(t)
	orderRecomend := storage.OrderRecomend()
	assert.NotNil(t, orderRecomend)
}

// Test funksiyasi Order uchun
func TestOrder(t *testing.T) {
	storage := setupDBCon(t)
	order := storage.Order()
	assert.NotNil(t, order)
}

// Test funksiyasi Product uchun
func TestProduct(t *testing.T) {
	storage := setupDBCon(t)
	product := storage.Product()
	assert.NotNil(t, product)
}
