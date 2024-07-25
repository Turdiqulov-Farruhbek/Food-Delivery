package postgres

import (
	"Notification/config"
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	ntf "Notification/storage"
	
)

type StorageStruct struct {
	DB            *pgx.Conn
	AdminAlirt_S    ntf.AdminAlirtInterface
	CourierNtf_S    ntf.CourierNtfInterface
	UserNtf_S    ntf.UserNtfInterface
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

func (s *StorageStruct) AdminAlirt() ntf.AdminAlirtInterface {
	if s.AdminAlirt_S == nil {
		s.AdminAlirt_S = NewAdminAlirt(s.DB)
	}

    return s.AdminAlirt_S
}


func (s *StorageStruct) CourierNtf() ntf.CourierNtfInterface {
	if s.CourierNtf_S == nil {
        s.CourierNtf_S = NewCourierNtf(s.DB)
    }

    return s.CourierNtf_S
}


func (s *StorageStruct) UserNtf() ntf.UserNtfInterface {
	if s.UserNtf_S == nil {
        s.UserNtf_S = NewUserNtf(s.DB)
    }

    return s.UserNtf_S
}