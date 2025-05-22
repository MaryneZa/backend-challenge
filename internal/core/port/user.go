package port

import (
	"context"
	"github.com/MaryneZa/backend-challenge/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
}

type UserService interface {

}