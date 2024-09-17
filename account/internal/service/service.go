package service

import (
	"context"
	"fmt"
	"time"

	"github.com/subrotokumar/bitforge-account/internal/contract"
	gen "github.com/subrotokumar/bitforge-pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type service struct {
	gen.UnimplementedAccountServiceServer
	startedAt time.Time
	repo      contract.AccountRepository
}

func New(repo contract.AccountRepository) *service {
	return &service{repo: repo, startedAt: time.Now()}
}

func (s *service) HealthCheck(ctx context.Context, empty *gen.Empty) (*gen.HealthCheckResponse, error) {
	return &gen.HealthCheckResponse{
		Service: "account",
		Status:  gen.HealthStatus_healthy,
		Component: &gen.HealthCheckComponent{
			Database:    gen.HealthStatus_unknown,
			Cache:       gen.HealthStatus_unknown,
			ExternalApi: gen.HealthStatus_unknown,
		},
		Uptime:    time.Since(s.startedAt).String(),
		Timestamp: timestamppb.Now(),
	}, nil
}

func (s *service) CreateAccount(ctx context.Context, request *gen.CreateAccountRequest) (*gen.CreateAccountRequest, error) {
	return &gen.CreateAccountRequest{
		Email: "subrotokumar@email.com",
	}, nil
}

func (s *service) Magiclink(ctx context.Context, request *gen.MagiclinkRequest) (*gen.MagiclinkResponse, error) {
	return &gen.MagiclinkResponse{
		Message: fmt.Sprintf("Magiclink is send to %s", request.GetEmail()),
	}, nil
}

func (s *service) Authenticate(context.Context, *gen.AuthenticateRequest) (*gen.AuthenticateResponse, error) {
	return &gen.AuthenticateResponse{
		AccessToken:          "access token demo",
		RefreshToken:         "refresh token gemo",
		AccessTokenExpireAt:  time.Now().Add(time.Hour * 24).Unix(),
		RefreshTokenExpireAt: time.Now().Add(time.Hour * 30).Unix(),
	}, nil
}
