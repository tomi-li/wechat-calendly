package controllers

import (
	"calendly/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EventController struct{}

var eventModel = new(models.Event)

func (u EventController) Retrieve(c *gin.Context) {
	events := eventModel.ListAll()
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": events})
	return
}

func (u EventController) Create(c *gin.Context) {
	var event models.Event
	c.Bind(&event)
	eventModel.Create(event)
	return
}
