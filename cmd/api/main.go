package main

import (
	"github.com/Go-Go-Go/internal/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	router := gin.Default()

	handlers.Handler(router)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Error(err)
	}
}
