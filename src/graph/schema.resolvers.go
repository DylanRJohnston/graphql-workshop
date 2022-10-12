package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql-workshop/src/graph/generated"
	"graphql-workshop/src/graph/gql"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*gql.User, error) {
	user, err := r.Users.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &gql.User{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Profile:  &user.Profile,
		Birthday: user.Birthday,
		Posts:    nil,
		Comments: nil,
	}, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*gql.Post, error) {
	post, err := r.Posts.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &gql.Post{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		Comments: nil,
		User:     nil,
	}, nil
}

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context, id string) (*gql.Comment, error) {
	comment, err := r.Comments.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &gql.Comment{
		ID:      comment.ID,
		Content: comment.Content,
		User:    nil,
		Post:    nil,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
