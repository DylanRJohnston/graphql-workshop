package usecases

import (
	"context"
	"graphql-workshop/src/models"
)

func (d *Dependencies) CommentGet(ctx context.Context, id string) (models.Comment, error) {
	// Auth or other business logic goes here
	// Are all comments visible to everyone? Are there block lists, etc?
	return d.Comments.Get(ctx, id)
}
