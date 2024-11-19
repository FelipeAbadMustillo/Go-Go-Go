package tools

//log "github.com/sirupsen/logrus"

//Database collections
//structs de respuestas de la db

type DatabaseInterface interface {
	//Distintos gets que dvuelven los structs de arriba desde la db
	SetupDatabase() error
}

// func NewDatabase() (*DatabaseInterface, error) {

// }
