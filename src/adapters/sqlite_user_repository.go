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

type SQLiteUserRepository struct {
	db *sql.DB
}

func NewSQLiteUserRepository(db *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{db}
}

// Interfaces in Go are duck typed, if we don't correctly implement the port, we want the error to show up here, not in main.go
// Go ahead a delete / break the methods below to see how the error gets reported here
var _implementsUserRepository ports.UserRepository = (*SQLiteUserRepository)(nil) // nolint:deadcode,unused,varcheck

var user_get_query string = `
	SELECT id, name, email, profile, birthday
	FROM users
	WHERE id = ?
`

func (s *SQLiteUserRepository) Get(ctx context.Context, id string) (models.User, error) {
	row := s.db.QueryRow(user_get_query, id)

	user := models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Profile, &user.Birthday)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

var user_create_query string = `
	INSERT INTO users(id, name, email, profile, birthday)
	VALUES(?, ?, ?, ?, ?)
	RETURNING id, name, email, profile, birthday
`

func (s *SQLiteUserRepository) Create(ctx context.Context, fields ports.UserCreate) (models.User, error) {
	id := uuid.New()
	row := s.db.QueryRow(user_create_query, id, fields.Name, fields.Email, fields.Profile, fields.Birthday)

	user := models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Profile, &user.Birthday)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
