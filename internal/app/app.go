package app

import (
	"errors"
	"fmt"
	"gateway/internal/config"
	"gateway/internal/controllers"
	"gateway/pkg/middleware"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type App struct {
	app    *fiber.App
	config *config.Config
	logger *slog.Logger

	AuthController *controllers.AuthController
}

func newApp(
	config *config.Config,
	log *slog.Logger,
	authController *controllers.AuthController,
) *App {
	app := fiber.New(fiber.Config{
		AppName:       "sochya-gateway",
		CaseSensitive: true,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			err = ctx.Status(code).JSON(fiber.Map{
				"message": e.Message,
			})

			return nil
		},
		BodyLimit: 10 << 20,
	})

	return &App{
		app:            app,
		config:         config,
		logger:         log,
		AuthController: authController,
	}
}

func (a *App) Run() error {
	host := a.config.App.Host
	port := a.config.App.Port

	a.app.Use(logger.New())
	a.app.Use(middleware.AttachRequestId())
	a.app.Use(middleware.AttachLogger(a.logger))

	a.app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://localhost:3001, http://localhost:3001",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	v1 := a.app.Group("/api")

	au := v1.Group("/auth")
	au.Post("/sign-in", a.AuthController.SignIn())
	au.Post("/sign-up", a.AuthController.SignUp())
	au.Post("/sign-out", a.AuthController.AuthRequired(), a.AuthController.SignOut())
	au.Post("/refresh", a.AuthController.Refresh())
	v1.Get("/profile", a.AuthController.AuthRequired(), a.AuthController.Profile())

	a.logger.Info("server started", slog.String("host", host), slog.Int("port", port))
	return a.app.Listen(fmt.Sprintf("%s:%d", host, port))
}

func (a *App) Shutdown() {
	a.app.Shutdown()
}
