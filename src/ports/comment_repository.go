package ports

import (
	"context"
	"graphql-workshop/src/models"
)

type CommentCreate struct {
	UserID  string
	PostID  string
	Content string
}

type CommentRepository interface {
	Get(ctx context.Context, id string) (models.Comment, error)
	ForPost(ctx context.Context, postID string) ([]models.Comment, error)
	Create(ctx context.Context, fields CommentCreate) (models.Comment, error)
}
