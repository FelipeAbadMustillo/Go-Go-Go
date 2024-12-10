package handlers

import (
	"net/http"

	"github.com/Go-Go-Go/api"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func getAlerts(context *gin.Context) {
	//Renders user alerts.
	username := context.MustGet("username")
	var err error

	userAlerts, err := DB.GetUserAlerts(context, username.(string))
	if err != nil { //Mejorar los errores y fijarse esto
		log.Error(err)
		api.RequestErrorHandler(context.Writer, err)
		return
	}

	context.HTML(http.StatusOK, "get_alerts.tmpl", userAlerts)
}
