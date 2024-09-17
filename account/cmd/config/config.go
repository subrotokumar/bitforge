package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/subrotokumar/bitforge-pkg/config"
)

type (
	Config struct {
		config.App  `yaml:"app"`
		config.HTTP `yaml:"http"`
		config.Log  `yaml:"logger"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("config path: " + dir)

	err = cleanenv.ReadConfig(dir+"/cmd/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
