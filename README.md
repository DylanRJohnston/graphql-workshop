# Workshop

## Converting to GraphQL

- Create a `tools.go` file to hold the import for gqlgen. This pattern for managing go dev tools is recommend and documented [here](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module)
- Add the following to the `tools.go` file.

  ```go
  package tools

  import (
    _ "github.com/99designs/gqlgen"
    _ "github.com/99designs/gqlgen/graphql/introspection"
  )
  ```

- Download the package `go get "github.com/99designs/gqlgen"`
- Run the initial code generation `go run github.com/99designs/gqlgen init`

  - This will create a `./graph` folder, a `gqlgen.yml`, and `server.go` file.

- Fix up the generated file locations.

  - Move the `graph` folder inside the `src` folder
  - Delete the `server.go` file
  - Edit `gqlgen.yml` for the new generated file locations, mainly prefixing things with src, we'll also change the generated models package to `gql` to distinguish are database model types from our graphql api types

    ```yml
    # Where are all the schema files located? globs are supported eg  src/**/*.graphqls
    schema:
      - src/graph/*.graphqls

    # Where should the generated server code go?
    exec:
      filename: src/graph/generated/generated.go
      package: generated

    # Uncomment to enable federation
    # federation:
    #   filename: graph/generated/federation.go
    #   package: generated

    # Where should any generated models go?
    model:
      filename: src/graph/gql/models_gen.go
      package: gql

    # Where should the resolver implementations go?
    resolver:
      layout: follow-schema
      dir: src/graph
      package: graph
    ```

- Create a new file in `src/routes`, `graphql.go` to bridge the gap between Graphql and Gin.

  ```go
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
  ```

- Connect the handlers in `src/routes/routes.go`

  ```go
    r.POST("/query", d.GraphQLHandler())
    r.GET("/", d.PlaygroundHandler())
  ```

- You should be able to launch the server and see the GraphQL playground. `go run src/entrypoints/server/main.go`

- You'll notice that the schema is the default todos schema, so now we'll edit the schema in `src/graph/schema.graphqls`. We'll start by just describing the data of each model without their connections.

  ```gql
  type User {
    id: String!
    name: String!
    email: String!
    profile: String
    birthday: String!
  }

  type Post {
    id: String!
    title: String!
    content: String!
  }

  type Comment {
    id: String!
    content: String!
  }
  ```

- Now comes the real power of **Graph**QL. We can connect our entities together.
  - Users have posts and comments so we'll add that to the users
    ```gql
    posts: [Post!]!
    comments: [Comment!]!
    ```
  - Similarly posts have comments, so we'll add that to the post, and they belong to a user so we'll add that back connection.
    ```gql
    comments: [Comment!]!
    user: User!
    ```
  - Comments are made by a user, so we'll connect that too, and they belong to a post, so we'll add the back connection.
    ```gql
    user: User!
    post: Post!
    ```
- But how do we actually query anything? GraphQL has a special root type `Query` that we can attached fields to that will server as our entrypoint into our data graph.
  ```gql
  type Query {
    user(id: ID!): User
    post(id: ID!): Post
    comment(id: ID!): Comment
  }
  ```
- We'll regenerate our resolvers `go run github.com/99designs/gqlgen generate`
- Now if we rerun our server we should have our fancy new schema. `go run src/entrypoints/server/main.go`.
- Trying to execute a query will result in an error because our resolvers are unimplemented

  ```gql
  {
    user(id: "") {
      posts {
        title
        content
        comments {
          content
          user {
            name
            profile
          }
        }
      }
    }
  }
  ```

- In `src/graph/resolver.go` add our `usecases.Dependencies` to the `Resolver` struct so we have access to our use case handlers in our resolvers.
- Start filling out the implementation of our resolvers in `src/graph/schema.resolvers.go`
- Make sure to connect the dependency injection in `src/routes/graphql.go`
  ```go
  h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
    Resolvers: &graph.Resolver{
      Dependencies: d.usecase,
    },
    Directives: generated.DirectiveRoot{},
    Complexity: generated.ComplexityRoot{},
  }))
  ```
- You should then be able to retrieve a user
  ```gql
  {
    user(id: "9b3b2c98-78ab-461a-9ed9-b99dabfc472b") {
      id
      name
      email
      profile
    }
  }
  ```
- If we try a more complex query we don't get any results for posts as we haven't yet created resolvers for linking our data together

```gql
{
  user(id: "9b3b2c98-78ab-461a-9ed9-b99dabfc472b") {
    id
    name
    email
    profile
    posts {
      id
      title
      content
    }
  }
}
```

- We need to tell `gqlgen` to generate resolvers for those fields in `gqlgen.yml`

```yml
models:
  User:
    fields:
      posts:
        resolver: true
      comments:
        resolver: true
  Post:
    fields:
      comments:
        resolver: true
      user:
        resolver: true
  Comment:
    fields:
      user:
        resolver: true
      post:
        resolver: true
```

- Run `gqlgen` again `go run github.com/99designs/gqlgen generate`
- Keep filling out resolvers until you can answer this query

```gql
{
  user(id: "cb4e6d63-64d2-4482-a46a-9c7c697ff72b") {
    id
    name
    email
    profile
    posts {
      id
      title
      content
      comments {
        content
        user {
          name
          profile
        }
      }
    }
  }
}
```

```json
{
  "data": {
    "user": {
      "id": "cb4e6d63-64d2-4482-a46a-9c7c697ff72b",
      "name": "Dylan Johnston",
      "email": "dylan.johnston@familyzone.com",
      "profile": "https://images.dog.ceo/breeds/bulldog-french/n02108915_3382.jpg",
      "posts": [
        {
          "id": "747bf43c-6be0-4fe2-9c59-658e1e3157d4",
          "title": "Avionics",
          "content": "According to all known laws of aviation ...",
          "comments": [
            {
              "content": "Great post man! Can't wait to read more.",
              "user": {
                "name": "Barry B Benson",
                "profile": "https://static.wikia.nocookie.net/beemovie/images/1/11/Barry-B-Benson.png/revision/latest?cb=20190513100654"
              }
            }
          ]
        }
      ]
    }
  }
}
```
