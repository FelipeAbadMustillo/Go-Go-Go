package tools

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresql struct{}

var dbPool *pgxpool.Pool

// Ver como proteger esto.
const DATABASE_URL = "postgresql://GoGoGo_owner:BFLaWMZI9Pn4@ep-long-band-a4fg2f6u.us-east-1.aws.neon.tech/GoGoGo?sslmode=require"

func (d *postgresql) SetupDatabase() error {
	// Creates Database connection pool to send requests.
	var err error

	dbPool, err = pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		return err
	}

	return nil
}

func (d *postgresql) CreateUser(context *gin.Context, newUser *Users) error {
	// Inserts new user into database.

	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"
	_, err := dbPool.Exec(context, query, newUser.Username, newUser.Email, newUser.Password)

	return err
}

func (d *postgresql) GetUser(context *gin.Context, username string, password string) (*Users, error) {
	// Selects User from database where username and password match.
	var user = Users{}

	query := "SELECT username, email, password FROM users WHERE username = $1 AND password = $2"
	err := dbPool.QueryRow(context, query, username, password).Scan(&user.Username, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *postgresql) CreateAlert(context *gin.Context, newAlert *Alerts) error {
	// Inserts new alert into database.

	// Validations

	query := `INSERT INTO alerts (username, price, condition, start_date, status, coin_id, coin_name) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := dbPool.Exec(context, query, newAlert.Username, newAlert.Price, newAlert.Condition, newAlert.StartDate, newAlert.Status, newAlert.CoinID, newAlert.CoinName)

	return err
}

func (d *postgresql) GetUserAlerts(context *gin.Context, username string) (*[]Alerts, error) {
	// Selects alerts from database where username matches.
	var alerts []Alerts

	query := `SELECT id, username, price, condition, start_date, status, coin_id, coin_name
			  FROM alerts
			  WHERE username = $1`

	rows, err := dbPool.Query(context, query, username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var alert Alerts
		err := rows.Scan(
			&alert.ID,
			&alert.Username,
			&alert.Price,
			&alert.Condition,
			&alert.StartDate,
			&alert.Status,
			&alert.CoinID,
			&alert.CoinName,
		)
		if err != nil {
			return nil, err
		}

		alerts = append(alerts, alert)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &alerts, nil
}

func (d *postgresql) Close() {
	dbPool.Close()
}
