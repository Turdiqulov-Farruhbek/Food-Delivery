package postgres

import (
	"context"
	"fmt"
	"log/slog"

	stg "product_ordering/storage"
	"product_ordering/config"

	"github.com/jackc/pgx/v5"
)

type StorageStruct struct {
	DB            *pgx.Conn
	CartItemS      stg.CartItemInterface
	CartS          stg.CartInterface
	OrderRecomendS stg.OrderRecomendInterface
	OrderS         stg.OrderInterface
	ProductS       stg.ProductInterface
}

func DbCon() (*StorageStruct, error) {
	var (
		db  *pgx.Conn
		err error
	)
	cfg := config.Load()
	dbcon := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	db, err = pgx.Connect(context.Background(), dbcon)
	if err != nil {
		slog.Error("Unable to connect to database:", "erroRRR:", err)
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		slog.Warn("unable to ping database:", "erroRRR:", err)
		return nil, err
	}

	return &StorageStruct{
		DB: db,
	}, nil
}

func (s *StorageStruct) CartItem() stg.CartItemInterface {
	if s.CartItemS == nil {
		s.CartItemS = NewCartItem(s.DB)
	}
	return s.CartItemS
}

func (s *StorageStruct) Cart() stg.CartInterface{
	if s.CartS == nil {
        s.CartS = NewCart(s.DB)
    }

	return s.CartS
}

func (s *StorageStruct) OrderRecomend() stg.OrderRecomendInterface{
	if s.OrderRecomendS == nil {
        s.OrderRecomendS = NewOrderRecomend(s.DB)
    }
	return s.OrderRecomendS
}

func (s *StorageStruct) Order() stg.OrderInterface{
	if s.OrderS == nil {
        s.OrderS = NewOrder(s.DB)
    }

	return s.OrderS
}

func (s *StorageStruct) Product() stg.ProductInterface{
	if s.ProductS == nil {
        s.ProductS = NewProduct(s.DB)
    }
	return s.ProductS
}