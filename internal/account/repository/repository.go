package repository

import (
	"context"

	"github.com/gofrs/uuid"
)

type repository struct {
}

func New() *repository {
	return &repository{}
}

func (r *repository) CreateAccount(ctx context.Context, email string) error {
	return nil
}
func (r *repository) GetUser(ctx context.Context, email string) error {
	return nil
}
func (r *repository) GetUserById(ctx context.Context, id uuid.UUID) {
	return
}
