package s_bookmark

import (
	"fmt"
	"my-project/src/common/com_msg"
	"my-project/src/config/conf_sql"
	"my-project/src/model"
	"strings"
)

const table string = `bookmark`

func Page(dto model.BookmarkDto) (bookmarks []model.Bookmark, total int64, err error) {
	sqlite := conf_sql.InitSqlite()
	err = sqlite.Task(func() error {
		w := sqlite.NewWhere()
		if len(dto.Keyword) > 0 {
			w.AndLike(`name`, dto.Keyword)
		}
		if len(dto.Tag) > 0 {
			w.AndEq(`tag`, dto.Tag)
		}

		total, err = sqlite.Count(table, *w)
		if err != nil {
			return err
		}

		fmt.Println(dto)
		if total > 0 {
			w.Desc(`sort`).Asc(`rowid`).Limit(dto.Page, dto.PageSize).Limit(dto.Page, dto.PageSize)
			err = sqlite.FindSlice(table, *w, &bookmarks, `rowid`, `*`)
		}
		return err
	})
	return
}

func ById(id int64) (bookmark model.Bookmark, err error) {
	sqlite := conf_sql.InitSqlite()
	w := sqlite.NewWhere().AndEq(`rowid`, id)
	err = sqlite.FindOne(table, *w, &bookmark, `rowid`, `*`)
	return
}

func Save(data *model.Bookmark) (msg string, ok bool) {
	data.Name = strings.TrimSpace(data.Name)
	if len(data.Name) == 0 {
		msg = com_msg.Required(`Name`)
		return
	}
	data.Tag = strings.TrimSpace(data.Tag)
	if len(data.Tag) == 0 {
		msg = com_msg.Required(`Tag`)
		return
	}
	data.Link = strings.TrimSpace(data.Link)
	if len(data.Link) == 0 {
		msg = com_msg.Required(`Link`)
		return
	}
	data.Description = strings.TrimSpace(data.Description)
	sort := data.Sort
	if sort < 0 || sort > 99 {
		msg = com_msg.Range(`Sort`, `0`, `99`)
		return
	}

	if data.Id == 0 {
		// insert
		return insert(*data)
	} else {
		// update
		return update(*data)
	}
}

func insert(data model.Bookmark) (msg string, ok bool) {
	sqlite := conf_sql.InitSqlite()
	insertSet := sqlite.NewColumn().Set(`name`, data.Name).Set(`tag`, data.Tag).Set(`link`, data.Link).Set(`description`, data.Description).Set(`sort`, data.Sort)
	_, err := sqlite.Insert(table, *insertSet)
	if err != nil {
		msg = fmt.Sprintf(`Save bookmark error: %v`, err)
		return
	}

	msg = com_msg.ADD_SUCCESS
	ok = true
	return
}

func update(data model.Bookmark) (msg string, ok bool) {
	sqlite := conf_sql.InitSqlite()
	updateSet := sqlite.NewColumn().Set(`name`, data.Name).Set(`tag`, data.Tag).Set(`link`, data.Link).Set(`description`, data.Description).Set(`sort`, data.Sort)
	w := sqlite.NewWhere().AndEq(`rowid`, data.Id)
	_, err := sqlite.Update(table, *updateSet, *w)
	if err != nil {
		msg = fmt.Sprintf(`Update bookmark error: %v`, err)
		return
	}

	msg = com_msg.UPD_SUCCESS
	ok = true
	return
}

func Delete(id int64) (msg string, ok bool) {
	sqlite := conf_sql.InitSqlite()
	w := sqlite.NewWhere().AndEq(`rowid`, id)
	_, err := sqlite.Delete(table, *w)
	if err != nil {
		msg = fmt.Sprintf(`Delete bookmark error: %v`, err)
		return
	}

	msg = com_msg.DEL_SUCCESS
	ok = true
	return
}

func ByTag(tag string) (bookmarks []model.Bookmark, err error) {
	sqlite := conf_sql.InitSqlite()
	w := sqlite.NewWhere().AndEq(`tag`, tag)
	err = sqlite.FindSlice(table, *w, &bookmarks, `*`, `rowid`)
	return
}
