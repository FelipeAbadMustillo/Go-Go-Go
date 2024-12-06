package handlers

import (
	middleware "github.com/Go-Go-Go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Handler(router *gin.Engine) {
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
		templatesPath+"/new_alert.tmpl",
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
	secured.POST("/logout", postLogout)
	secured.GET("/alerts", getAlerts)

	secured.GET("/alerts/new", getNewAlert)
	secured.POST("/alerts/new", postNewAlert)
}
