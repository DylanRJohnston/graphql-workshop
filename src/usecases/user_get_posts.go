package usecases

import (
	"context"
	"graphql-workshop/src/models"
)

func (d *Dependencies) UserGetPosts(ctx context.Context, userID string) ([]models.Post, error) {
	// Other business logic goes here
	// Are all posts visible to everyone? Are there block lists, etc?
	return d.Posts.ForUser(ctx, userID)
}
