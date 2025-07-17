package middleware

import (
	"net/http"

	"example.com/go_rest_api_backend_project/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		// as this middle ware will be used by other users and for deferent functions so we use abort for that instance so to not affect other workings
		return
	}
	userId, err := utils.CheckToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}
	context.Set("userId",userId)
	context.Next()
}