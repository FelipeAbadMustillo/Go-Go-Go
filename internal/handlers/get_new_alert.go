package handlers

import (
	"net/http"

	"github.com/Go-Go-Go/api"
	"github.com/gin-gonic/gin"
)

func getNewAlert(context *gin.Context) {
	//Renders new alert form.
	var templateData api.NewAlertTemplateData
	context.HTML(http.StatusOK, "new_alert.tmpl", templateData)
}
