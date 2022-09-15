package init_db

import (
	"encoding/json"
	"my-project/src/api/api_exrate"
	"my-project/src/config/conf_sql"
	"my-project/src/model"
)

func insertExrateCode() error {
	sqlite := conf_sql.InitSqlite()

	return sqlite.Task(func() error {
		count, err := sqlite.Count(`exrate_code`, *sqlite.NewQuery())
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

		codes := codeData.SupportedCodes
		exrateCodes := make([]model.ExrateCode, 0)
		for index, vals := range codes {
			exrateCodes = append(exrateCodes, model.ExrateCode{Code: vals[0], Name: vals[1], Sort: index})
		}
		_, err = sqlite.InsertMore(`exrate_code`, exrateCodes)
		if err != nil {
			return err
		}
		return nil
	})
}

func insertExrateRate() error {
	sqlite := conf_sql.InitSqlite()

	return sqlite.Task(func() error {
		count, err := sqlite.Count(`exrate_rate`, *sqlite.NewQuery())
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
		update := sqlite.NewUpdate().Set(`date_unix`, rateData.TimeLastUpdateUnix).Set(`code`, rateData.BaseCode).Set(`rates`, jsonBytes)
		_, err = sqlite.InsertByUpdate(`exrate_rate`, *update)
		if err != nil {
			return err
		}
		return nil
	})
}
