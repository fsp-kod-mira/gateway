// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"fmt"
	"gateway/api/cv"
	"gateway/internal/config"
	"gateway/internal/controllers"
	"gateway/internal/logger"
	"gateway/internal/service/authservice"
	"gateway/internal/service/cvservice"
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

// Injectors from wire.go:

func InitApp() (*App, func(), error) {
	configConfig := config.New()
	slogLogger := logger.New(configConfig)
	service := authservice.New(configConfig, slogLogger)
	authController := controllers.NewAuthController(service, service)
	cvServiceClient, cleanup, err := InitCVService(configConfig, slogLogger)
	if err != nil {
		return nil, nil, err
	}
	cvserviceService := cvservice.New(cvServiceClient, slogLogger)
	cvController := controllers.NewCvController(cvserviceService)
	app := newApp(configConfig, slogLogger, authController, cvController)
	return app, func() {
		cleanup()
	}, nil
}

// wire.go:

func InitCVService(cfg *config.Config, logger2 *slog.Logger) (cv.CvServiceClient, func(), error) {
	l := logger2.With("service", "auth")
	host := cfg.Services.CvService.Host
	port := cfg.Services.CvService.Port

	l.Info("connecting to grpc service", slog.String("host", host), slog.Int("port", port))

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("error with connection to grpc service", slog.String("err", err.Error()))
		return nil, nil, err
	}

	client := cv.NewCvServiceClient(conn)
	return client, func() {
		conn.Close()
	}, nil
}
