package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	"html/template"
)

// moneda representa información al respecto de una bitCoin.
type moneda struct {
	ID         int     //`json:"id"`
	Nombre     string  //`json:"nombre"`
	Cotizacion float64 //`json:"cotizacion"`
	Logo       string  //`json:"logo"`
}

// slice de monedas para hacer pruebas.
var monedas = []moneda{
	{ID: 1, Nombre: "BitCoin", Cotizacion: 56.99, Logo: ""},
	{ID: 2, Nombre: "Ethereum", Cotizacion: 17.99, Logo: ""},
	{ID: 3, Nombre: "DogeCoin", Cotizacion: 39.99, Logo: ""},
}

func main() {
	router := gin.Default()

	// 1. Make a database connection
	ctx := context.Background()
	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return
	}
	fmt.Println("Connected to the database", ctx)

	// Carga las plantillas desde la carpeta "templates" y los estilos desde la carpeta "static"
	router.SetHTMLTemplate(template.Must(template.ParseFiles("templates/index.html")))
	router.Static("/static", "./static")

	//EndPoints
	router.GET("/Monedas", getMonedas)

	router.Run("localhost:8080")
}

// getMonedas devuelve la lista de monedas en formato JSON.
func getMonedas(conversor *gin.Context) {
	conversor.HTML(http.StatusOK, "index.html", monedas[0])
	//conversor.IndentedJSON(http.StatusOK, monedas)
}
