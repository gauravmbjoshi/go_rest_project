package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) { // as we are pointing to the server we do not need to return anything

	server.GET("/events", getEvents)
    server.GET("/events/:id", getEvent) // dynamic path handler setup using /:id
    server.POST("/events", createEvent)

}