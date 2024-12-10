package main

import (
	"github.com/Go-Go-Go/internal/handlers"
	"github.com/Go-Go-Go/internal/tools"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	router := gin.Default()

	database, err := tools.NewDatabase()
	if err != nil {
		log.Error(err)
	}
	defer database.Close()

	handlers.Handler(router, &database)

	err = router.Run("localhost:8080")
	if err != nil {
		log.Error(err)
	}
}
