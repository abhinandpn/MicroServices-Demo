// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/api"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/api/handler"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/config"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/database"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/repository"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServeHTTP, error) {
	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	productRepository := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductRepository(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)
	serveHTTP := http.NewServeHTTP(productHandler)
	return serveHTTP, nil
}
