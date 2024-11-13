package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

const ROOT string = "https://api.coingecko.com/api/v3"
const HEADER_AUTH_KEY string = "x_cg_demo_api_key"
const API_KEY string = "CG-4YbsrgnB5a45UdF9gYyqXJfi"

// TrendingResponse representa el formato de respuesta del endpoint "Trending" de CG.
type TrendingResponse struct {
	Coins []struct {
		Item struct {
			Name     string  `json:"name"`
			Symbol   string  `json:"symbol"`
			Rank     int     `json:"market_cap_rank"`
			Logo     string  `json:"small"`
			PriceBTC float64 `json:"price_btc"`
		} `json:"item"`
	} `json:"coins"`
}

// CoinsListResponse representa el formato de respuesta del endpoint "List" de CG.
type CoinListResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CoinMarketDataResponse representa el formato de respuesta del endpoint "Market" de CG.
type CoinMarketDataResponse struct {
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	Rank     int     `json:"market_cap_rank"`
	Logo     string  `json:"image"`
	PriceBTC float64 `json:"current_price"`
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
	router.LoadHTMLFiles("templates/search.html", "templates/index.html")
	router.Static("/static", "./static")

	//EndPoints
	router.GET("/", getIndex)
	router.GET("/search/:nameLike", getCoinsByNameLike)

	router.Run("localhost:8080")
}

func getIndex(context *gin.Context) {
	//En el index se muestran las monedas en tendencia con su información básica.
	var trendingCoins TrendingResponse

	CGRequest("/search/trending", &trendingCoins)

	context.HTML(http.StatusOK, "index.html", trendingCoins.Coins) //Explicar bien como funcionan los templates
}

func getCoinsByNameLike(context *gin.Context) {
	//Busca las monedas que contengan dentro de su nombre el string argumentado y se muestran resultados.
	var IDsParam string
	var coinsMarketData []CoinMarketDataResponse
	nameLike := context.Param("nameLike")

	IDsParam = getCoinIDs(nameLike)

	//Ahora hacer la request a CG /coins/markets con los ids filtrados para despues mostrar la lista
	params := url.Values{}
	params.Add("vs_currency", "usd") // Pensar bien que moneda usar
	params.Add("ids", IDsParam)

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s?%s", ROOT+"/coins/markets", params.Encode()), nil) //Averiguar bien que es lo de _ que maneja errores
	req.Header.Add("accept", "application/json")                                                        //Averiguar bien porque es necesario este header
	req.Header.Add(HEADER_AUTH_KEY, API_KEY)                                                            //Averiguar la forma mas segura de pasar la api key

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close() //Averiguar bien que es defer para documentar como caracteristica del lenguaje que creo que hayu algo ahi
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &coinsMarketData)

	context.HTML(http.StatusOK, "search.html", coinsMarketData) //Explicar bien como funcionan los templates
}

func getCoinIDs(nameFilter string) string {
	//Devuelve un string separado por comas de todos los ids de las monedas cuyo nombre contiene nameFilter.
	var coinIDsMap []CoinListResponse
	var filteredIDs []string
	coinLimit := 5

	CGRequest("/coins/list", &coinIDsMap)

	for _, coin := range coinIDsMap {
		if strings.Contains(strings.ToLower(coin.Name), strings.ToLower(nameFilter)) && len(filteredIDs) <= coinLimit {
			filteredIDs = append(filteredIDs, coin.ID)
		}
	}
	return strings.Join(filteredIDs, ",")
}

func CGRequest(url string, response any) {
	req, _ := http.NewRequest("GET", ROOT+url, nil) //Averiguar bien que es lo de _ que maneja errores
	req.Header.Add("accept", "application/json")    //Averiguar bien porque es necesario este header
	req.Header.Add(HEADER_AUTH_KEY, API_KEY)        //Averiguar la forma mas segura de pasar la api key

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close() //Averiguar bien que es defer para documentar como caracteristica del lenguaje que creo que hayu algo ahi
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, response)
}
