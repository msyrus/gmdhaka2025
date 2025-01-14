package repository

import (
	"context"
	"errors"

	"github.com/msyrus/gmdhaka2025/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	collection *mongo.Collection
}

func NewUser(db *mongo.Database) *User {
	return &User{collection: db.Collection("users")}
}

func (u *User) FindByID(ctx context.Context, id string) (*model.User, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidID
	}

	var t model.User
	if err := u.collection.FindOne(ctx, bson.D{{Key: "_id", Value: oID}}).Decode(&t); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (u *User) Add(ctx context.Context, user *model.User) error {
	res, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (u *User) Update(ctx context.Context, user *model.User) error {
	// implementation
	panic("not implemented")
}

func (u *User) Delete(ctx context.Context, id string) error {
	// implementation
	panic("not implemented")
}

func (u *User) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	// implementation
	panic("not implemented")
}
