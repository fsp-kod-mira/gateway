//go:build wireinject
// +build wireinject

package app

import (
	"fmt"
	"gateway/api/cv"
	objectstoragev1 "gateway/api/objectstorage"
	"gateway/api/templates"
	"gateway/internal/config"
	"gateway/internal/controllers"
	"gateway/internal/logger"
	"gateway/internal/service/authservice"
	"gateway/internal/service/cvservice"
	"gateway/internal/service/objectstorage"
	"gateway/internal/service/templatesservice"
	"log/slog"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitApp() (*App, func(), error) {
	panic(wire.Build(
		newApp,
		InitCVService,
		InitObjectStorageService,
		InitTemplatesService,

		wire.NewSet(logger.New),
		wire.NewSet(config.New),

		wire.NewSet(authservice.New),
		wire.Bind(new(controllers.AuthService), new(*authservice.Service)),
		wire.Bind(new(controllers.UserService), new(*authservice.Service)),

		wire.NewSet(cvservice.New),
		wire.Bind(new(controllers.CvService), new(*cvservice.Service)),

		wire.NewSet(templatesservice.New),
		wire.Bind(new(controllers.TemplatesService), new(*templatesservice.Service)),

		wire.NewSet(objectstorage.New),
		wire.Bind(new(controllers.ObjectStorage), new(*objectstorage.ObjectStorage)),

		wire.NewSet(controllers.NewAuthController),
		wire.NewSet(controllers.NewCvController),
		wire.NewSet(controllers.NewTemplatesController),
	))
}

func InitCVService(cfg *config.Config, logger *slog.Logger) (cv.CvServiceClient, func(), error) {
	l := logger.With("service", "auth")
	host := cfg.Services.CvService.Host
	port := cfg.Services.CvService.Port

	l.Info("connecting to grpc service", slog.String("host", host), slog.Int("port", port))

	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Error("error with connection to grpc service", slog.String("err", err.Error()))
		return nil, nil, err
	}

	client := cv.NewCvServiceClient(conn)
	return client, func() {
		conn.Close()
	}, nil
}

func InitObjectStorageService(cfg *config.Config, logger *slog.Logger) (objectstoragev1.ObjectStorageClient, func(), error) {
	l := logger.With("service", "auth")
	host := cfg.Services.ObjectStorage.Host
	port := cfg.Services.ObjectStorage.Port

	l.Info("connecting to grpc service", slog.String("host", host), slog.Int("port", port))

	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Error("error with connection to grpc service", slog.String("err", err.Error()))
		return nil, nil, err
	}

	client := objectstoragev1.NewObjectStorageClient(conn)
	return client, func() {
		conn.Close()
	}, nil
}

func InitTemplatesService(cfg *config.Config, logger *slog.Logger) (templates.TemplatesClient, func(), error) {
	l := logger.With("service", "auth")
	host := cfg.Services.TemplatesService.Host
	port := cfg.Services.TemplatesService.Port

	l.Info("connecting to grpc service", slog.String("host", host), slog.Int("port", port))

	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Error("error with connection to grpc service", slog.String("err", err.Error()))
		return nil, nil, err
	}

	client := templates.NewTemplatesClient(conn)
	return client, func() {
		conn.Close()
	}, nil
}
