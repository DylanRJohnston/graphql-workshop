package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var migrations []string = []string{
	`CREATE TABLE users (
		id       TEXT PRIMARY KEY NOT NULL,
		name     TEXT             NOT NULL,
		email    TEXT             NOT NULL,
		profile  TEXT                     ,
		birthday TEXT             NOT NULL
	)`,

	`CREATE TABLE posts (
		id       TEXT PRIMARY KEY NOT NULL,
		userID   TEXT             NOT NULL,
		title    TEXT             NOT NULL,
		content  TEXT             NOT NULL,

		FOREIGN KEY (userID) REFERENCES users(id)
	)`,

	`CREATE TABLE comments (
		id       TEXT PRIMARY KEY NOT NULL,
		userID   TEXT             NOT NULL,
		postID   TEXT             NOT NULL,
		content  TEXT             NOT NULL,

		FOREIGN KEY (userID) REFERENCES users(id),
		FOREIGN KEY (postID) REFERENCES posts(id)
	)`,
}

func main() {
	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		panic(err)
	}

	for _, migration := range migrations {
		_, err := db.Exec(migration)
		if err != nil {
			fmt.Printf("Error running migration %s\n", migration)
			panic(err)
		}
	}
}
