package handlers

import (
	"net/http"

	"github.com/Go-Go-Go/api"
	"github.com/gin-gonic/gin"
)

func getSignUp(context *gin.Context) {
	//Renders login form.
	var templateData api.SignUpTemplateData
	context.HTML(http.StatusOK, "sign_up.html", templateData)
}
