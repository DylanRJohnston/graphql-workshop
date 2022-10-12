package usecases

import (
	"context"
	"graphql-workshop/src/models"
	"graphql-workshop/src/ports"
)

type CommentCreateFields struct {
	UserID  string
	Content string
}

func (d *Dependencies) CommentCreate(ctx context.Context, postID string, comment CommentCreateFields) (models.Comment, error) {
	// Auth or other business logic goes here
	// Validate the comment content, should we ban swear words, is there a block list, maximum length, etc
	return d.Comments.Create(ctx, ports.CommentCreate{
		UserID:  comment.UserID,
		PostID:  postID,
		Content: comment.Content,
	})
}
