package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/likimiad/wbtech0/internal/handlers"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type DatabaseConfig struct {
	Name     string `env:"DB_NAME"     env-required:"true"`
	User     string `env:"DB_USER"     env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
	Port     string `env:"DB_PORT"     env-required:"true"`
	Host     string `env:"DB_HOST"     env-required:"true"`
}

type NatsConfig struct {
	ClusterID string `env:"NATS_CLUSTER_ID"            env-required:"true"`
	ClientID  string `env:"NATS_CLIENT_ID"             env-required:"true"`
	URL       string `env:"NATS_SERVER_URL"            env-required:"true"`
	Ticker    int    `env:"NATS_CLIENT_TICKER_SECONDS" env-required:"true"`
}

type HTTPServer struct {
	Address string `env:"HTTP_ADDRESS"             env-default:"0.0.0.0:8080"`
}

type Config struct {
	NatsConfig
	HTTPServer
	DatabaseConfig
}

func GetConfig() *Config {
	defer func(start time.Time) {
		slog.Info("successfully initialized the config file", "duration", time.Since(start))
	}(time.Now())
	return loadConfig()
}

func loadConfig() *Config {
	exePath, err := os.Executable()
	if err != nil {
		handlers.FatalError("error_handlers getting executable path", err)
	}

	exeDir := filepath.Dir(exePath)

	configPath := filepath.Join(exeDir, ".env")
	if _, err := os.Stat(configPath); err != nil {
		handlers.FatalError("handlers opening config file", err)
	}

	var cfg Config

	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		handlers.FatalError("handlers reading config file", err)
	}

	return &cfg
}
