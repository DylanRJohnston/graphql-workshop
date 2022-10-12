package usecases

import "graphql-workshop/src/ports"

// As an app grows a single dependency struct will become very wide, you should should split it into smaller sections
type Dependencies struct {
	Users    ports.UserRepository
	Comments ports.CommentRepository
	Posts    ports.PostRepository
}
