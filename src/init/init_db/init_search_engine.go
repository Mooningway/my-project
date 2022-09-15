package init_db

import (
	"my-project/src/config/conf_sql"
	"my-project/src/model"
)

func insertSearchEngine() error {
	sqlite := conf_sql.InitSqlite()

	return sqlite.Task(func() error {
		count, err := sqlite.Count(`search_engine`, *sqlite.NewQuery())
		if err != nil {
			return err
		}
		if count > 0 {
			return nil
		}

		searchEngines := make([]model.SearchEngine, 0)
		// Bing
		searchEngines = append(searchEngines, model.SearchEngine{Name: `Bing`, Url: `https://www.bing.com`, Search: `/search?q=`})
		_, err = sqlite.InsertMore(`search_engine`, searchEngines)
		if err != nil {
			return err
		}
		return nil
	})
}
