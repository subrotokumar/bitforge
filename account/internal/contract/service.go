package contract

import (
	"context"

	gen "github.com/subrotokumar/bitforge-pkg/pb"
)

type AccountService interface {
	HealthCheck(ctx context.Context, _ gen.Empty) (*gen.HealthCheckResponse, error)
	CreateAccount(ctx context.Context, request *gen.CreateAccountRequest) (*gen.CreateAccountRequest, error)
	Magiclink(context.Context, *gen.MagiclinkRequest) (*gen.MagiclinkResponse, error)
	Authenticate(context.Context, *gen.AuthenticateRequest) (*gen.AuthenticateResponse, error)
}
