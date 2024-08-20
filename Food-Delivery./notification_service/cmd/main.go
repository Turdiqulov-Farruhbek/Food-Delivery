package main

import (
	"delivery/notification_service/config"
	"delivery/notification_service/server"
)

func main() {
	//configurations
	cfg := config.Load()

	// Run
	server.Run(&cfg)
}
