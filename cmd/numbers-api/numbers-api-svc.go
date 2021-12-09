package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://numbersapi.com/13/trivia?fragment"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("Code:[%d] Body:[%s]", res.StatusCode, string(body))
}
