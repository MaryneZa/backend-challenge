package service

import (
	"context"
	"errors"

	"github.com/MaryneZa/backend-challenge/internal/core/port"
	"github.com/MaryneZa/backend-challenge/internal/core/util"

)

type AuthService struct {
	userRepo port.UserRepository
	JWTSecret string
}

func NewAuthService(userRepo port.UserRepository, JWTSecret string) port.AuthService {
	return &AuthService{userRepo: userRepo, JWTSecret: JWTSecret}
}

func (as *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := as.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("no existing account with " + email)
	}

	if ok := util.CheckPasswordHash(password, user.Password); !ok {
		return "", errors.New("invalid password")
	}

	token, err := util.CreateToken(user, as.JWTSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}