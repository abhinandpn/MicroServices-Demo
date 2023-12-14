package main

import (
	"log"

	config "github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/config"
	di "github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/di"
)

func main() {

	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		log.Fatal("cannot load config: ", cfgErr)
	}

	server, diErr := di.InitializeAPI(cfg)

	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Run()
	}
}
