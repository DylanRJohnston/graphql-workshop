package usecases

import (
	"context"
	"graphql-workshop/src/models"
	"graphql-workshop/src/ports"
)

type PostCreateFields struct {
	Title   string
	Content string
}

func (d *Dependencies) PostCreate(ctx context.Context, userID string, newPost PostCreateFields) (models.Post, error) {
	// Auth or other business logic goes here
	// Validate titles, content, do we ban swear words, length limits, etc?
	return d.Posts.Create(ctx, ports.PostCreate{
		UserID:  userID,
		Title:   newPost.Title,
		Content: newPost.Content,
	})
}
