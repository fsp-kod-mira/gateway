package app

import (
	"errors"
	"fmt"
	"gateway/api/auth"
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

	AuthController     *controllers.AuthController
	CvController       *controllers.CvController
	TemplateController *controllers.TemplatesController
}

func newApp(
	config *config.Config,
	log *slog.Logger,
	authController *controllers.AuthController,
	cvController *controllers.CvController,
	templateController *controllers.TemplatesController,
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
		app:                app,
		config:             config,
		logger:             log,
		AuthController:     authController,
		CvController:       cvController,
		TemplateController: templateController,
	}
}

func (a *App) Run() error {
	host := a.config.App.Host
	port := a.config.App.Port

	a.app.Use(logger.New())
	a.app.Use(middleware.AttachRequestId())
	a.app.Use(middleware.AttachLogger(a.logger))

	a.logger.Info("setup CORS", slog.String("AllowOrigins", a.config.CORS.Origin))

	a.app.Use(cors.New(cors.Config{
		// AllowOrigins:     "https://localhost:3001, http://localhost:3001",
		AllowOrigins:     a.config.CORS.Origin,
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

	cv := v1.Group("/cv")
	cv.Post("/", a.AuthController.AuthRequired(auth.Role_recruiter, auth.Role_hiring_manager), a.CvController.Upload())
	cv.Get("/", a.AuthController.AuthRequired(auth.Role_hiring_manager, auth.Role_resource_manager), a.CvController.GetAll())
	cv.Get("/:id", a.AuthController.AuthRequired(), a.CvController.Get())

	v1.Get("/templates", a.AuthController.AuthRequired(auth.Role_resource_manager, auth.Role_hiring_manager), a.TemplateController.GetAll())
	v1.Get("/features", a.AuthController.AuthRequired(auth.Role_resource_manager, auth.Role_hiring_manager), a.TemplateController.GetFeaturesByTemplate())

	a.logger.Info("server started", slog.String("host", host), slog.Int("port", port))
	return a.app.Listen(fmt.Sprintf("%s:%d", host, port))
}

func (a *App) Shutdown() {
	a.app.Shutdown()
}
