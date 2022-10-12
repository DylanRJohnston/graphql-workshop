package adapters

import (
	"context"
	"database/sql"
	"graphql-workshop/src/models"
	"graphql-workshop/src/ports"

	_ "embed"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type SQLitePostRepository struct {
	db *sql.DB
}

func NewSQLitePostRepository(db *sql.DB) *SQLitePostRepository {
	return &SQLitePostRepository{db}
}

// Interfaces in Go are duck typed, if we don't correctly implement the port, we want the error to show up here, not in main.go
// Go ahead a delete / break the methods below to see how the error gets reported here
var _implementsPostRepository ports.PostRepository = (*SQLitePostRepository)(nil) // nolint:deadcode,unused,varcheck

var post_get_query string = `
	SELECT id, userID, title, content
	FROM posts
	WHERE id = ?
`

func (s *SQLitePostRepository) Get(ctx context.Context, id string) (models.Post, error) {
	row := s.db.QueryRow(post_get_query, id)

	post := models.Post{}
	err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Content)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

var post_create_query string = `
	INSERT INTO posts(id, userID, title, content)
	VALUES (?, ?, ?, ?)
	RETURNING id, userID, title, content
`

func (s *SQLitePostRepository) Create(ctx context.Context, fields ports.PostCreate) (models.Post, error) {
	id := uuid.New()
	row := s.db.QueryRow(post_create_query, id, fields.UserID, fields.Title, fields.Content)

	post := models.Post{}
	err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Content)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

var posts_for_user_query = `
	SELECT id, userID, title, content
	FROM posts
	WHERE userID = ?
`

func (s *SQLitePostRepository) ForUser(ctx context.Context, userID string) ([]models.Post, error) {
	rows, err := s.db.Query(posts_for_user_query, userID)
	if err != nil {
		return nil, err
	}

	posts := []models.Post{}

	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
