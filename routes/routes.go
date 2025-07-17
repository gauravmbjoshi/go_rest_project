package routes

import (
	"example.com/go_rest_api_backend_project/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) { // as we are pointing to the server we do not need to return anything

	server.GET("/events", getEvents)
    server.GET("/events/:id", getEvent) // dynamic path handler setup using /:id
    server.POST("/events",middleware.Authenticate, createEvent) // This is one way to apply middleware directly to a specific route

	// Alternatively, you can create a group and apply middleware to all its routes:
	// 1st you create a variable to create a group on base route
	authenticated := server.Group("/") 
	// 2nd you use middleware 
	authenticated.Use(middleware.Authenticate)
	// 3rd you protect the routes below that
	authenticated.PUT("/events/:id",updateEvent)
	authenticated.DELETE("/events/:id",deleteEvent)
	
	server.POST("/signup",signup)
	server.POST("/login",login)
}