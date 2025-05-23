package service

import (
	"context"

	"github.com/MaryneZa/backend-challenge/internal/core/domain"
	"github.com/MaryneZa/backend-challenge/internal/core/port"
)

type UserService struct {
	userRepo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) port.UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) Register(ctx context.Context, user *domain.User) error {
	
	return
}
