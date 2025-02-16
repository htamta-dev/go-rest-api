package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/htamta-dev/go-rest-api/models"
	"github.com/htamta-dev/go-rest-api/utils"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create a user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
