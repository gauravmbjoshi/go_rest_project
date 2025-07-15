package main

import (
	"example.com/go_rest_api_backend_project/db"
	"example.com/go_rest_api_backend_project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    db.InitDB()      // ✅ Initialize DB first
    server := gin.Default()
    routes.RegisterRoutes(server) // as gin.Defalut() already returns pointer so you don't need to specify pointer while passing argument
    server.Run(":8080")
}

