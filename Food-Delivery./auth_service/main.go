package main

import (
	"log"

	"delivery/auth_service/api"
	"delivery/auth_service/api/handler"
	"delivery/auth_service/config"
	"delivery/auth_service/kafka"
	"delivery/auth_service/server"
	"delivery/auth_service/service"
	"delivery/auth_service/storage/postgres"
)

func main() {
	cfg := config.Load()

	go server.Server(&cfg)

	kafka, err := kafka.NewKafkaProducer([]string{cfg.KafkaUrl})
	if err != nil {
		log.Fatal(err)
		return
	}
	db, err := postgres.ConnectDB(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	auth_service := service.NewUserService(db)

	h := handler.NewHandler(auth_service, kafka)
	r := api.NewGin(h)
	r.Run(cfg.HTTPPort)
}
