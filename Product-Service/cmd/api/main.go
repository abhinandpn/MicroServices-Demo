package main

import (
	"log"

	di "github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/di"

	config "github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/config"
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
