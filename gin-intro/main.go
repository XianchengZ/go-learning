package main

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":5029")
}

func getEvents(context *gin.Context) {
	var events = models.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(
		http.StatusCreated,
		gin.H{
			"Message": "Event created",
			"event":   event,
		})
}
