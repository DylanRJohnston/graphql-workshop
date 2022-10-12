package routes

import (
	"graphql-workshop/src/usecases"

	"github.com/gin-gonic/gin"
)

type Dependences struct {
	usecase usecases.Dependencies
}

// No business logic should exist in these gin controllers
// only packing and unpacking of data to and from the http layer

func New(deps usecases.Dependencies) *gin.Engine {
	d := Dependences{deps}

	r := gin.Default()

	r.POST("/query", d.GraphQLHandler())
	r.GET("/", d.PlaygroundHandler())

	r.GET("/users/:userID", d.UserGet)
	r.GET("/users/:userID/posts", d.UserGetPosts)
	r.GET("/posts/:postID", d.PostGet)
	r.GET("/posts/:postID/comments", d.PostGetComments)
	r.GET("/comments/:commentID", d.CommentGet)

	r.POST("/users", d.UserCreate)
	r.POST("/users/:userID/posts", d.PostCreate)
	r.POST("/posts/:postID/comments", d.CommentCreate)

	return r
}
