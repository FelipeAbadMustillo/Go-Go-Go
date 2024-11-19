package handlers

import (
	"github.com/gin-gonic/gin"
	//ginmiddle "github.com/gin-gonic/gin/middle"
	//"github.com/Go-Go-Go/internal/middleware"
)

func Handler(router *gin.Engine) {
	// Loads templates and static services
	templatesPath := "../../web/templates"
	router.LoadHTMLFiles(
		templatesPath+"/index.html",
		templatesPath+"/coin.html",
	)
	router.Static("static", "../../web/static")

	router.GET("/", getIndex)
	router.GET("/coins", getCoins)
}
