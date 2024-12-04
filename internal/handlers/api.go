package handlers

import (
	"database/sql"

	middleware "github.com/Go-Go-Go/internal/middleware"
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
		templatesPath+"/sign_up.html",
		templatesPath+"/get_alerts.tmpl",
		templatesPath+"/footer.tmpl",
		templatesPath+"/header.tmpl",
		templatesPath+"/navbar.tmpl",
		templatesPath+"/view_alerts.tmpl",
	)
	router.Static("static", "../../web/static")

	router.GET("/", getIndex)
	router.GET("/coins", getCoins)
	router.GET("/signup", getSignUp)
	router.POST("/signup", postSignUp)
	router.GET("/login", getLogin)
	router.POST("/login", postLogin)

	secured := router.Group("/account")
	secured.Use(middleware.JWTAuthorization())
	secured.GET("/alerts", getAlerts)
	secured.POST("/logout", postLogout)
}
