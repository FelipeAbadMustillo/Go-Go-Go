package handlers

import (
	"net/http"

	"github.com/Go-Go-Go/api"
	"github.com/gin-gonic/gin"
)

func getLogin(context *gin.Context) {
	//Renders login form.
	var templateData api.LoginTemplateData
	context.HTML(http.StatusOK, "login.html", templateData)
}
