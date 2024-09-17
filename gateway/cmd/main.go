package main

import (
	"github.com/subrotokumar/bitforge-gateway/cmd/config"
	"github.com/subrotokumar/bitforge-pkg/logger"
)

func main() {
	logger, _ := logger.NewLogger()
	defer logger.Sync()

	log := logger.Sugar()

	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load config : %s", err.Error())
	}

	log.Info(config)
}
