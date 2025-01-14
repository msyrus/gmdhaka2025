package repository

import (
	"context"

	"github.com/msyrus/gmdhaka2025/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserGeneric struct {
	mongoRepo[*model.User]
}

func NewUserGeneric(db *mongo.Database) *UserGeneric {
	return &UserGeneric{
		mongoRepo[*model.User]{collection: db.Collection("users")},
	}
}

func (u *UserGeneric) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	// implementation
	panic("not implemented")
}
