//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/api"
	"github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/api/handler"
	"github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/config"
	"github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/database"
	"github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/repository"
	"github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(

		// database
		database.ConnectDatabase,

		// handler
		handler.NewAdminHandler,

		// usecase
		usecase.NewAdminUseCase,

		// repo
		repository.NewAdminRepository,

		// http
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
