package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"graphql-workshop/src/adapters"
	"graphql-workshop/src/models"
	"graphql-workshop/src/routes"
	"graphql-workshop/src/usecases"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	_ "github.com/mattn/go-sqlite3"
)

func setup() *gin.Engine {
	db, err := sql.Open("sqlite3", "../../blog.db")
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

	return routes.New(deps)
}

func newRequest(t *testing.T, method string, path string, object interface{}) (*httptest.ResponseRecorder, *http.Request) {
	buffer := &bytes.Buffer{}
	err := json.NewEncoder(buffer).Encode(object)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, path, buffer)
	if err != nil {
		t.Fatal(err)
	}

	return w, req
}

func Test_Integration(t *testing.T) {
	gin := setup()

	w, req := newRequest(t, "POST", "/users", usecases.UserCreateFields{
		Name:  "Dylan Johnston",
		Email: "dylan.johnston@familyzone.com",
		//  https://dog.ceo/dog-api/
		Profile:  "https://images.dog.ceo/breeds/bulldog-french/n02108915_3382.jpg",
		Birthday: time.Now().String(),
	})

	gin.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	blogUser := models.User{}
	err := json.NewDecoder(w.Body).Decode(&blogUser)
	if err != nil {
		t.Fatal(err)
	}

	w, req = newRequest(t, "POST", "/users", usecases.UserCreateFields{
		Name:     "Barry B Benson",
		Email:    "B.B.B@gmail.com",
		Profile:  "https://static.wikia.nocookie.net/beemovie/images/1/11/Barry-B-Benson.png/revision/latest?cb=20190513100654",
		Birthday: time.Now().String(),
	})

	gin.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	commentUser := models.User{}
	err = json.NewDecoder(w.Body).Decode(&commentUser)
	if err != nil {
		t.Fatal(err)
	}

	w, req = newRequest(t, "POST", fmt.Sprintf("/users/%s/posts", blogUser.ID), usecases.PostCreateFields{
		Title:   "Avionics",
		Content: "According to all known laws of aviation ...",
	})

	gin.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	post := models.Post{}
	err = json.NewDecoder(w.Body).Decode(&post)
	if err != nil {
		t.Fatal(err)
	}

	w, req = newRequest(t, "POST", fmt.Sprintf("/posts/%s/comments", post.ID), usecases.CommentCreateFields{
		UserID:  commentUser.ID,
		Content: "Great post man! Can't wait to read more.",
	})

	gin.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	comment := models.Comment{}
	err = json.NewDecoder(w.Body).Decode(&comment)
	if err != nil {
		t.Fatal(err)
	}

	w, req = newRequest(t, "GET", fmt.Sprintf("/users/%s", blogUser.ID), nil)
	gin.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	expectedUser := models.User{}
	err = json.NewDecoder(w.Body).Decode(&expectedUser)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, blogUser, expectedUser)

	w, req = newRequest(t, "GET", fmt.Sprintf("/posts/%s", post.ID), nil)
	gin.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	expectedPost := models.Post{}
	err = json.NewDecoder(w.Body).Decode(&expectedPost)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, post, expectedPost)

	w, req = newRequest(t, "GET", fmt.Sprintf("/comments/%s", comment.ID), nil)
	gin.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	expectedComment := models.Comment{}
	err = json.NewDecoder(w.Body).Decode(&expectedComment)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, comment, expectedComment)
}
