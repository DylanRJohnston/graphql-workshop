package ports

import (
	"context"
	"graphql-workshop/src/models"
)

type UserCreate struct {
	Name     string
	Email    string
	Profile  string
	Birthday string
}

type UserRepository interface {
	Get(ctx context.Context, id string) (models.User, error)
	Create(ctx context.Context, fields UserCreate) (models.User, error)
}
