//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/api"
	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/api/handler"
	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/database"
	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/repository"
	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/usecase"
	"github.com/google/wire"

	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/config"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(

		//database
		database.ConnectDatabase,

		//handler
		handler.NewUserUseHandler,

		//usecase
		usecase.NewUserUsecase,

		//repo
		repository.NewUserRepository,

		//
		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
