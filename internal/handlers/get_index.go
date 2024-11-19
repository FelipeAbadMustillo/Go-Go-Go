package handlers

import (
	"net/http"
	"net/url"

	"github.com/Go-Go-Go/api"
	"github.com/Go-Go-Go/internal/tools"
	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"
)

func getIndex(context *gin.Context) {
	//Renders index with trending coins basic information in a table.
	var trendingCoins api.TrendingResponse
	var indexData api.IndexTemplateData

	tools.CGRequest("/search/trending", url.Values{}, &trendingCoins)

	indexData.Title = "Monedas en tendencia"
	for _, coin := range trendingCoins.Coins {
		coin.Item.Logo = coin.Item.LogoTrending
		coin.Item.PriceBTC = coin.Item.PriceBTCTrending
		indexData.Coins = append(indexData.Coins, coin.Item)
	}
	context.HTML(http.StatusOK, "index.html", indexData)
}
