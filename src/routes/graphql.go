package routes

import (
	"graphql-workshop/src/graph"
	"graphql-workshop/src/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func (d *Dependences) GraphQLHandler() gin.HandlerFunc {

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers:  &graph.Resolver{},
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (d *Dependences) PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
