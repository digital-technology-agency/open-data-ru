package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.coinranking.com/v2/coin/Qwsogvtv82FCd/history?timePeriod=30d"
	//url := "https://api.coinranking.com/v2/coin/Qwsogvtv82FCd/history"
	//url := "https://api.coinranking.com/v2/exchanges"
	//url := "https://api.coinranking.com/v2/markets"
	//url := "https://api.coinranking.com/v2/stats"
	//url := "https://api.coinranking.com/v2/coins?timePeriod=3h"
	//url := "https://api.coinranking.com/v2/coins?search=bitcoin"
	//url := "https://api.coinranking.com/v2/coins?referenceCurrencyUuid=Qwsogvtv82FCd"
	//url := "https://api.coinranking.com/v2/coins?symbols[]=BTC&symbols[]=ETH&symbols=XRP"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("Code:[%d] Body:[%s]", res.StatusCode, string(body))
}
