package service

import (
	"context"
	"errors"

	"github.com/MaryneZa/backend-challenge/internal/core/domain"
	"github.com/MaryneZa/backend-challenge/internal/core/port"
	"github.com/MaryneZa/backend-challenge/internal/core/util"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserService struct {
	userRepo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) port.UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) Register(ctx context.Context, email string, password string) error {
	user, err := us.FindByEmail(ctx, email)


	if err != nil && err.Error() != util.ErrUserNotFound.Error(){
		return err
	}

	if user != nil {
		return errors.New("this email is already in use")
	}

	hashPassword, err := util.HashPassword(password)
	if err != nil {
		return errors.New("hash password failed:" + err.Error())
	}
	newUser := domain.User{
		Email:    email,
		Password: hashPassword,
	}
	err = us.userRepo.Create(ctx, &newUser)
	if mongo.IsDuplicateKeyError(err) {
		return errors.New("email already exists")
	}
	return nil
}

func (us *UserService) FindByID(ctx context.Context, id string) (*domain.User, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	user, err := us.userRepo.FindByID(ctx, objectID)
	if err != nil {
		return nil, errors.New("failed to check existing user:" + err.Error())
	}
	if user == nil {
		return nil, util.ErrUserNotFound
	}
	return user, nil
}

func (us *UserService) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := us.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("failed to check existing user:" + err.Error())
	}
	if user == nil {
		return nil, util.ErrUserNotFound
	}
	return user, nil
}

func (us *UserService) GetAllUser(ctx context.Context) ([]*domain.User, error) {
	return us.userRepo.GetAllUser(ctx)
}

func (us *UserService) UpdateEmail(ctx context.Context, id string, email string) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	user, err := us.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return errors.New("failed to check existing email:" + err.Error())
	}

	if user != nil {
		if user.ID != objectID {
			return errors.New("this email is already in use")
		}
		return errors.New("You’re already using this email.")
	}

	if err := us.userRepo.UpdateEmail(ctx, objectID, email); err != nil {
		return errors.New("failed to update email:" + err.Error())
	}

	return nil
}

func (us *UserService) UpdateName(ctx context.Context, id string, name string) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	user, err := us.userRepo.FindByID(ctx, objectID)
	if err != nil {
		return errors.New("failed to check existing user:" + err.Error())
	}

	if user == nil {
		return errors.New("user not found")
	}

	if err := us.userRepo.UpdateName(ctx, objectID, name); err != nil {
		return errors.New("failed to update name:" + err.Error())
	}

	return nil
}
func (us *UserService) Delete(ctx context.Context, email string) error {

	user, err := us.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return errors.New("failed to check existing user:" + err.Error())
	}

	if user == nil {
		return errors.New("user not found")
	}

	if err := us.userRepo.Delete(ctx, email); err != nil {
		return errors.New("failed to delete user:" + err.Error())
	}

	return nil
}
