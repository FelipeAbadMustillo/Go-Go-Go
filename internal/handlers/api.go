package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

var conectionDB *sql.DB

func Handler(router *gin.Engine) {

	//conectionDB = tools.GetDBConection()

	// Loads templates and static services
	templatesPath := "../../web/templates"
	router.LoadHTMLFiles(
		templatesPath+"/index.html",
		templatesPath+"/coin.html",
		templatesPath+"/login.html",
	)
	router.LoadHTMLGlob(templatesPath + "/*.tmpl")

	router.Static("static", "../../web/static")

	router.GET("/", getIndex)
	router.GET("/coins", getCoins)
	router.GET("/login", getLogin)
	router.POST("/login", postLogin)

	secured := router.Group("/account")
	secured.Use(JWTAuthorization())
	secured.GET("/alerts", getAlerts)

	router.POST("/alerts", postAlerts)
	router.GET("/viewalerts", viewAlerts) //Esto es el get del form para crear nuevas alertas

}
