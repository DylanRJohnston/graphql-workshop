package routes

import (
	"graphql-workshop/src/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

// No business logic should exist in these gin controllers
// only packing and unpacking of data to and from the http layer

func (d *Dependences) UserGet(c *gin.Context) {
	userID := c.Param("userID")

	user, err := d.usecase.UserGet(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (d *Dependences) UserGetPosts(c *gin.Context) {
	userID := c.Param("userID")

	user, err := d.usecase.UserGetPosts(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (d *Dependences) UserCreate(c *gin.Context) {
	userFields := usecases.UserCreateFields{}

	err := c.BindJSON(&userFields)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	user, err := d.usecase.UserCreate(c, userFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
