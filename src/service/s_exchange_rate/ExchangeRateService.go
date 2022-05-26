package s_exchange_rate

import (
	"encoding/json"
	"fmt"
	"my-project/src/api/api_exrate"
	"my-project/src/config/conf_sql"
	"my-project/src/utils/u_math"
	"my-project/src/utils/u_string"
	"my-project/src/utils/u_time"
	"strings"
)

type code struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Index int    `json:"sort"`
}

type rate struct {
	DateString string `json:"data_string"`
	DateUnix   int64  `json:"date_unix"`
	Code       string `json:"code"`
	RateString string `json:"rates"`
}

func Codes() ([]code, error) {
	sqlite := conf_sql.InitSqlite()
	codes := make([]code, 0)

	w := sqlite.NewWHere().Asc(`sort`)
	err := sqlite.FindSlice(`exrate_code`, *w, &codes)
	return codes, err
}

func RatesData() ([]rate, error) {
	sqlite := conf_sql.InitSqlite()
	rates := make([]rate, 0)

	w := sqlite.NewWHere().Desc(`date_unix`)
	err := sqlite.FindSlice(`exrate_rate`, *w, &rates, `date_unix`, `code`)
	if err != nil {
		return rates, err
	}
	rates1 := make([]rate, 0)
	for _, val := range rates {
		rates1 = append(rates1, rate{Code: val.Code, DateString: u_time.FmtYmdUnix(val.DateUnix)})
	}
	return rates1, nil
}

func PullAndSaveRates(currencyCode string) (msg string) {
	code := strings.ToUpper(currencyCode)

	sqlite := conf_sql.InitSqlite()
	w := sqlite.NewWHere().AndEq(`code`, code)
	codeCount, err := sqlite.Count(`exrate_code`, *w)
	if err != nil {
		msg = fmt.Sprintf(`Update data error: %v`, err)
		return
	}
	if codeCount == 0 {
		msg = fmt.Sprintf(`%s is not supported.`, code)
		return
	}
	rateData, err := api_exrate.PullExchangeRates(code)
	if err != nil {
		msg = fmt.Sprintf(`Update data error: %v`, err)
		return
	}

	err = sqlite.Task(func() error {
		whereRate := sqlite.NewWHere().AndEq(`code`, code)
		rateCount, err := sqlite.Count(`exrate_rate`, *whereRate)
		if err != nil {
			return err
		}

		baseCode := rateData.BaseCode
		dateUnix := rateData.TimeLastUpdateUnix
		jsonBytes, _ := json.Marshal(rateData.ConversionRates)
		if rateCount > 0 {
			// Update
			updateSet := sqlite.NewColumn().Set(`date_unix`, dateUnix).Set(`rates`, jsonBytes)
			updateWhere := sqlite.NewWHere().AndEq(`code`, baseCode)
			_, err = sqlite.Update(`exrate_rate`, *updateSet, *updateWhere)
		} else {
			// Insert
			insetSet := sqlite.NewColumn().Set(`date_unix`, dateUnix).Set(`code`, baseCode).Set(`rates`, jsonBytes)
			_, err = sqlite.Insert(`exrate_rate`, *insetSet)
		}
		return err
	})
	if err != nil {
		msg = fmt.Sprintf(`Update data error: %v`, err)
		return
	}
	return
}

func Exchange(fromCode, toCode string, amount string) (msg string, ok bool) {
	amountF64, err := u_string.Float64(amount)
	if err != nil || amountF64 < 0 {
		msg = fmt.Sprintf(`Amount must be greater than %v.`, 0)
		return
	}

	codeF := strings.ToUpper(fromCode)
	codeT := strings.ToUpper(toCode)

	rate := rate{}
	var defaultRate bool = false

	sqlite := conf_sql.InitSqlite()
	err = sqlite.Task(func() (err error) {
		// Get rates
		w := sqlite.NewWHere().AndEq(`code`, codeF)
		err = sqlite.FindOne(`exrate_rate`, *w, &rate)
		fmt.Println(`==========`)
		fmt.Println(err)
		fmt.Println(`==========`)
		if err != nil {
			return
		}
		if len(rate.RateString) == 0 {
			defaultRate = true
			w1 := sqlite.NewWHere().AndEq(`code`, `USD`)
			err = sqlite.FindOne(`exrate_rate`, *w1, &rate)
		}
		return
	})
	if err != nil {
		msg = fmt.Sprintf(`Convert error: %v`, err)
		return
	}
	rateMap := make(map[string]float64)
	json.Unmarshal([]byte(rate.RateString), &rateMap)
	if len(rateMap) == 0 {
		msg = `Please update the data.`
		return
	}
	if rateMap[codeF] <= 0 {
		msg = fmt.Sprintf(`%s is not supported.`, codeF)
		return
	}
	if rateMap[codeT] <= 0 {
		msg = fmt.Sprintf(`%s is not supported.`, codeT)
		return
	}

	var amountResult float64
	if defaultRate {
		// The source currency code is in the configuration
		// Example: [JPY] -> [KRW]
		amountResult = u_math.Multiply4F(amountF64, rateMap[codeT])
	} else {
		// Use default currency code, convert amount using USD exchange rate
		if codeF == `USD` {
			// Example: [USD] -> [JPY]
			amountResult = u_math.Multiply4F(amountF64, rateMap[codeT])
		} else {
			if codeT == `USD` {
				// Example: [KRW] -> [USD]
				amountResult = u_math.Divide4F(amountF64, rateMap[codeF])
			} else {
				// Example: [SGD] -> [USD] -> [AUD]
				amountResult = u_math.Divide4F(amountF64, rateMap[codeF])
				amountResult = u_math.Multiply4F(amountResult, rateMap[codeT])
			}
		}
	}
	msg = fmt.Sprintf(`%v %s = %v %s.`, amount, codeF, amountResult, codeT)
	ok = true
	return
}

func DeleteRateByCode(currencyCode string) error {
	code := strings.ToUpper(currencyCode)
	sqlite := conf_sql.InitSqlite()
	if len(code) == 0 || code == `USD` {
		return nil
	}

	w := sqlite.NewWHere().AndEq(`code`, code)
	_, err := sqlite.Delete(`exrate_rate`, *w)
	return err
}
