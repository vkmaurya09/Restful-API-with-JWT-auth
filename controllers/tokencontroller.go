package controllers

import (
	"net/http"
	"restapi-auth/auth"
	"restapi-auth/database"
	"restapi-auth/models"

	"github.com/gin-gonic/gin"
)

// struct for login
type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// genrating token for user validation
func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check email in in DB
	record := database.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	// check password is correct
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	// generate token
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
		context.Abort()
		return
	}
	// return token
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
