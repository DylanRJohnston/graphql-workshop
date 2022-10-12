package usecases

import (
	"context"
	"graphql-workshop/src/models"
	"graphql-workshop/src/ports"
)

type UserCreateFields struct {
	Name     string
	Email    string
	Profile  string
	Birthday string
}

func (d *Dependencies) UserCreate(ctx context.Context, user UserCreateFields) (models.User, error) {
	// Other business logic goes here
	// Validate emails, birthdays, etc here
	return d.Users.Create(ctx, ports.UserCreate{
		Name:     user.Name,
		Email:    user.Email,
		Profile:  user.Profile,
		Birthday: user.Birthday,
	})
}
