package api_exrate

// ExchangeRate-API
// https://www.exchangerate-api.com/
import (
	"encoding/json"
	"errors"
	"fmt"
	"my-project/src/utils/u_http"
)

type Codes struct {
	Result         string     `json:"result"`
	ErrorType      string     `json:"error-type"`
	Documentation  string     `json:"documentation"`
	TermsOfUse     string     `json:"terms_of_use"`
	SupportedCodes [][]string `json:"supported_codes"`
}

type Rates struct {
	Result             string             `json:"result"`
	ErrorType          string             `json:"error-type"`
	Documentation      string             `json:"documentation"`
	TermsOfUse         string             `json:"terms_of_use"`
	TimeLastUpdateUnix int64              `json:"time_last_update_unix"`
	TimeLastUpdateUtc  string             `json:"time_last_update_utc"`
	TimeNextUpdateUnix int64              `json:"time_next_update_unix"`
	TimeNextUpdateUtc  string             `json:"time_next_update_utc"`
	BaseCode           string             `json:"base_code"`
	ConversionRates    map[string]float64 `json:"conversion_rates"`
}

type Quota struct {
	Result            string `json:"result"`
	ErrorType         string `json:"error-type"`
	Documentation     string `json:"documentation"`
	TermsOfUse        string `json:"terms_of_use"`
	PlanQuota         int64  `json:"plan_quota"`
	RequestsRemaining int64  `json:"requests_remaining"`
	RefreshDayOfMonth int64  `json:"refresh_day_of_month"`
}

// Api key is required
const ApiKey string = ``

// https://www.exchangerate-api.com/docs/supported-codes-endpoint
func PullCodes() (result Codes, err error) {
	if len(ApiKey) == 0 {
		err = errors.New(`api key must be required`)
		return
	}

	url := fmt.Sprintf(`https://v6.exchangerate-api.com/v6/%s/codes`, ApiKey)
	body, err := u_http.Get(url, nil, nil)
	if err != nil {
		return
	}

	json.Unmarshal(body, &result)
	if result.Result != `success` {
		err = errors.New(result.ErrorType)
	}
	return
}

// https://www.exchangerate-api.com/docs/standard-requests
func PullExchangeRates(currencyCode string) (result Rates, err error) {
	if len(ApiKey) == 0 {
		err = errors.New(`api key must be required`)
		return
	}
	if len(currencyCode) == 0 {
		err = errors.New(`currency code must be required`)
		return
	}

	url := fmt.Sprintf(`https://v6.exchangerate-api.com/v6/%s/latest/%s`, ApiKey, currencyCode)
	body, err := u_http.Get(url, nil, nil)
	if err != nil {
		return
	}

	json.Unmarshal(body, &result)
	if result.Result != `success` {
		err = errors.New(result.ErrorType)
	}
	return
}

// https://www.exchangerate-api.com/docs/request-quota-endpoint
func GetQuota() (result Quota, err error) {
	if len(ApiKey) == 0 {
		err = errors.New(`api key must be required`)
		return
	}

	url := fmt.Sprintf(`https://v6.exchangerate-api.com/v6/%s/quota`, ApiKey)
	body, err := u_http.Get(url, nil, nil)
	if err != nil {
		return
	}

	json.Unmarshal(body, &result)
	if result.Result != `success` {
		err = errors.New(result.ErrorType)
	}
	return
}
