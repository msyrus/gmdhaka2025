package service

import (
	"context"

	"github.com/msyrus/gmdhaka2025/model"
)

type User struct {
	users UserRepository
}

func NewUser(users UserRepository) *User {
	return &User{
		users: users,
	}
}

func (u *User) Get(ctx context.Context, id string) (*model.User, error) {
	return u.users.FindByID(ctx, id)
}
