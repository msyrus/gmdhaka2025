package repository

import (
	"context"
	"errors"

	"github.com/msyrus/gmdhaka2025/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	collection *mongo.Collection
}

func NewProduct(db *mongo.Database) *Product {
	return &Product{collection: db.Collection("products")}
}

func (p *Product) FindByID(ctx context.Context, id string) (*model.Product, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidID
	}

	var t model.Product
	if err := p.collection.FindOne(ctx, bson.D{{Key: "_id", Value: oID}}).Decode(&t); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (p *Product) Add(ctx context.Context, product *model.Product) error {
	res, err := p.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	product.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (p *Product) Update(ctx context.Context, product *model.Product) error {
	// implementation
	panic("not implemented")
}

func (p *Product) Delete(ctx context.Context, id string) error {
	// implementation
	panic("not implemented")
}

func (p *Product) FindAll(ctx context.Context) ([]*model.Product, error) {
	// implementation
	panic("not implemented")
}
