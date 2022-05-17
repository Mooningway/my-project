package init_db

import (
	"encoding/json"
	"fmt"
	"my-project/src/api/api_exrate"
	"my-project/src/config/conf_sql"
	"strings"
)

func insertExrateCode() error {
	sqlite := conf_sql.InitSqlite()

	return sqlite.Exec(func() error {
		count, err := sqlite.CountExec(`SELECT COUNT(*) c FROM exrate_code`)
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

		values := make([]interface{}, 0)
		insertValues := make([]string, 0)
		for index, vals := range codeData.SupportedCodes {
			insertValues = append(insertValues, `(?, ?, ?)`)
			values = append(values, vals[0], vals[1], index)
		}
		insertSql := fmt.Sprintf(`INSERT INTO exrate_code (code, name, sort) VALUES %s`, strings.Join(insertValues, `,`))
		_, err = sqlite.Db.Exec(insertSql, values...)
		if err != nil {
			return err
		}
		return nil
	})
}

func insertExrateRate() error {
	sqlite := conf_sql.InitSqlite()

	return sqlite.Exec(func() error {
		count, err := sqlite.CountExec(`SELECT COUNT(*) c FROM exrate_rate`)
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

		baseCode := rateData.BaseCode
		dateUnix := rateData.TimeLastUpdateUnix
		jsonBytes, _ := json.Marshal(rateData.ConversionRates)
		_, err = sqlite.Db.Exec(`INSERT INTO exrate_rate (date_unix, code, rates) VALUES (?, ?, ?)`, dateUnix, baseCode, jsonBytes)
		if err != nil {
			return err
		}
		return nil
	})
}
