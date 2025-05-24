package port

import (
	"context"
	"github.com/MaryneZa/backend-challenge/internal/core/domain"
)

type TokenService interface {
	CreateToken(ctx context.Context, user domain.User) (string, error)
	VerifyToken(ctx context.Context, token string) error
}

type AuthService interface {
	Login(ctx context.Context, email string, password string) (string, error)
}