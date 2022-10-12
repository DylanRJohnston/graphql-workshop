package routes

import (
	"graphql-workshop/src/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

// No business logic should exist in these gin controllers
// only packing and unpacking of data to and from the http layer

func (d *Dependences) PostGet(c *gin.Context) {
	postID := c.Param("postID")

	user, err := d.usecase.PostGet(c, postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (d *Dependences) PostGetComments(c *gin.Context) {
	postID := c.Param("postID")

	user, err := d.usecase.PostGetComments(c, postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (d *Dependences) PostCreate(c *gin.Context) {
	userID := c.Param("userID")
	postFields := usecases.PostCreateFields{}

	err := c.BindJSON(&postFields)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	user, err := d.usecase.PostCreate(c, userID, postFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
