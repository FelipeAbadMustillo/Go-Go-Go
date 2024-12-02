package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database collections
type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Logins struct { //Esto no iria en auth.go?
	AuthToken string
	Username  string
}

type Alerts struct {
	IdAlert   string  `json:"id_alert"`
	Username  string  `json:"username"`
	Price     float64 `json:"price"`
	Condition string  `json:"condition"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
	IsActive  string  `json:"is_active"`
}

type DatabaseInterface interface {
	GetUser(username string, password string) *Users
	GetUserLoginDetails(username string) *Logins
	GetUserAlerts(username string) *[]Alerts
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{} //Despues cambiar a postgres online

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}

//----------------Utilidades de Base de datos de cesar------------------------------
/* // conexion a la base de datos
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
} */

/* //Query para alertas de usuario
rows, err := conectionDB.Query("SELECT id_alert,username,price,condition,start_date,end_date,is_active FROM ALERTS")
if err != nil {
	fmt.Println(err)
	log.Fatal(err)
	return
}
defer rows.Close()

var alerts []tools.Alert
for rows.Next() {

	var a tools.Alert
	err := rows.Scan(&a.Id_alert, &a.Username, &a.Price, &a.Condition, &a.Start_date, &a.End_date, &a.Is_active)
	if err != nil {
		log.Fatal(err)
	}

	alerts = append(alerts, a)
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}
*/
