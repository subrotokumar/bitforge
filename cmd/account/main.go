package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/subrotokumar/bitforge/cmd/account/config"
	"github.com/subrotokumar/bitforge/pkg/logger"
)

func main() {
	logger, _ := logger.NewLogger()
	defer logger.Sync()

	log := logger.Sugar()

	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load config : %s", err.Error())
	}
	print(config.App.Name)
}
