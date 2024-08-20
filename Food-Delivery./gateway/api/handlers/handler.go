package handlers

import (
	"delivery/gateway/clients"
	"delivery/gateway/config"
	"delivery/gateway/kafka"
	m "delivery/gateway/minio"

	"github.com/minio/minio-go/v7"
	// "github.com/minio/minio-go/v7/pkg/credentials"
)

type Handler struct {
	Clients  *clients.Clients
	Producer kafka.KafkaProducer
	minio    *minio.Client
}

func NewHandler(cfg *config.Config) (*Handler, error) {
	// grpc clients
	clients, err := clients.NewClients(cfg)
	if err != nil {
		return nil, err
	}

	// kafka producer
	kafkaProd, err := kafka.NewKafkaProducer([]string{cfg.KafkaHost + cfg.KafkaPort})
	if err != nil {
		return nil, err
	}
	// connect to minio
	minio_c, err := m.MinIOConnect()
	if err != nil {
		return nil, err
	}

	return &Handler{
		Clients:  clients,
		Producer: kafkaProd,
		minio:    minio_c,
	}, nil
}
