package init_db

import (
	"encoding/json"
	"my-project/src/api/api_exrate"
	"my-project/src/config/conf_sql"
)

func insertExrateCode() error {
	sqlite := conf_sql.InitSqlite()

	return sqlite.Task(func() error {
		count, err := sqlite.Count(`exrate_code`, *sqlite.NewWhere())
		if err != nil {
			return err
		}
		if count > 0 {
			return nil
		}
		codeData, err := api_exrate.PullCodes()
		if err != nil {
			return err
		}

		columns := []string{`code`, `name`, `sort`}
		codes := codeData.SupportedCodes
		values := make([]interface{}, 0)
		for index, vals := range codes {
			values = append(values, vals[0], vals[1], index)
		}
		_, err = sqlite.InsertMore(`exrate_code`, columns, len(codes), values...)
		if err != nil {
			return err
		}
		return nil
	})
}

func insertExrateRate() error {
	sqlite := conf_sql.InitSqlite()

	return sqlite.Task(func() error {
		count, err := sqlite.Count(`exrate_rate`, *sqlite.NewWhere())
		if err != nil {
			return err
		}
		if count > 0 {
			return nil
		}
		rateData, err := api_exrate.PullExchangeRates(`USD`)
		if err != nil {
			return err
		}

		jsonBytes, _ := json.Marshal(rateData.ConversionRates)

		insetSet := sqlite.NewColumn().Set(`date_unix`, rateData.TimeLastUpdateUnix).Set(`code`, rateData.BaseCode).Set(`rates`, jsonBytes)
		_, err = sqlite.Insert(`exrate_rate`, *insetSet)
		if err != nil {
			return err
		}
		return nil
	})
}
