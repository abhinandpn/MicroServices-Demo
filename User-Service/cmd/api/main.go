package main

import (
	"log"

	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/config"
	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/di"
)

func main() {
	cfg, configErr := config.LoadConfig()

	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(cfg)

	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
