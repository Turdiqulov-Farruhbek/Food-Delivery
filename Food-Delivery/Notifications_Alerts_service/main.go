package main

import (
	cfg "Notification/config"
	er "Notification/genproto"
	src "Notification/services"
	post "Notification/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

// var (
// 	_, b, _, _ = runtime.Caller(0)
// 	basepath   = filepath.Dir(b)
// )

func main() {
	// logger := logger.NewLogger(basepath, "logs/info.log")
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

	er.RegisterCourierNotificationServiceServer(server, src.NewCourierNotificationService(db))
	er.RegisterUserNotificationServiceServer(server, src.NewUserNotificationService(db))

	log.Println("Server is running on port :4040")
	if err := server.Serve(liss); err != nil {
		log.Fatalf("error listening: %v", err)
	}

}