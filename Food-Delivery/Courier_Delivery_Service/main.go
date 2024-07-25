package main

import (
	"log"
	"net"
	"path/filepath"
	"runtime"

	cfg "courier_delivery/config"
	"courier_delivery/config/logger"
	gen "courier_delivery/genproto"
	src "courier_delivery/service"
	post "courier_delivery/storage/postgres"

	"google.golang.org/grpc"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	logger := logger.NewLogger(basepath, "logs/info.log")

	db, err := post.DbCon()
	if err != nil {
		log.Fatal(err)
	}

	// config := cfg.Load()
	liss, err := net.Listen("tcp", cfg.Load().HTTPPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	gen.RegisterCourierOrderServiceServer(server, src.NewCourierOrderService(db, *logger))
	gen.RegisterCourierServiceServer(server, src.NewCourierService(db, *logger))
	gen.RegisterOrderServiceServer(server, src.NewOrderService(db, *logger))

	log.Println("Server is running on port :2020")
	if err := server.Serve(liss); err != nil {
		log.Fatalf("error listening: %v", err)
	}
}
