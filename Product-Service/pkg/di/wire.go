//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/api"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/api/handler"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/config"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/database"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/repository"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServeHTTP, error) {
	wire.Build(

		//database
		database.ConnectDatabase,

		//handler
		handler.NewProductHandler,

		//usecase
		usecase.NewProductRepository,

		//repository
		repository.NewProductRepository,

		http.NewServeHTTP,
	)
	return &http.ServeHTTP{}, nil
}
