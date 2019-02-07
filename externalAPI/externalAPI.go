package externalAPI

import (
	"fmt"
)

const (
	rateLimitPerSec = 30
	retryLimit      = 3
	apiEndPoint     = "http://data.coa.gov.tw/Service/OpenData/TransService.aspx"
	defaultQuery    = "?UnitId=QcbUEzN6E6DL&animal_kind=貓&animal_colour="
)

var (
	ErrExceedRateLimit = fmt.Errorf("request exceed rate limit")
)

func QueryCatColor(color string) (string, error) {
	return "", fmt.Errorf("not implement yet")
}
