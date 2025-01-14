package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func findByID(ctx context.Context, collection *mongo.Collection, id string, result interface{}) error {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID
	}

	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: oID}}).Decode(result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil
		}
		return err
	}
	return nil
}

func add(ctx context.Context, collection *mongo.Collection, model DBModel) error {
	res, err := collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}
	model.SetID(res.InsertedID.(primitive.ObjectID).Hex())
	return nil
}

func addV2[M DBModel](ctx context.Context, collection *mongo.Collection, model M) error {
	res, err := collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}
	model.SetID(res.InsertedID.(primitive.ObjectID).Hex())
	return nil
}
