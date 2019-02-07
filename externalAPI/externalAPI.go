package externalAPI

import (
	"fmt"
)

const (
	endPoint           = ""
	rateLimitPerSec    = 30
	retryLimit         = 3
	ErrExceedRateLimit = fmt.Errorf("request exceed rate limit")
)

const (
	apiEndPoint  = "http://data.coa.gov.tw/Service/OpenData/TransService.aspx"
	defaultQuery = "?UnitId=QcbUEzN6E6DL&animal_kind=貓&animal_colour="
)

func queryCatColor(color string) (string, error) {
	return "", fmt.Errorf("not implement yet")
}
