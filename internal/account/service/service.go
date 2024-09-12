package service

import (
	"context"

	"github.com/subrotokumar/bitforge/internal/account/contract"
)

type service struct {
	repo contract.AccountRepository
}

func New(repo contract.AccountRepository) *service {
	return &service{repo: repo}
}

func (s *service) CreateAccount(ctx context.Context, email string) error {
	return s.repo.CreateAccount(ctx, email)
}
