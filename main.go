package main

import (
	"net/http"

	"example.com/go_rest_api_backend_project/db"
	"example.com/go_rest_api_backend_project/models"
	"github.com/gin-gonic/gin"
)

func main() {
    db.InitDB()      // ✅ Initialize DB first
    server := gin.Default()
    server.GET("/events", getEvents)
    server.POST("/events", createEvent)
    server.Run(":8080")
}

func getEvents(context *gin.Context){
	events,err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
    var event models.Event
    if err := c.ShouldBindJSON(&event); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event format."})
        return
    }

    if err := event.Save(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event."})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully!", "event": event})
}