package main

import (
	"delivery/product_service/config"
	"delivery/product_service/internal/app"
)

func main() {
	// Configuration
	cfg := config.Load()

	// Run
	app.Run(cfg)
}
