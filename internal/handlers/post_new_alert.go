package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/Go-Go-Go/api"
	"github.com/Go-Go-Go/internal/tools"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func postNewAlert(context *gin.Context) {
	// Structure of the body request with user input.
	var userData struct {
		CoinName  string  `json:"coin_name"`
		Price     float64 `json:"price"`
		Condition string  `json:"condition"`
	}

	// Links request´s body with userData struct.
	err := context.ShouldBindJSON(&userData)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validates CoinName and assigns CoinID.
	coinID := tools.GetCoinID(userData.CoinName)
	if coinID == "" {
		err = errors.New("invalid coin name")
		log.Error(err)
		api.RequestErrorHandler(context.Writer, err)
		return
	}

	// Create new alert on DB.
	newAlert := tools.Alerts{}
	newAlert.Username = context.MustGet("username").(string)
	newAlert.Price = userData.Price
	newAlert.Condition = userData.Condition
	newAlert.StartDate = time.Now()
	newAlert.Status = "pending"
	newAlert.CoinID = coinID
	newAlert.CoinName = userData.CoinName

	err = DB.CreateAlert(context, &newAlert)
	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(context.Writer, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "New Alert Succesfully Created!"})
}
