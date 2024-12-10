package handlers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/Go-Go-Go/api"
	"github.com/Go-Go-Go/internal/tools"
	"github.com/gin-gonic/gin"
)

func getPrices(context *gin.Context) {

	amount := context.Query("amount")
	ids := context.Query("ids")
	vs_currencies := context.Query("vs_currencies")

	//byidVsCurrency(context, ids, vs_currencies)
	byidVsCurrency(context, amount, ids, vs_currencies)

}

func byidVsCurrency(context *gin.Context, amount string, ids string, vs_currencies string) {
	//Renders search.html template with a table of coins wich name contains the query param.

	var priceData map[string]any //map[string]json.RawMessage  //interface []interface{}
	var coinData api.CoinPriceTemplateData

	params := url.Values{}
	params.Add("ids", ids)
	params.Add("vs_currencies", vs_currencies) // Pensar bien que moneda usar

	tools.CGRequest("/simple/price", params, &priceData)

	for key_id, vs_currencies_data := range priceData {
		coinData.Base = key_id
		for key_vs_currencies, value := range vs_currencies_data.(map[string]any) {
			coinData.BaseTo = key_vs_currencies

			if amountValue, err := strconv.ParseFloat(amount, 64); err == nil {
				coinData.Amount = amountValue
				coinData.Rates.Vs_currencies = amountValue * value.(float64)
			}

		}
	}

	context.IndentedJSON(http.StatusOK, coinData)

	//https://api.coingecko.com/api/v3/simple/price?ids=moonpot&vs_currencies=ltc

}
