package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		App     `yaml:"app"`
		HTTP    `yaml:"http"`
		Log     `yaml:"logger"`
		MySQL   `yaml:"mysql"`
		Goulash `yaml:"goulash"`
	}

	// App -.
	App struct {
		IsGoulash bool `env-required:"true" yaml:"is_goulash" env:"APP_IS_GOULASH" env-default:"true"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT" env-default:"8083"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
		Out   string `                    yaml:"log_out"   env:"LOG_OUT"`
	}

	MySQL struct {
		Username        string        `yaml:"username"          env:"DB_USERNAME"`
		Password        string        `yaml:"password"          env:"DB_PASSWORD"`
		Host            string        `yaml:"host"              env:"DB_HOST"`
		Port            string        `yaml:"port"              env:"DB_PORT"`
		DBName          string        `yaml:"db_name"           env:"DB_NAME"`
		Timeout         time.Duration `yaml:"timeout"           env:"DB_TIMEOUT"          env-default:"3s"`
		RefreshInterval time.Duration ` yaml:"refresh_interval"  env:"DB_REFRESH_INTERVAL" env-default:"5m"`
	}

	Goulash struct {
		APIUrl  string `yaml:"api_url"     env:"GOULASH_API_URL"`
		APIKey  string `yaml:"api_key"     env:"GOULASH_API_KEY"`
		APIUuid string `yaml:"api_uuid"     env:"GOULASH_API_UUID"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	// Загрузка переменных окружения
	_ = godotenv.Load(".env")

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
