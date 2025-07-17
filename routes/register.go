package routes

import (
	"net/http"
	"strconv"

	"example.com/go_rest_api_backend_project/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId,err := strconv.ParseInt(context.Param("id"),10,64)

    if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event,err := models.GetEventById(eventId)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"could not fetch event"})
		return
	}

	err = event.Register(userId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"could not register user for event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message":"registered for event"})
}

func cancleRegistration(context *gin.Context) {
	
}