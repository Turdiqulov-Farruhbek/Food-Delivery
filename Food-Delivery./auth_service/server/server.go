package server

import (
	"fmt"
	"log"
	"net"

	// "time"

	"delivery/auth_service/config"
	pb "delivery/auth_service/genproto"
	"delivery/auth_service/kafka"
	"delivery/auth_service/service"
	"delivery/auth_service/storage/postgres"

	"google.golang.org/grpc"
)

func Server(cfg *config.Config) {
	db, err := postgres.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Error with dbconnection", err)
		return
	}

	userService := service.NewUserService(db)

	//*kafka\\//\\
	// time.Sleep(time.Second * 20)
	brokers := []string{cfg.KafkaUrl} //////////////////////////////////////////////////////

	kcm := kafka.NewKafkaConsumerManager()

	if err := kcm.RegisterConsumer(brokers, "user-create", "user", kafka.UserCreateHandler(userService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'user-create' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}
	if err := kcm.RegisterConsumer(brokers, "forgot_password", "user", kafka.ForgotPasswordHandler(userService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'forgot_password' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}
	if err := kcm.RegisterConsumer(brokers, "courier-create", "courier", kafka.CourierCreateHandler(userService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'courier-create' already exists")
		} else {
			log.Fatalf("Error registering courier: %v", err)
		}
	}
	if err := kcm.RegisterConsumer(brokers, "kitchen-manager-create", "kitchen-manager", kafka.ManagerCreateHandler(userService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic'kitchen-manager-create' already exists")
		} else {
			log.Fatalf("Error registering kitchen-manager: %v", err)
		}
	}

	// *kafka\\//\\

	newServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(newServer, service.NewUserService(db))

	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running on ", cfg.GRPCPort)
	err = newServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
