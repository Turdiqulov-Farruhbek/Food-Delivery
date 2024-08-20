package main

import (
	"delivery/gateway/api"
	"delivery/gateway/api/handlers"
	"delivery/gateway/config"
)

func main() {
	cfg := config.Load()

	h, err := handlers.NewHandler(&cfg)
	if err != nil {
		panic(err)
	}

	r := api.NewGin(h)

	r.Run(cfg.HTTPPort)
}
