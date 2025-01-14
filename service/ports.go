package service

import (
	"context"

	"github.com/msyrus/gmdhaka2025/model"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Add(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type ProductRepository interface {
	FindByID(ctx context.Context, id string) (*model.Product, error)
	FindAll(ctx context.Context) ([]*model.Product, error)
	Add(ctx context.Context, product *model.Product) error
	Update(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, id string) error
}
