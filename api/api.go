package api

import (
	"encoding/json"
	"net/http"
)

const JWT_EXPIRATION_MINUTES = 60

// IndexTemplateData specifies data passed to index.html template.
type IndexTemplateData struct {
	Title string
	Coins []CoinBasicInformation
}

// LoginTemplateData specifies data passed to login.html template
type LoginTemplateData struct {
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

// CoinsListResponse specifies the structure of the "/list" CG endpoint.
type CoinListResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CoinTemplateData specifies data passed to coin.html template.
type CoinTemplateData struct {
	Name         string
	CurrentPrice float64
	MaxPrice     float64
	MinPrice     float64
	Logo         string
	History      []HistoryTemplateData
}

// HistoryTemplateData specifies data passed to history coin prices table on coin.html template.
type HistoryTemplateData struct {
	Time         string
	Price        float64
	MarketCap    float64
	TotalVolumes float64
}

// CoinDataResponse specifies the structure of the "/coins/{id}" CG endpoint.
type CoinDataResponse struct {
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
	MarketData struct {
		CurrentPrice struct {
			USD float64 `json:"usd"`
		} `json:"current_price"`
		AllTimeHigh struct {
			USD float64 `json:"usd"`
		} `json:"ath"`
		AllTimeLow struct {
			USD float64 `json:"usd"`
		} `json:"atl"`
		MarketCap struct {
			USD float64 `json:"usd"`
		} `json:"market_cap"`
		Rank int `json:"market_cap_rank"`
	} `json:"market_data"`
}

// CoinHistoryResponse specifies the structure of the "/coins/{id}/market_chart/range" CG endpoint.
type CoinHistoryResponse struct {
	Prices       [][]float64 `json:"prices"`
	MarketCaps   [][]float64 `json:"market_caps"`
	TotalVolumes [][]float64 `json:"total_volumes"`
}

const TIME int = 0
const VALUE int = 1

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
