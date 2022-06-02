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
			w.AndLike(`tag`, dto.Tag)
		}

		total, err = sqlite.Count(table, *w)
		if err != nil {
			return err
		}

		if total > 0 {
			w.Desc(`sort`).Desc(`rowid`).Limit(dto.Page, dto.PageSize)
			err = sqlite.FindSlice(table, *w, &bookmarks)
		}
		return err
	})
	return
}

func ById(id int64) (bookmark model.Bookmark, err error) {
	sqlite := conf_sql.InitSqlite()
	w := sqlite.NewWhere().AndEq(`rowid`, id)
	err = sqlite.FindOne(table, *w, &bookmark)
	return
}

func Insert(data model.Bookmark) (msg string, ok bool) {
	if data.Id != 0 {
		msg = com_msg.ADD_FAIL
	}
	name := strings.TrimSpace(data.Name)
	if len(name) == 0 {
		msg = com_msg.Required(`Name`)
		return
	}
	tag := strings.TrimSpace(data.Tag)
	if len(tag) == 0 {
		msg = com_msg.Required(`Tag`)
		return
	}
	link := strings.TrimSpace(data.Link)
	if len(link) == 0 {
		msg = com_msg.Required(`Link`)
		return
	}
	sort := data.Sort
	if sort < 0 {
		msg = com_msg.PositiveInteger(`Sort`)
		return
	}

	sqlite := conf_sql.InitSqlite()
	insertSet := sqlite.NewColumn().Set(`name`, name).Set(`tag`, tag).Set(`link`, link).Set(`sort`, sort)
	_, err := sqlite.Insert(tag, *insertSet)
	if err != nil {
		msg = fmt.Sprintf(`Save bookmark error: %v`, err)
		return
	}

	ok = true
	return
}

func Update(data model.Bookmark) (msg string, ok bool) {
	id := data.Id
	if id <= 0 {
		return com_msg.DEL_SUCCESS, true
	}
	name := strings.TrimSpace(data.Name)
	if len(name) == 0 {
		msg = com_msg.Required(`Name`)
		return
	}
	tag := strings.TrimSpace(data.Tag)
	if len(tag) == 0 {
		msg = com_msg.Required(`Tag`)
		return
	}
	link := strings.TrimSpace(data.Link)
	if len(link) == 0 {
		msg = com_msg.Required(`Link`)
		return
	}
	sort := data.Sort
	if sort < 0 {
		msg = com_msg.PositiveInteger(`Sort`)
		return
	}

	sqlite := conf_sql.InitSqlite()
	updateSet := sqlite.NewColumn().Set(`name`, name).Set(`tag`, tag).Set(`link`, link).Set(`sort`, sort)
	w := sqlite.NewWhere().AndEq(`rowid`, id)
	_, err := sqlite.Update(table, *updateSet, *w)
	if err != nil {
		msg = fmt.Sprintf(`Update bookmark error: %v`, err)
		return
	}

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

	ok = true
	return
}

func ByTag(tag string) (bookmarks []model.Bookmark, err error) {
	sqlite := conf_sql.InitSqlite()
	w := sqlite.NewWhere().AndEq(`tag`, tag)
	err = sqlite.FindSlice(table, *w, &bookmarks, `*`, `rowid`)
	return
}
