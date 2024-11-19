package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Go-Go-Go/api"
)

const ROOT string = "https://api.coingecko.com/api/v3"
const HEADER_AUTH_KEY string = "x_cg_demo_api_key"
const API_KEY string = "CG-4YbsrgnB5a45UdF9gYyqXJfi"

func CGRequest(url string, params url.Values, response any) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s?%s", ROOT+url, params.Encode()), nil)
	req.Header.Add("accept", "application/json") //Averiguar bien porque es necesario este header
	req.Header.Add(HEADER_AUTH_KEY, API_KEY)     //Averiguar la forma mas segura de pasar la api key

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close() //Averiguar bien que es defer para documentar como caracteristica del lenguaje que creo que hayu algo ahi
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, response)
}

func FilterCoinIDs(nameFilter string) string {
	//Returns a comma separated list of all the coin ids wich names contain the filter.
	var coinIDsMap []api.CoinListResponse
	var filteredIDs []string
	coinLimit := 5 //Pensar bien el coin limit como lo hago

	CGRequest("/coins/list", url.Values{}, &coinIDsMap) //Ver como no hacer esto siempre y solo hacerlo cada 5 min.

	for _, coin := range coinIDsMap {
		if strings.Contains(strings.ToLower(coin.Name), strings.ToLower(nameFilter)) && len(filteredIDs) <= coinLimit {
			filteredIDs = append(filteredIDs, coin.ID)
		}
	}
	return strings.Join(filteredIDs, ",")
}
