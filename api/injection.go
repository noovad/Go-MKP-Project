//go:build wireinject
// +build wireinject

package api

import (
	"go-gin-project/api/controller"
	"go-gin-project/api/repository"
	"go-gin-project/api/service"
	"go-gin-project/config"

	"github.com/google/wire"
)

func InitializeTerminalController() *controller.TerminalController {
	wire.Build(
		controller.NewTerminalController,
		service.NewTerminalServiceImpl,
		repository.NewTerminalRepositoryImpl,
		config.DatabaseConnection,
		config.NewValidator,
	)
	return &controller.TerminalController{}
}
