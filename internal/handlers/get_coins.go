package handlers

import (
	"net/http"
	"net/url"

	"github.com/Go-Go-Go/api"
	"github.com/Go-Go-Go/internal/tools"
	"github.com/gin-gonic/gin"
)

func getCoins(context *gin.Context) {
	//Searches coin information.
	nameLike := context.Query("nameLike")
	id := context.Query("id")

	if id != "" {
		byId(context, id)
	} else {
		byNameLike(context, nameLike)
	}
}

func byNameLike(context *gin.Context, nameLike string) {
	//Renders search.html template with a table of coins wich name contains the query param.
	var coinsMarketData []api.CoinBasicInformation
	var indexData api.IndexTemplateData

	IDsParam := tools.FilterCoinIDs(nameLike)

	params := url.Values{}
	params.Add("vs_currency", "usd") // Pensar bien que moneda usar
	params.Add("ids", IDsParam)

	tools.CGRequest("/coins/markets", params, &coinsMarketData)

	indexData.Title = "Resultados de la busqueda '" + nameLike + "'"
	indexData.Coins = coinsMarketData
	context.HTML(http.StatusOK, "index.html", indexData)
}

func byId(context *gin.Context, id string) {
	//Renders coin.html template with the extended coin data of the query param.
	var coinData api.CoinDataResponse

	params := url.Values{}
	params.Add("localization", "false")
	params.Add("tickers", "false")
	params.Add("community_data", "false")
	params.Add("developer_data", "false")

	tools.CGRequest("/coins/"+id, params, &coinData)

	context.HTML(http.StatusOK, "coin.html", coinData)
}
