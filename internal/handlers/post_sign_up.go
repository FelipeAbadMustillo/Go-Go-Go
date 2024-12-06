package handlers

import (
	"net/http"

	"github.com/Go-Go-Go/api"
	middleware "github.com/Go-Go-Go/internal/middleware"
	"github.com/Go-Go-Go/internal/tools"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func postSignUp(context *gin.Context) {
	// Structure of the body request with user input.
	var userData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Links request´s body with credentials struct.
	err := context.ShouldBindJSON(&userData)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate User Input with DB query for username-password match.
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil { //Mejorar los errores y fijarse esto
		log.Error("Could not connect to DB")
		api.InternalErrorHandler(context.Writer)
		return
	}

	newUser, err := (*database).CreateUser(userData.Username, userData.Email, userData.Password) //Aca pasar un objeto user
	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(context.Writer, err)
		return
	}

	// Generate JWT
	token, err := middleware.GenerateJWT(newUser.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Configurar la cookie para el token JWT
	context.SetCookie(
		"authToken",                   // cookie name
		token,                         // cookie value
		api.JWT_EXPIRATION_MINUTES*60, // cookie duration in seconds
		"/",                           // Path (all routes) //Cambiar a solo /account???
		"",                            // Dominio (empty for current domain)
		true,                          // Secure (only send trough HTTPS)
		true,                          // HttpOnly (not JS accesible)
	)

	context.JSON(http.StatusOK, gin.H{"message": "Sign Up successful"})
}
