package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-workshop/src/graph/generated"
	"graphql-workshop/src/graph/gql"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// User is the resolver for the user field.
func (r *commentResolver) User(ctx context.Context, obj *gql.Comment) (*gql.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Post is the resolver for the post field.
func (r *commentResolver) Post(ctx context.Context, obj *gql.Comment) (*gql.Post, error) {
	panic(fmt.Errorf("not implemented: Post - post"))
}

// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *gql.Post) ([]*gql.Comment, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// User is the resolver for the user field.
func (r *postResolver) User(ctx context.Context, obj *gql.Post) (*gql.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*gql.User, error) {
	user, err := r.Deps.Users.Get(ctx, id)
	if err != nil {
		return nil, gqlerror.Errorf("Error encountered while fetching User %s", err.Error())
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
	post, err := r.Deps.Posts.Get(ctx, id)
	if err != nil {
		return nil, gqlerror.Errorf("Error encountered while fetching Posts %s", err.Error())
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
	comment, err := r.Deps.Comments.Get(ctx, id)
	if err != nil {
		return nil, gqlerror.Errorf("Error encountered while fetching comment %s", err.Error())
	}

	return &gql.Comment{
		ID:      comment.ID,
		Content: comment.Content,
		User:    nil,
		Post:    nil,
	}, nil
}

// Posts is the resolver for the posts field.
func (r *userResolver) Posts(ctx context.Context, obj *gql.User) ([]*gql.Post, error) {
	posts, err := r.Deps.Posts.ForUser(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	gqlPosts := []*gql.Post{}
	for _, post := range posts {
		gqlPosts = append(gqlPosts, &gql.Post{
			ID:       post.ID,
			Title:    post.Title,
			Content:  post.Content,
			Comments: nil,
			User:     obj,
		})
	}

	return gqlPosts, nil
}

// Comments is the resolver for the comments field.
func (r *userResolver) Comments(ctx context.Context, obj *gql.User) ([]*gql.Comment, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type commentResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
