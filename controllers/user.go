package controllers

import (
	"calendly/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := userModel.GetByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}

func (u UserController) RetrieveAll(c *gin.Context) {
	users := userModel.ListAll()
	c.JSON(http.StatusOK, gin.H{"message": "Success", "users": users})
	c.Abort()
	return
}
