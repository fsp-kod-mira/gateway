package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App struct {
		Host string `env:"APP_HOST" env-default:"0.0.0.0"`
		Port int    `env:"APP_PORT" env-default:"8081"`
	}
	Services struct {
		AuthService struct {
			Host string `env:"AUTH_SERVICE_HOST" env-default:"localhost"`
			Port int    `env:"AUTH_SERVICE_PORT" env-default:"50053"`
		}
		CvService struct {
			Host string `env:"CV_SERVICE_HOST" env-default:"localhost"`
			Port int    `env:"CV_SERVICE_PORT" env-default:"50054"`
		}
		ObjectStorage struct {
			Host string `env:"OBJECT_STORAGE_SERVICE_HOST" env-default:"localhost"`
			Port int    `env:"OBJECT_STORAGE_SERVICE_PORT" env-default:"50054"`
		}
		TemplatesService struct {
			Host string `env:"TEMPLATES_SERVICE_HOST" env-default:"localhost"`
			Port int    `env:"TEMPLATES_SERVICE_PORT" env-default:"50054"`
		}
	}
	CORS struct {
		Origin string `env:"CORS_ORIGIN" env-default:"http://localhost:3001"`
	}

	LogLevel string `env:"LOG_LEVEL" env-default:"debug"`
}

func New() *Config {
	config := &Config{}
	if err := cleanenv.ReadEnv(config); err != nil {
		header := "GATEWAY"
		f := cleanenv.FUsage(os.Stdout, config, &header)
		f()
		panic(err)
	}

	return config
}
