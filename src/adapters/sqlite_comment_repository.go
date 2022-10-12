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

type SQLiteCommentRepository struct {
	db *sql.DB
}

func NewSQLiteCommentRepository(db *sql.DB) *SQLiteCommentRepository {
	return &SQLiteCommentRepository{db}
}

// Interfaces in Go are duck typed, if we don't correctly implement the port, we want the error to show up here, not in main.go
// Go ahead a delete / break the methods below to see how the error gets reported here
var _implementsCommentRepository ports.CommentRepository = (*SQLiteCommentRepository)(nil) // nolint:deadcode,unused,varcheck

var comment_get_query string = `
	SELECT id, userID, postID, content
	FROM comments
	WHERE id = ?
`

func (s *SQLiteCommentRepository) Get(ctx context.Context, id string) (models.Comment, error) {
	row := s.db.QueryRow(comment_get_query, id)

	comment := models.Comment{}
	err := row.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content)
	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

var comment_create_query string = `
	INSERT INTO comments(id, userID, postID, content)
	VALUES (?, ?, ?, ?)
	RETURNING id, userID, postID, content
`

func (s *SQLiteCommentRepository) Create(ctx context.Context, fields ports.CommentCreate) (models.Comment, error) {
	id := uuid.New()
	row := s.db.QueryRow(comment_create_query, id, fields.UserID, fields.PostID, fields.Content)

	comment := models.Comment{}
	err := row.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content)
	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

func (s *SQLiteCommentRepository) ForPost(ctx context.Context, postID string) ([]models.Comment, error) {
	panic("unimplemented")
}
