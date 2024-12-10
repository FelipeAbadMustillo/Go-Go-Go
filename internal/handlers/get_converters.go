package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"
)

func getConverters(ctx *gin.Context) {
	//llamada al template para crear nuevo alerta

	//context.HTML(http.StatusOK, "user.tmpl", nil) //Explicar bien como funcionan los templates
	ctx.HTML(http.StatusOK, "get_converters.tmpl", gin.H{
		"title": "Currency Converter"})

}
