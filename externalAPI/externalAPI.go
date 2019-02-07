package externalAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	rateLimitPerSec = 30
	retryLimit      = 3
	apiEndPoint     = "http://data.coa.gov.tw/Service/OpenData/TransService.aspx"
	defaultQuery    = "?UnitId=QcbUEzN6E6DL&animal_kind=貓&animal_colour="
)

var (
	ErrExceedRateLimit = fmt.Errorf("request exceed rate limit")
	QueryMsgs          = make(chan string, 100)
	RateLimit          = time.Tick(time.Second / rateLimitPerSec)
)

func QueryCatByColor(color string) (string, error) {
	req, err := http.NewRequest("GET", apiEndPoint+defaultQuery+color, nil)
	if err != nil {
		fmt.Println("QueryCatcolor err :%v", err)
		return "", nil
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("http req err :%v", err)
		return "", nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("external request fail , status code ", resp.Status)
		return "", nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("readbody err :%v", err)
		return "", nil
	}

	result := []map[string]interface{}{}
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		fmt.Printf("external API response json unmarshal fail%v\n", err)
		return "", nil
	}

	b, err := json.Marshal(result)
	if err != nil {
		return "", nil
	}

	return string(b), nil
}
