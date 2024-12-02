package handlers

import (
	"database/sql"

	"github.com/Go-Go-Go/internal/tools"
	"github.com/gin-gonic/gin"
	//ginmiddle "github.com/gin-gonic/gin/middle"
	//"github.com/Go-Go-Go/internal/middleware"
)

var conectionDB *sql.DB

func Handler(router *gin.Engine) {

	conectionDB = tools.GetDBConection()

	// Loads templates and static services
	templatesPath := "../../web/templates"
	router.LoadHTMLFiles(
		templatesPath+"/index.html",
		templatesPath+"/coin.html",
	)
	router.LoadHTMLGlob(templatesPath + "/*.tmpl")

	router.Static("static", "../../web/static")

	router.GET("/", getIndex)
	router.GET("/coins", getCoins)
	router.GET("/alerts", getAlerts)
	router.GET("/alerts/:username", getAlertsByUsername)
	router.GET("/alerts/new", getAlertsNew)

	router.POST("/alerts", postAlerts)

}
