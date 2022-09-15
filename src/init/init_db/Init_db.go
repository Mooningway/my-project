package init_db

// Project initialization
import (
	"fmt"
	"log"
	"my-project/src/config/conf_sql"
	"my-project/src/model"
)

func Init() bool {
	// Create Table
	errors := createTable()
	if len(errors) > 0 {
		for _, err := range errors {
			out := fmt.Sprintf(`Create Table error: %v`, err)
			log.Println(out)
		}
		return false
	}

	// Insert Data
	errors = insertData()
	if len(errors) > 0 {
		for _, err := range errors {
			out := fmt.Sprintf(`Insert Data error: %v`, err)
			log.Println(out)
		}
		return false
	}

	return true
}

func createTable() []error {
	tabMap := make(map[string]interface{})
	tabMap[TABLE_ENGINE] = model.SearchEngine{}
	tabMap[TABLE_BOOKMARK] = model.Bookmark{}
	tabMap[TABLE_BOOKMARK_TAG] = model.BookmarkTag{}
	tabMap[TABLE_EXRATE_CODE] = model.ExrateCode{}
	tabMap[TABLE_EXRATE_RATE] = model.ExrateRate{}
	tabMap[TABLE_NOTE_FOLDER] = model.NoteFolder{}
	tabMap[TABLE_NOTE] = model.Note{}

	sqlite := conf_sql.InitSqlite()
	return sqlite.CreateTables(tabMap)
}

func insertData() []error {
	tempErrors := make([]error, 0)
	tempErrors = append(tempErrors, insertExrateCode())   // exrate_code
	tempErrors = append(tempErrors, insertExrateRate())   // exrate_rate
	tempErrors = append(tempErrors, insertSearchEngine()) // search_engine
	errors := make([]error, 0)
	for _, err := range tempErrors {
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
