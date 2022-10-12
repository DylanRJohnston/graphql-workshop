package main

import (
	"database/sql"
	"graphql-workshop/src/adapters"
	"graphql-workshop/src/routes"
	"graphql-workshop/src/usecases"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		panic(err)
	}

	users := adapters.NewSQLiteUserRepository(db)
	comments := adapters.NewSQLiteCommentRepository(db)
	posts := adapters.NewSQLitePostRepository(db)

	deps := usecases.Dependencies{
		Users:    users,
		Comments: comments,
		Posts:    posts,
	}

	gin := routes.New(deps)

	err = gin.Run()
	if err != nil {
		panic(err)
	}
}
