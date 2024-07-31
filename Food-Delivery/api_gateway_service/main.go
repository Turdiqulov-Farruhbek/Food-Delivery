package main

import (
	"fmt"
	"gateway/api"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up gRPC connections
	courierService, err := grpc.NewClient("localhost:2020", grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		slog.Error("Failed to connect to courierService service %v", "error", err)
	}
	defer courierService.Close()

	productService, err := grpc.NewClient("localhost:2030", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!= nil {
        slog.Error("Failed to connect to productService service %v", "error", err)
    }
	defer productService.Close()

	notificationService, err := grpc.NewClient("localhost:2040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!= nil {
        slog.Error("Failed to connect to notificationService service %v", "error", err)
    }
	defer notificationService.Close()


	router := api.NewGin(courierService, productService, notificationService)
	fmt.Println("API Gateway running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		slog.Error("Failed to connect to gin engine: %v", "error",  err)
	}

}