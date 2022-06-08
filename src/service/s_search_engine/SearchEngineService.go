package s_search_engine

import (
	"my-project/src/config/conf_sql"
	"my-project/src/model"
)

const table string = `search_engine`

func All() (searchEngines []model.SearchEngine, err error) {
	sqlite := conf_sql.InitSqlite()
	err = sqlite.FindSlice(table, *sqlite.NewWhere(), &searchEngines)
	return
}
