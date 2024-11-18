package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
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

// CoinDataResponse representa el formato de respuesta del endpoint "Coins" de CG.
type CoinDataResponse struct {
	ID          string   `json:"id"`
	Symbol      string   `json:"symbol"`
	Name        string   `json:"name"`
	Categories  []string `json:"categories"`
	Description struct {
		English string `json:"en"`
	} `json:"description"`
	Image struct {
		Thumbnail string `json:"thumb"`
		Small     string `json:"small"`
		Large     string `json:"large"`
	} `json:"image"`
	Origin      string `json:"country_origin"`
	GenesisDate string `json:"genesis_date"`
	WatchList   int    `json:"watchlist_portfolio_users"`
	Rank        int    `json:"market_cap_rank"`
}

func main() {
	router := gin.Default()

	// Carga las plantillas desde la carpeta "templates" y los estilos desde la carpeta "static"
	router.LoadHTMLFiles("templates/search.html", "templates/index.html", "templates/moneda.html")
	router.Static("/static", "./static")

	//EndPoints
	router.GET("/", getIndex)
	router.GET("/coin/:id", getCoinByID)
	router.GET("/coins/:nameLike", getCoinsByNameLike)

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

func getCoinByID(context *gin.Context) {
	//Busca la moneda segun el id enviado por url y muestra todos sus datos en una pagina aparte.
	var coinData CoinDataResponse
	id := context.Param("id")

	//Ahora hacer la request a CG /coins/:id con los ids filtrados para despues mostrar la lista
	params := url.Values{}
	params.Add("localization", "false")
	params.Add("tickers", "false")
	params.Add("community_data", "false")
	params.Add("developer_data", "false")

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s?%s", ROOT+"/coins/"+id, params.Encode()), nil) //Averiguar bien que es lo de _ que maneja errores
	req.Header.Add("accept", "application/json")                                                    //Averiguar bien porque es necesario este header
	req.Header.Add(HEADER_AUTH_KEY, API_KEY)                                                        //Averiguar la forma mas segura de pasar la api key

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close() //Averiguar bien que es defer para documentar como caracteristica del lenguaje que creo que hayu algo ahi
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &coinData)

	context.HTML(http.StatusOK, "moneda.html", coinData) //Explicar bien como funcionan los templates
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
