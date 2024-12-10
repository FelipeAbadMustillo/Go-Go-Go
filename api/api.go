package api

import (
	"encoding/json"
	"net/http"
)

// IndexTemplateData specifies data passed to index.html template.
type IndexTemplateData struct {
	Title string
	Coins []CoinBasicInformation
}

// CoinBasicInformation specifies the basic info for a sigle coin.
type CoinBasicInformation struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	Rank             int     `json:"market_cap_rank"`
	Logo             string  `json:"image"`
	PriceBTC         float64 `json:"current_price"`
	LogoTrending     string  `json:"small"`
	PriceBTCTrending float64 `json:"price_btc"`
}

// TrendingResponse specifies the structure of the "/trending" CG endpoint
type TrendingResponse struct {
	Coins []struct {
		Item CoinBasicInformation `json:"item"`
	} `json:"coins"`
}

// CoinsListResponse representa el formato de respuesta del endpoint "List" de CG.
type CoinListResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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

type CoinPriceTemplateData2 struct {
	IdFrom    string
	IdTo      string
	PriceFrom float64
	PriceTo   float64
}

// //{"amount":10.0,"base":"BRL","date":"2024-12-06","rates":{"AUD":2.5975}}

type CoinPriceTemplateData struct {
	Base   string     `json:"base"`
	Amount float64    `json:"amount"`
	BaseTo string     `json:"baseTo"`
	Rates  Currencies `json:"rates"`
}

type Currencies struct {
	Vs_currencies float64 `json:"vs_currencies"`
}
type PriceResponse struct {
	Ids Currencies
}

type Error struct {
	//Error code
	Code int

	//Error message
	Message string
}

func writeError(writer http.ResponseWriter, message string, code int) {
	response := Error{
		Code:    code,
		Message: message,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)

	json.NewEncoder(writer).Encode(response)
}

var (
	RequestErrorHandler = func(writer http.ResponseWriter, err error) {
		writeError(writer, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(writer http.ResponseWriter) {
		writeError(writer, "An Unexpected Error Ocurred.", http.StatusInternalServerError)
	}
)
