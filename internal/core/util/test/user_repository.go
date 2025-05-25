package test

import (
	"errors"
	"context"
	"github.com/MaryneZa/backend-challenge/internal/core/port"
	"github.com/MaryneZa/backend-challenge/internal/core/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateFuncSuccess(ctx context.Context, user *domain.User) error {
	return nil
}

func CreateFuncFailed(ctx context.Context, user *domain.User) error {
	return errors.New("email already exists")
}

func FindByEmailFuncFound(ctx context.Context, email string) (*domain.User, error) {
	return &domain.User{Email: email}, nil
}

func FindByEmailFuncNotFound(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}

type MockUserRepository struct {
	CreateFunc      func(ctx context.Context, user *domain.User) error
	FindByEmailFunc func(ctx context.Context, email string) (*domain.User, error)
}

func NewMockUserRepository(
	CreateFunc      func(ctx context.Context, user *domain.User) error ,
	FindByEmailFunc func(ctx context.Context, email string) (*domain.User, error)) port.UserRepository {
	return &MockUserRepository{FindByEmailFunc: FindByEmailFunc, CreateFunc: CreateFunc}
}

func (mu *MockUserRepository) Create(ctx context.Context, user *domain.User) error {
	if mu.CreateFunc != nil {
		return mu.CreateFunc(ctx, user)
	}
	return nil

}

func (mu *MockUserRepository) FindByID(ctx context.Context, id bson.ObjectID) (*domain.User, error){
	return &domain.User{Name: "Test ID"}, nil
}

func (mu *MockUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error){
	if mu.FindByEmailFunc != nil {
		return mu.FindByEmailFunc(ctx, email)
	}
	return &domain.User{Email: email}, nil
}

func (mu *MockUserRepository) GetAllUser(ctx context.Context) ([]*domain.User, error) {
	return []*domain.User{}, nil
}

func (mu *MockUserRepository) UpdateEmail(ctx context.Context, id bson.ObjectID, email string) error {
	return nil
}

func (mu *MockUserRepository) UpdateName(ctx context.Context, id bson.ObjectID, name string) error {
	return nil
}

func (mu *MockUserRepository) Delete(ctx context.Context, email string) error {
	return nil
}

func (mu *MockUserRepository) SetUpIndexes(ctx context.Context) error {
	return nil
}