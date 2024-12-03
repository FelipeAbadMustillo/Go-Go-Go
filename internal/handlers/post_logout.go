package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postLogout(context *gin.Context) {
	// Set cookie duration to expire auutomatically.
	context.SetCookie(
		"authToken", // cookie name
		"",          // cookie value
		-1,          // cookie duration in seconds
		"/",         // Path (all routes) //Cambiar a solo /account???
		"",          // Dominio (empty for current domain)
		true,        // Secure (only send trough HTTPS)
		true,        // HttpOnly (not JS accesible)
	)

	context.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
