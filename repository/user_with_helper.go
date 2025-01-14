package repository

import (
	"context"

	"github.com/msyrus/gmdhaka2025/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserWithHelper struct {
	collection *mongo.Collection
}

func NewUserWithHelper(db *mongo.Database) *User {
	return &User{collection: db.Collection("users")}
}

func (u *UserWithHelper) FindByID(ctx context.Context, id string) (*model.User, error) {
	var t *model.User
	return t, findByID(ctx, u.collection, id, &t)
}

func (u *UserWithHelper) Add(ctx context.Context, user *model.User) error {
	return add(ctx, u.collection, user)
}

func (u *UserWithHelper) Update(ctx context.Context, user *model.User) error {
	// implementation
	panic("not implemented")
}

func (u *UserWithHelper) Delete(ctx context.Context, id string) error {
	// implementation
	panic("not implemented")
}

func (u *UserWithHelper) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	// implementation
	panic("not implemented")
}
