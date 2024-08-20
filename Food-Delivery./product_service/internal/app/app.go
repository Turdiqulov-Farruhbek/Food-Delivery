package app

import (
	// "errors"
	"fmt"
	"log"
	"net"

	"delivery/product_service/config"
	pb "delivery/product_service/genproto"
	"delivery/product_service/kafka"
	"delivery/product_service/service"
	"delivery/product_service/storage/postgres"

	"google.golang.org/grpc"
)

func Run(cfg config.Config) {
	brokers := []string{cfg.KafkaHost + cfg.KafkaPort}

	//connect to database

	kf_producer, err := kafka.NewKafkaProducer(brokers)
	if err != nil {
		log.Fatal("Error with kf_producer", err)
		return
	}

	db, err := postgres.ConnectDB(&kf_producer, &cfg)
	if err != nil {
		log.Fatal("Error with dbconnection", err)
		return
	}

	//register kafka handlers
	err = Regitries(cfg, db)
	if err != nil {
		log.Fatal("Error with registries", err)
		return
	}

	//register services
	newServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(newServer, service.NewOrderService(db))
	pb.RegisterProductServiceServer(newServer, service.NewProductService(db))
	pb.RegisterItemServiceServer(newServer, service.NewOrderItemService(db))
	pb.RegisterKitchenServiceServer(newServer, service.NewKitchenService(db))
	pb.RegisterLocationServiceServer(newServer, service.NewLocationService(db))

	lis, err := net.Listen("tcp", cfg.HTTPPort)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running on :" + cfg.HTTPPort)
	err = newServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
