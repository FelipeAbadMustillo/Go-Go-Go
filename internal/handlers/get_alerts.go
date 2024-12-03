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
		log.Error(err)
		api.InternalErrorHandler(context.Writer)
		return
	}

	userAlerts := (*database).GetUserAlerts(username.(string))

	context.HTML(http.StatusOK, "get_alerts.tmpl", userAlerts)
}
