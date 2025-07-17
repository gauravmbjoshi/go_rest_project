package routes

import (
	"net/http"
	"strconv"

	"example.com/go_rest_api_backend_project/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context){
	events,err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context){
    eventId,err := strconv.ParseInt(context.Param("id"),10,64)
    // this method provides you the parameter values from the endpoint
    // as id field in our database model is int so we need to parse it to Int so Parse int needs the value we need to convert and in what base int and what bitSide
    if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
    events,err := models.GetEventById(eventId)
    if err !=nil{
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
    }
    context.JSON(http.StatusOK,events)
}

func createEvent(c *gin.Context) {
    var event models.Event
    err := c.ShouldBindJSON(&event)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event format."})
        return
    }
    userId := c.GetInt64("userId") //this will convert the value in right type int64 which we want
    event.UserID = userId
    err = event.Save()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event."})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully!", "event": event})
}

func updateEvent(context *gin.Context){
	eventId,err := strconv.ParseInt(context.Param("id"),10,64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
    _,err = models.GetEventById(eventId)
    if err !=nil{
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
    }
	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent);
	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event format."})
        return
    }
	updateEvent.ID = eventId
	err = updateEvent.Update()
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
        return
    }
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}
func deleteEvent(context *gin.Context){
    eventId,err := strconv.ParseInt(context.Param("id"),10,64)
    if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
    event,err := models.GetEventById(eventId)
    if err !=nil{
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
    }

    err = event.Delete()
    if err !=nil{
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
    }
    context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}