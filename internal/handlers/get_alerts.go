package handlers

import (
	"net/http"

	"github.com/Go-Go-Go/api"
	"github.com/Go-Go-Go/internal/tools"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func getAlerts(context *gin.Context) {
	//Renders user alerts.
	username := context.MustGet("username")
	var err error

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil { //Mejorar los errores y fijarse esto
		api.InternalErrorHandler(context.Writer)
		return
	}

	var userAlerts *[]tools.Alerts
	userAlerts = (*database).GetUserAlerts(username.(string))
	if userAlerts == nil {
		log.Error("User not found on DB")
		api.InternalErrorHandler(context.Writer)
		return
	}

	context.HTML(http.StatusOK, "get_alerts.tmpl", userAlerts)
}
