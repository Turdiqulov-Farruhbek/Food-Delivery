package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"courier_delivery/config"
	stg "courier_delivery/storage"

	"github.com/jackc/pgx/v5"
)

type StorageStruct struct {
	DB                *pgx.Conn
	CourierOrder_S    stg.CourierOrderInterface
	Courier_S         stg.CourierInterface
	Order_S           stg.OrderInterface
	Task_S            stg.TaskInterface
	CourierLocation_S stg.CourierLocationInterface
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

func (s *StorageStruct) CourierOrder() stg.CourierOrderInterface {
	if s.CourierOrder_S == nil {
		s.CourierOrder_S = NewCourierOrder(s.DB)
	}
	return s.CourierOrder_S
}

func (s *StorageStruct) Courier() stg.CourierInterface {
	if s.Courier_S == nil {
		s.Courier_S = NewCourier(s.DB)
	}
	return s.Courier_S
}

func (s *StorageStruct) Order() stg.OrderInterface {
	if s.Order_S == nil {
		s.Order_S = NewOrder(s.DB)
	}
	return s.Order_S
}


func (s *StorageStruct) Task() stg.TaskInterface {
	if s.Task_S == nil {
        s.Task_S = NewTask(s.DB)
    }
    return s.Task_S
}

func (s *StorageStruct) CourierLocation() stg.CourierLocationInterface {
	if s.CourierLocation_S == nil {
        s.CourierLocation_S = NewCourierLocation(s.DB)
    }
    return s.CourierLocation_S
}