package usecases

import (
	"context"
	"graphql-workshop/src/models"
)

func (d *Dependencies) PostGet(ctx context.Context, id string) (models.Post, error) {
	// Auth or other business logic goes here
	// Are all posts visible to everyone? Are there block lists, etc?

	return d.Posts.Get(ctx, id)
}
