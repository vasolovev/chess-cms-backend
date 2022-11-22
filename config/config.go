package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		Log   `yaml:"logger"`
		Mongo `yaml:"mongo"`
		GRPC  `yaml:"grpc"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	Mongo struct {
		URL      string `env-required:"true" yaml:"url" env:"MONGO_URL"`
		User     string `yaml:"user" env:"MONGO_USER"`
		Password string `yaml:"password" env:"MONGO_PASSWORD"`
		Name     string `env-required:"true" yaml:"name" env:"MONGO_NAME"`
	}
	GRPC struct {
		IP   string `env-required:"true" yaml:"ip" env:"GRPC_IP"`
		Port int    `env-required:"true" yaml:"port" env:"GRPC_PORT"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	ex, er := os.Executable()
	if er != nil {
		panic(er)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	err := cleanenv.ReadConfig("/app/config/config.yml", cfg)
	if err != nil {
		err = cleanenv.ReadConfig("./config/config.yml", cfg)
		if err != nil {
			err = cleanenv.ReadConfig("../config/config.yml", cfg)
			if err != nil {
				err = cleanenv.ReadConfig("../../config/config.yml", cfg)
				if err != nil {
					return nil, fmt.Errorf("config error: %w", err)
				}
			}
		}
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
