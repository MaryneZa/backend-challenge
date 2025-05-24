package port

import (
	"context"
	"github.com/MaryneZa/backend-challenge/internal/core/domain"
	"go.mongodb.org/mongo-driver/v2/bson"

)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id bson.ObjectID) (*domain.User, error) 
	FindByEmail(ctx context.Context, email string) (*domain.User, error) 
	GetAllUser(ctx context.Context) ([]*domain.User, error) 
	UpdateEmail(ctx context.Context, id bson.ObjectID, email string) error 
	UpdateName(ctx context.Context, id bson.ObjectID, name string) error 
	Delete(ctx context.Context, email string) error

	SetUpIndexes(ctx context.Context) error
}

type UserService interface {
	Register(ctx context.Context, email string, password string) error
	FindByID(ctx context.Context, id string) (*domain.User, error) 
	FindByEmail(ctx context.Context, email string) (*domain.User, error) 
	GetAllUser(ctx context.Context) ([]*domain.User, error) 
	UpdateEmail(ctx context.Context, id string, email string) error 
	UpdateName(ctx context.Context, id string, name string) error 
	Delete(ctx context.Context, email string) error 
}