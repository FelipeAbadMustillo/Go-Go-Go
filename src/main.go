package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	//	"github.com/uptrace/bun"
	//	"github.com/uptrace/bun/dialect/pgdialect"
	//	"github.com/uptrace/bun/driver/pgdriver"
	//	"github.com/uptrace/bun/extra/bundebug"
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

type user struct {
	User_id          int    `json:"user_id"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Email            string `json:"email"`
	Date_entry       string `json:"fecha_alta"`
	Date_last_access string `json:"fecha_ultimo_acceso"`
}

var db *sql.DB

func main() {
	router := gin.Default()

	// 1. Make a database connection
	//ctx := context.Background()
	//dsn := "postgres://postgres:admin@localhost:5432/db_test?sslmode=disable"
	//sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	//db := bun.NewDB(sqldb, pgdialect.New())

	//db.AddQueryHook(bundebug.NewQueryHook(
	//	bundebug.WithVerbose(true),
	//	bundebug.FromEnv("BUNDEBUG"),
	//	))
	//err := db.Ping()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Connected to the database", ctx)

	var err error
	db, err = sql.Open("postgres", "postgres://postgres:admin@localhost:5432/db_test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Carga las plantillas desde la carpeta "templates" y los estilos desde la carpeta "static"
	router.LoadHTMLFiles("templates/search.html", "templates/index.html")
	router.Static("/static", "./static")

	//EndPoints
	router.GET("/", getIndex)
	router.GET("/search/:nameLike", getCoinsByNameLike)
	router.GET("/users", getUsers)
	router.POST("/users", createUser)

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

func getUsers(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	rows, err := db.Query("SELECT user_id,username,email,fecha_alta,fecha_ultimo_acceso FROM USUARIOS")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {

		var u user
		err := rows.Scan(&u.User_id, &u.Username, &u.Email, &u.Date_entry, &u.Date_last_access)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	context.IndentedJSON(http.StatusOK, users)

}

func createUser(context *gin.Context) {
	var newUser user

	if err := context.BindJSON(&newUser); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid  Request payload"})
		return
	}

	stmt, err := db.Prepare("INSERT INTO usuarios (user_id,username,password,email,fecha_alta,fecha_ultimo_acceso) values ($1, $2,$3,$4,$5,$6)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newUser.User_id, newUser.Username, newUser.Password, newUser.Email, newUser.Date_entry, newUser.Date_last_access); err != nil {
		log.Fatal(err)
	}

	context.JSON(http.StatusCreated, newUser)

}
