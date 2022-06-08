package init_db

import (
	"my-project/src/config/conf_sql"
	"my-project/src/model"
)

func insertSearchEngine() error {
	sqlite := conf_sql.InitSqlite()

	return sqlite.Task(func() error {
		count, err := sqlite.Count(`search_engine`, *sqlite.NewWhere())
		if err != nil {
			return err
		}
		if count > 0 {
			return nil
		}

		columns := []string{`name`, `url`, `search`, `sort`}
		data := make([]model.SearchEngine, 0)
		// Bing
		data = append(data, model.SearchEngine{Name: `Bing`, Url: `https://www.bing.com`, Search: `/search?q=`})
		values := make([]interface{}, 0)
		for index, se := range data {
			values = append(values, se.Name, se.Url, se.Search, index)
		}
		_, err = sqlite.InsertMore(`search_engine`, columns, len(data), values...)
		if err != nil {
			return err
		}
		return nil
	})
}
