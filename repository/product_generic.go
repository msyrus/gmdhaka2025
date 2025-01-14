package repository

import (
	"context"

	"github.com/msyrus/gmdhaka2025/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductGeneric struct {
	mongoRepo[*model.Product]
}

func NewProductGeneric(db *mongo.Database) *ProductGeneric {
	return &ProductGeneric{
		mongoRepo[*model.Product]{collection: db.Collection("products")},
	}
}

func (u *ProductGeneric) FindAll(ctx context.Context) ([]*model.Product, error) {
	// implementation
	panic("not implemented")
}
