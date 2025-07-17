package routes

import (
	"net/http"

	"example.com/go_rest_api_backend_project/models"
	"example.com/go_rest_api_backend_project/utils"
	"github.com/gin-gonic/gin"
)
func signup(context *gin.Context){
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user format."})
        return
    }
	err := user.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user, please try again."})
        return
	}
	context.JSON(http.StatusOK, gin.H{"message":"user created successfully!!!"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if  err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user format."})
        return
    }
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message":err.Error()})
        return
	}
	token,err := utils.GenerateToken(user.Email,user.ID)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message":err.Error()})
        return
	}

	context.JSON(http.StatusOK, gin.H{"message":"Login successful!","token":token})
}