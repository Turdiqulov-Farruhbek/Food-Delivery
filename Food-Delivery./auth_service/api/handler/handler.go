package handler

import (
	"delivery/auth_service/kafka"
	"delivery/auth_service/service"
)

type Handler struct {
	Producer kafka.KafkaProducer
	AuthS    *service.UserService
}

func NewHandler(auth *service.UserService, kafka kafka.KafkaProducer) *Handler {
	return &Handler{AuthS: auth, Producer: kafka}
}
