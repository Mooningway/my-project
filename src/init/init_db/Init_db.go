package init_db

// Project initialization
import (
	"errors"
	"fmt"
	"my-project/src/config/conf_sql"
)

func Init() error {
	// Create Table
	err := createTable()
	if err != nil {
		out := fmt.Sprintf(`Create Table error: %v`, err)
		return errors.New(out)
	}
	// Insert Data
	err = insertData()
	if err != nil {
		out := fmt.Sprintf(`Insert Data error: %v`, err)
		return errors.New(out)
	}
	return nil
}

func createTable() error {
	sqlite := conf_sql.InitSqlite()
	return sqlite.Task(func() (err error) {
		// exrate_code
		_, err = sqlite.DB.Exec(sql_create_exrate_code)
		if err != nil {
			return
		}
		// exrate_rate
		_, err = sqlite.DB.Exec(sql_create_exrate_rate)
		if err != nil {
			return
		}
		// bookmark
		_, err = sqlite.DB.Exec(sql_create_bookmark)
		if err != nil {
			return
		}
		// bookmark_tag
		_, err = sqlite.DB.Exec(sql_create_bookmark_tag)
		if err != nil {
			return
		}
		// search_engine
		_, err = sqlite.DB.Exec(sql_create_search_engine)
		if err != nil {
			return
		}
		return
	})
}

func insertData() error {
	// exrate_code
	err := insertExrateCode()
	if err != nil {
		return err
	}
	// exrate_rate
	err = insertExrateRate()
	if err != nil {
		return err
	}
	// search_engine
	err = insertSearchEngine()
	if err != nil {
		return err
	}
	return nil
}
