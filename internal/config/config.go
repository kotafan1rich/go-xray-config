package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	XRayConfigPath string `env:"X_RAY_CONFIG_PATH" envDefault:"/usr/local/etc/xray/config.json"`
	XRayHost       string `env:"X_RAY_HOST_IP"`
	XrayPublicKey  string `env:"X_RAY_PUBLIC_KEY"`
	XRayServerName string `env:"SERVER_NAME"`

	ServerPort string `env:"SERVER_PORT" envDefault:"8080"`
	GinMode    string `env:"GIN_MODE" envDefault:"debug"`

	Token string `env:"TOKEN" required:"true"`

	AppName string `env:"APP_NAME" envDefault:"go-xray-config"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
