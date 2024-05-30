package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App struct {
		Host string `env:"APP_HOST" env-default:"0.0.0.0"`
		Port int    `env:"APP_PORT" env-default:"8080"`
	}
	Services struct {
		AuthService struct {
			Host string `env:"AUTH_SERVICE_HOST" env-default:"localhost"`
			Port int    `env:"AUTH_SERVICE_PORT" env-default:"50053"`
		}
	}

	LogLevel string `env:"LOG_LEVEL" env-default:"debug"`
}

func New() *Config {
	config := &Config{}
	if err := cleanenv.ReadEnv(config); err != nil {
		header := "SOCHYA GATEWAY"
		f := cleanenv.FUsage(os.Stdout, config, &header)
		f()
		panic(err)
	}

	return config
}