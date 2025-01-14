package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DBModel interface {
	SetID(id string)
}

type mongoRepo[M DBModel] struct {
	collection *mongo.Collection
}

func NewMongoRepo[M DBModel](db *mongo.Database, collectionName string) *mongoRepo[M] {
	return &mongoRepo[M]{collection: db.Collection(collectionName)}
}

func (m *mongoRepo[M]) FindByID(ctx context.Context, id string) (M, error) {
	var t M
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return t, ErrInvalidID
	}

	if err := m.collection.FindOne(ctx, bson.D{{Key: "_id", Value: oID}}).Decode(&t); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return t, nil
		}
		return t, err
	}
	return t, nil
}

func (m *mongoRepo[M]) Add(ctx context.Context, model M) error {
	res, err := m.collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}
	model.SetID(res.InsertedID.(primitive.ObjectID).Hex())
	return nil
}

func (m *mongoRepo[M]) Update(ctx context.Context, model M) error {
	// implementation
	panic("not implemented")
}

func (m *mongoRepo[M]) Delete(ctx context.Context, id string) error {
	// implementation
	panic("not implemented")
}
