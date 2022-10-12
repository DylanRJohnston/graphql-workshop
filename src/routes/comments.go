package routes

import (
	"graphql-workshop/src/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

// No business logic should exist in these gin controllers
// only packing and unpacking of data to and from the http layer

func (d *Dependences) CommentGet(c *gin.Context) {
	commentID := c.Param("commentID")

	user, err := d.usecase.CommentGet(c, commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (d *Dependences) CommentCreate(c *gin.Context) {
	postID := c.Param("postID")
	commentFields := usecases.CommentCreateFields{}

	err := c.BindJSON(&commentFields)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	user, err := d.usecase.CommentCreate(c, postID, commentFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
