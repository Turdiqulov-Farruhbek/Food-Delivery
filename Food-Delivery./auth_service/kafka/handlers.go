package kafka

import (
	"context"
	"log"

	pb "delivery/auth_service/genproto"
	"delivery/auth_service/service"

	"google.golang.org/protobuf/encoding/protojson"
)

func UserCreateHandler(userService *service.UserService) func(message []byte) {
	return func(message []byte) {
		var user pb.UserCreateReq
		if err := protojson.Unmarshal(message, &user); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := userService.Register(context.Background(), &user)
		if err != nil {
			log.Printf("Cannot create user via Kafka: %v", err)
			return
		}
		log.Printf("Created user: %+v", res)
	}
}

func ForgotPasswordHandler(userService *service.UserService) func(message []byte) {
	return func(message []byte) {
		var email pb.ForgotPasswordReq
		if err := protojson.Unmarshal(message, &email); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := userService.ForgotPassword(context.Background(), &email)
		if err != nil {
			log.Printf("Cannot send forgot password via Kafka: %v", err)
			return
		}
		log.Printf("Sent forgot password email: %+v", res)
	}
}
func CourierCreateHandler(userService *service.UserService) func(message []byte) {
	return func(message []byte) {
		var user pb.UserCreateReq
		if err := protojson.Unmarshal(message, &user); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := userService.RegisterCourier(context.Background(), &user)
		if err != nil {
			log.Printf("Cannot create courier via Kafka: %v", err)
			return
		}
		log.Printf("Created courier: %+v", res)
	}
}
func ManagerCreateHandler(userService *service.UserService) func(message []byte) {
	return func(message []byte) {
		var user pb.UserCreateReq
		if err := protojson.Unmarshal(message, &user); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := userService.CreateManager(context.Background(), &user)
		if err != nil {
			log.Printf("Cannot create manager via Kafka: %v", err)
			return
		}
		log.Printf("Created manager: %+v", res)
	}
}
