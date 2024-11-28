package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

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
	timeTo := time.Now().Add(-7 * 24 * time.Hour).Unix()
	timeFrom := time.Now().Unix()
	var coinData api.CoinDataResponse
	var coinHistory api.CoinHistoryResponse
	var templateData api.CoinTemplateData

	params := url.Values{}
	params.Add("localization", "false")
	params.Add("tickers", "false")
	params.Add("community_data", "false")
	params.Add("developer_data", "false")
	tools.CGRequest("/coins/"+id, params, &coinData)

	for k := range params {
		delete(params, k)
	}

	params.Add("vs_currency", "usd")
	params.Add("from", fmt.Sprintf("%d", timeTo))
	params.Add("to", fmt.Sprintf("%d", timeFrom))
	tools.CGRequest("/coins/"+id+"/market_chart/range", params, &coinHistory)

	templateData.Name = coinData.Name
	templateData.CurrentPrice = coinData.MarketData.CurrentPrice.USD
	templateData.MaxPrice = coinData.MarketData.AllTimeHigh.USD
	templateData.MinPrice = coinData.MarketData.AllTimeLow.USD
	templateData.Logo = coinData.Image.Small
	for index := 0; index < len(coinHistory.Prices); index++ {
		templateData.History = append(templateData.History, api.HistoryTemplateData{
			Time:         time.Unix(int64(coinHistory.Prices[index][api.TIME])/1000, 0).Format("02/01/2006"),
			Price:        coinHistory.Prices[index][api.VALUE],
			MarketCap:    coinHistory.MarketCaps[index][api.VALUE],
			TotalVolumes: coinHistory.TotalVolumes[index][api.VALUE],
		})
	}

	context.HTML(http.StatusOK, "coin.html", templateData)
}
