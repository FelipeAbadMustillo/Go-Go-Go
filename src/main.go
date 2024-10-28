package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
