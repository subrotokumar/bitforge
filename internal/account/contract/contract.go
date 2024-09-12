package contract

import (
	"context"

	"github.com/gofrs/uuid"
)

type (
	AccountRepository interface {
		CreateAccount(ctx context.Context, email string) error
		GetUser(ctx context.Context, email string) error
		GetUserById(ctx context.Context, id uuid.UUID)
	}

	AccountService interface {
		CreateAccount(ctx context.Context, email string) error
	}
)
