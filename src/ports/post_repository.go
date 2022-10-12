package ports

import (
	"context"
	"graphql-workshop/src/models"
)

type PostCreate struct {
	UserID  string
	Title   string
	Content string
}

type PostRepository interface {
	Get(ctx context.Context, id string) (models.Post, error)
	ForUser(ctx context.Context, userID string) ([]models.Post, error)
	Create(ctx context.Context, fields PostCreate) (models.Post, error)
}
