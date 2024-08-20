package postgres

import (
	"database/sql"
	"fmt"

	"delivery/auth_service/config"
	"delivery/auth_service/storage"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

type Storage struct {
	db    *sql.DB
	rdb   *redis.Client
	UserS storage.UserI
}

func ConnectDB(cfg *config.Config) (*Storage, error) {
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
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisUrl,
	})

	userS := NewUserRepo(db, rdb)
	return &Storage{
		db:    db,
		rdb:   rdb,
		UserS: userS,
	}, nil
}
func (s *Storage) User() storage.UserI {
	if s.UserS == nil {
		s.UserS = NewUserRepo(s.db, s.rdb)
	}
	return s.UserS
}
