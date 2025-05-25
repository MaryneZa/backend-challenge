package service

import (
	"context"
	"testing"

	"github.com/MaryneZa/backend-challenge/internal/core/domain"
	"github.com/MaryneZa/backend-challenge/internal/core/util/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	email := "test@example.com"
	password := "password"

	// success
	t.Run("Register successful", func(t *testing.T) {
		userRepo := test.NewMockUserRepository(test.CreateFuncSuccess, test.FindByEmailFuncNotFound)
		userService := NewUserService(userRepo)
		err := userService.Register(context.TODO(), email, password)
		assert.NoError(t, err)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	// not success
	t.Run("Register unsuccessful", func(t *testing.T) {
		userRepo := test.NewMockUserRepository(test.CreateFuncFailed, test.FindByEmailFuncFound)
		userService := NewUserService(userRepo)
		err := userService.Register(context.TODO(), email, password)
		assert.Error(t, err)
		if err == nil {
			t.Errorf("expected error for invalid ObjectID, got nil")
		} else {
			t.Logf("expected failure occurred: %v", err)
		}
	})
}

func TestFindUser(t *testing.T) {

	type Input struct {
		ID    string
		Email string
	}

	// BY ID

	// success
	t.Run("find by id successful", func(t *testing.T) {
		userRepo := test.NewMockUserRepository(nil, nil)
		userService := NewUserService(userRepo)
		_, err := userService.FindByID(context.TODO(), "6831b53acf66afd6d203efbe")
		assert.NoError(t, err)

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

	})

	// not success
	t.Run("find by id unsuccessful", func(t *testing.T) {
		userRepo := test.NewMockUserRepository(nil, nil)
		userService := NewUserService(userRepo)
		input := []Input{
			{ID: "6831b53a"},
			{ID: "test test"},
		}

		for _, tc := range input {
			_, err := userService.FindByID(context.TODO(), tc.ID)
			assert.Error(t, err)

			if err == nil {
				t.Errorf("expected error for invalid ObjectID, got nil")
			} else {
				t.Logf("expected failure occurred: %v", err)
			}
		}

	})

	// BY EMAIL
	t.Run("find by email successful", func(t *testing.T) {
		userRepo := test.NewMockUserRepository(nil, test.FindByEmailFuncFound)
		userService := NewUserService(userRepo)
		_, err := userService.FindByEmail(context.TODO(), "test@example.com")
		assert.NoError(t, err)

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("find by email unsuccessful", func(t *testing.T) {
		userRepo := test.NewMockUserRepository(nil, test.FindByEmailFuncNotFound)
		userService := NewUserService(userRepo)
		_, err := userService.FindByEmail(context.TODO(), "test@example.com")
		assert.Error(t, err)

		if err == nil {
			t.Errorf("expected error for invalid ObjectID, got nil")
		} else {
			t.Logf("expected failure occurred: %v", err)
		}

	})
}

func TestUpdateUser(t *testing.T) {

	// EMAIL

	// success
	t.Run("update user email successful", func(t *testing.T) {
		userRepo := test.NewMockUserRepository(nil,test. FindByEmailFuncNotFound)
		userService := NewUserService(userRepo)
		err := userService.UpdateEmail(context.TODO(), "6831b53acf66afd6d203efbe", "test@example.com")
		assert.NoError(t, err)

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	// not success (duplicate)
	t.Run("update user email unsuccessful", func(t *testing.T) {
		userRepo := test.NewMockUserRepository(nil, test.FindByEmailFuncFound)
		userService := NewUserService(userRepo)
		err := userService.UpdateEmail(context.TODO(), "6831b53acf66afd6d203efbe", "test@example.com")
		assert.Error(t, err)

		if err == nil {
			t.Errorf("expected error for invalid ObjectID, got nil")
		} else {
			t.Logf("expected failure occurred: %v", err)
		}

	})

	// Name

	//success
	t.Run("update user name successful", func(t *testing.T) {
		userRepo := test.NewMockUserRepository(nil, nil)
		userService := NewUserService(userRepo)
		err := userService.UpdateName(context.TODO(), "6831b53acf66afd6d203efbe", "testtest")
		assert.NoError(t, err)

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

}

func TestGetUsers(t *testing.T) {

	t.Run("get users successful", func(t *testing.T) {
		userRepo := test.NewMockUserRepository(nil, nil)
		userService := NewUserService(userRepo)
		users, err := userService.GetAllUser(context.TODO())
		assert.NoError(t, err)
		assert.Equal(t, []*domain.User{}, users)

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

	})
}

func TestDeleteUser(t *testing.T) {
	// success
	t.Run("delete user successful", func(t *testing.T) {
		userRepo := test.NewMockUserRepository(nil, test.FindByEmailFuncFound)
		userService := NewUserService(userRepo)
		err := userService.Delete(context.TODO(), "test@example.com")
		assert.NoError(t, err)

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	// not success (email not found)
	t.Run("delete user unsuccessful", func(t *testing.T) {
		userRepo := test.NewMockUserRepository(nil,test. FindByEmailFuncNotFound)
		userService := NewUserService(userRepo)
		err := userService.Delete(context.TODO(), "test@example.com")
		assert.Error(t, err)

		if err == nil {
			t.Errorf("expected error for invalid ObjectID, got nil")
		} else {
			t.Logf("expected failure occurred: %v", err)
		}

	})

}
