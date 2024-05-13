package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse requested data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user"})
		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{
			"Message": "User created",
			"user":    user,
		})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse requested data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user, wrong credential"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user, wrong with token generation"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "loged in!", "token": token})
}
