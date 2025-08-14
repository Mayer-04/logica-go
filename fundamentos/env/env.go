package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload" // Carga automáticamente el .env
)

type LogConfig struct {
	Style string
	Level string
}

type PostgresConfig struct {
	Username string
	Password string
	URL      string
	Port     string
}

type Config struct {
	Logs LogConfig
	DB   PostgresConfig
	Port string
}

func main() {

}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port: os.Getenv("PORT"),
		Logs: LogConfig{
			Style: os.Getenv("LOG_STYLE"),
			Level: os.Getenv("LOG_LEVEL"),
		},
		DB: PostgresConfig{
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PWD"),
			URL:      os.Getenv("POSTGRES_URL"),
			Port:     os.Getenv("POSTGRES_PORT"),
		},
	}

	// Validaciones
	if cfg.Port == "" {
		return nil, fmt.Errorf("PORT is required")
	}
	if cfg.DB.Username == "" || cfg.DB.Password == "" {
		return nil, fmt.Errorf("POSTGRES_USER and POSTGRES_PWD are required")
	}
	if cfg.Logs.Level != "debug" && cfg.Logs.Level != "info" && cfg.Logs.Level != "error" {
		return nil, fmt.Errorf("LOG_LEVEL must be one of: debug, info, error")
	}

	return cfg, nil
}
