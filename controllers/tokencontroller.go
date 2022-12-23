package controllers

import (
	"fmt"
	"net/http"
	"restapi-auth/auth"
	"restapi-auth/database"
	"restapi-auth/models"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTToken struct {
	Token string `json:"token"`
}

// GenerateToken godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Authentication
// @Accept json
// @Produce json
// @Param {email,password} body TokenRequest true "Email email, Password password"
// @Success 200 {object} JWTToken
// @Router /user/token [post]
func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		fmt.Println("error in req", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check email is in DB
	record := database.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		fmt.Println("email not found")
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	// check password is correct
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		fmt.Println("password incorrect")
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
	context.JSON(http.StatusOK, JWTToken{
		Token: tokenString,
	})
}
