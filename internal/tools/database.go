package tools

import (
	"time"

	"github.com/Go-Go-Go/api"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Database collections
type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Alerts struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Price     float64   `json:"price"`
	Condition string    `json:"condition"`
	StartDate time.Time `json:"start_date"`
	Status    string    `json:"status"`
	CoinID    string    `json:"coin_id"`
	CoinName  string    `json:"coin_name"`
}

type DatabaseInterface interface {
	SetupDatabase() error
	CreateUser(context *gin.Context, newUser *Users) error
	GetUser(context *gin.Context, username string, password string) (*Users, error)
	CreateAlert(context *gin.Context, newAlert *Alerts) error
	GetUserAlerts(context *gin.Context, username string) (*[]Alerts, error)
	Close()
}

func NewDatabase() (DatabaseInterface, error) {
	var database DatabaseInterface = &postgresql{}
	var mockDB DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)

		if api.PRODUCTION {
			return nil, err
		} else {
			log.Info("now using mockDB")
			return mockDB, nil
		}
	}

	return database, nil
}
