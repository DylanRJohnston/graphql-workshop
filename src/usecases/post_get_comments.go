package usecases

import (
	"context"
	"graphql-workshop/src/models"
)

func (d *Dependencies) PostGetComments(ctx context.Context, postID string) ([]models.Comment, error) {
	// Other business logic goes here
	// Are all comments visible to everyone? Are there block lists, etc?
	return d.Comments.ForPost(ctx, postID)
}
