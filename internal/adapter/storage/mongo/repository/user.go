package repository

import (
	"github.com/MaryneZa/backend-challenge/internal/core/port"

)

type UserRepository struct {

}

func NewUserRepository() port.UserRepository {
	return &UserRepository{}
}