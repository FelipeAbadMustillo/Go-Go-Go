package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ROOT string = "https://api.coingecko.com/api/v3"
const HEADER_AUTH_KEY string = "x_cg_demo_api_key"
const API_KEY string = "CG-4YbsrgnB5a45UdF9gYyqXJfi"

// CoinGeckoTrendingResponse representa el formato de respuesta del endpoint "Trending" de CG.
type CoinGeckoTrendingResponse struct {
	Monedas []struct {
		Moneda struct {
			Nombre          string  `json:"name"`
			Abreviatura     string  `json:"symbol"`
			Rango           int     `json:"market_cap_rank"`
			Logo            string  `json:"small"`
			CotizacionEnBTC float64 `json:"price_btc"`
		} `json:"item"`
	} `json:"coins"`
}

func main() {
	router := gin.Default()

	// Carga las plantillas desde la carpeta "templates" y los estilos desde la carpeta "static"
	router.SetHTMLTemplate(template.Must(template.ParseFiles("templates/index.html")))
	router.Static("/static", "./static")

	//EndPoints
	router.GET("/", getIndex)

	router.Run("localhost:8080")
}

func getIndex(conversor *gin.Context) {
	//En el index se muestran las monedas en tendencia con su información básica.
	url := ROOT + "/search/trending"
	var respuesta CoinGeckoTrendingResponse

	req, _ := http.NewRequest("GET", url, nil)   //Averiguar bien que es lo de _ que maneja errores
	req.Header.Add("accept", "application/json") //Averiguar bien porque es necesario este header
	req.Header.Add(HEADER_AUTH_KEY, API_KEY)     //Averiguar la forma mas segura de pasar la api key

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close() //Averiguar bien que es defer para documentar como caracteristica del lenguaje que creo que hayu algo ahi
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &respuesta)

	conversor.HTML(http.StatusOK, "index.html", respuesta.Monedas)
}
