package main

import (
	"fmt"
	"net"

	_ "github.com/joho/godotenv/autoload"
	"github.com/subrotokumar/bitforge-account/cmd/config"
	"github.com/subrotokumar/bitforge-account/internal/repository"
	"github.com/subrotokumar/bitforge-account/internal/service"
	"github.com/subrotokumar/bitforge-pkg/logger"
	gen "github.com/subrotokumar/bitforge-pkg/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger, _ := logger.NewLogger()
	defer logger.Sync()

	log := logger.Sugar()

	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load config : %s", err.Error())
	}

	grpcServer := grpc.NewServer()

	listen := fmt.Sprintf("%s:%d", config.HTTP.Host, config.HTTP.Port)
	lis, err := net.Listen("tcp", listen)
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	repo := repository.New()
	svc := service.New(repo)

	reflection.Register(grpcServer)
	gen.RegisterAccountServiceServer(grpcServer, svc)

	log.Infoln(config.HTTP)
	log.Infof("Account server started at %s", listen)
	grpcServer.Serve(lis)
}
