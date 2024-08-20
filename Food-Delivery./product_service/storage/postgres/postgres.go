package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"delivery/product_service/config"
	"delivery/product_service/kafka"
	"delivery/product_service/storage"
)

type Storage struct {
	db         *sql.DB
	ProductS   storage.ProductI
	OrderS     storage.OrderI
	OrderItemS storage.OrderItemI
	KitchenS   storage.KitchenI
	LocationS  storage.LocationI
	CourierS   storage.CourierI
}

func ConnectDB(kf_producer *kafka.KafkaProducer, cfg *config.Config) (*Storage, error) {

	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	prodS := NewProductRepo(db)
	orderS := NewOrderRepo(db, *kf_producer)
	orderItemS := NewOrderItemRepo(db)
	kitchenS := NewKitchenRepo(db)
	locationS := NewLocationRepo(db)
	courierS := NewCourierRepo(db)

	return &Storage{
		db:         db,
		ProductS:   prodS,
		OrderS:     orderS,
		OrderItemS: orderItemS,
		KitchenS:   kitchenS,
		LocationS:  locationS,
		CourierS:   courierS,
	}, nil
}
func (s *Storage) Product() storage.ProductI {
	if s.ProductS == nil {
		s.ProductS = NewProductRepo(s.db)
	}
	return s.ProductS
}
func (s *Storage) Order() storage.OrderI {
	return s.OrderS
}
func (s *Storage) OrderItem() storage.OrderItemI {
	if s.OrderItemS == nil {
		s.OrderItemS = NewOrderItemRepo(s.db)
	}
	return s.OrderItemS
}
func (s *Storage) Kitchen() storage.KitchenI {
	if s.KitchenS == nil {
		s.KitchenS = NewKitchenRepo(s.db)
	}
	return s.KitchenS
}
func (s *Storage) Location() storage.LocationI {
	if s.LocationS == nil {
		s.LocationS = NewLocationRepo(s.db)
	}
	return s.LocationS
}
func (s *Storage) Courier() storage.CourierI {
	return s.CourierS
}
