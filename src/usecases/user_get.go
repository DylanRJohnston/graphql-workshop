package usecases

import (
	"context"
	"graphql-workshop/src/models"
)

func (d *Dependencies) UserGet(ctx context.Context, id string) (models.User, error) {
	// Other business logic goes here
	// Are all users visible to everyone? Are there block lists, etc?
	return d.Users.Get(ctx, id)
}
