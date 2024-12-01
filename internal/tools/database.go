package tools

//log "github.com/sirupsen/logrus"

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	//	"github.com/uptrace/bun"
	//	"github.com/uptrace/bun/dialect/pgdialect"
	//	"github.com/uptrace/bun/driver/pgdriver"
	//	"github.com/uptrace/bun/extra/bundebug"
)

//Database collections
//structs de respuestas de la db

/*type DatabaseInterface interface {
	//Distintos gets que dvuelven los structs de arriba desde la db
	SetupDatabase() error
}*/

// func NewDatabase() (*DatabaseInterface, error) {

// }

type Alert struct {
	Id_alert   string  `json:"id_alert"`
	Username   string  `json:"username"`
	Price      float64 `json:"price"`
	Condition  string  `json:"condition"`
	Start_date string  `json:"start_date"`
	End_date   string  `json:"end_date"`
	Is_active  string  `json:"is_active"`
	Coin_Code  string  `json:"coin_code"`
}

// conexion a la base de datos
const (
	dbdriver   = "postgres"
	dbhostname = "localhost"
	dbport     = "5432"
	dbuser     = "postgres"
	dbpassword = "admin"
	dbname     = "db_test"
)

func GetDBConection() (conection *sql.DB) {
	var err error
	connectionStr := dbdriver + "://" + dbuser + ":" + dbpassword + "@" + dbhostname + ":" + dbport + "/" + dbname + "?sslmode=disable"
	//fmt.Println(connectionStr)
	conection, err = sql.Open(dbdriver, connectionStr)
	if err != nil {
		log.Fatal(err)
	}
	return conection
}
