//go:build wireinject
// +build wireinject

package app

import (
	"gateway/internal/config"
	"gateway/internal/controllers"
	"gateway/internal/logger"
	"gateway/internal/service/authservice"

	"github.com/google/wire"
)

func InitApp() *App {
	panic(wire.Build(
		newApp,
		wire.NewSet(logger.New),
		wire.NewSet(config.New),

		wire.NewSet(authservice.New),
		wire.Bind(new(controllers.AuthService), new(*authservice.Service)),
		wire.Bind(new(controllers.UserService), new(*authservice.Service)),
		wire.NewSet(controllers.NewAuthController),
	))
}
