package s_bookmark_tag

import (
	"fmt"
	"my-project/src/common/com_msg"
	"my-project/src/config/conf_sql"
	"my-project/src/model"
	"strings"
)

const table string = `bookmark_tag`

func Page(dto model.BookmarkDto) (bookmarkTags []model.BookmarkTag, total int64, err error) {
	sqlite := conf_sql.InitSqlite()
	err = sqlite.Task(func() error {
		w := sqlite.NewWhere()
		if len(dto.Keyword) > 0 {
			w.AndLike(`name`, dto.Keyword)
		}

		total, err = sqlite.Count(table, *w)
		if err != nil {
			return err
		}

		if total > 0 {
			w.Desc(`sort`).Desc(`rowid`).Limit(dto.Page, dto.PageSize)
			err = sqlite.FindSlice(table, *w, &bookmarkTags, `*`, `rowid`)
		}
		return err
	})
	return
}

func All() (tags []model.BookmarkTag, err error) {
	sqlite := conf_sql.InitSqlite()
	w := sqlite.NewWhere().Desc(`sort`).Desc(`rowid`)
	err = sqlite.FindSlice(table, *w, &tags, `rowid`, `*`)
	return
}

func ById(id int64) (tag model.BookmarkTag, err error) {
	sqlite := conf_sql.InitSqlite()
	w := sqlite.NewWhere().AndEq(`rowid`, id)
	sqlite.FindOne(table, *w, &tag, `rowid`, `*`)
	return
}

func Insert(data model.BookmarkTag) (msg string, ok bool) {
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
	sort := data.Sort
	if sort < 0 {
		msg = com_msg.PositiveInteger(`Sort`)
		return
	}

	sqlite := conf_sql.InitSqlite()
	w1 := sqlite.NewWhere().AndNq(`tag`, tag)
	count1, err := sqlite.Count(table, *w1)
	if err != nil {
		msg = fmt.Sprintf(`Save bookmark tag error: %v`, err)
		return
	}
	if count1 > 0 {
		msg = `The tag already exists and cannot be added.`
		return
	}

	insertSet := sqlite.NewColumn().Set(`name`, name).Set(`tag`, tag).Set(`sort`, sort)
	_, err = sqlite.Insert(table, *insertSet)
	if err != nil {
		msg = fmt.Sprintf(`Save bookmark tag error: %v`, err)
		return
	}
	ok = true
	return
}

func Update(data model.BookmarkTag) (msg string, ok bool) {
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
	sort := data.Sort
	if sort < 0 {
		msg = com_msg.PositiveInteger(`Sort`)
		return
	}

	sqlite := conf_sql.InitSqlite()
	w1 := sqlite.NewWhere().AndEq(`tag`, tag).AndNq(`rowid`, id)
	count1, err := sqlite.Count(table, *w1)
	if err != nil {
		msg = fmt.Sprintf(`Update bookmark tag error: %v`, err)
		return
	}
	if count1 > 0 {
		msg = com_msg.ExistsNotUpdate(`Tag`)
		return
	}

	w2 := sqlite.NewWhere().AndEq(`tag`, tag)
	count2, err := sqlite.Count(`bookmark`, *w2)
	if err != nil {
		msg = fmt.Sprintf(`Update bookmark tag error: %v`, err)
		return
	}
	if count2 > 0 {
		msg = com_msg.HasBeenUsedNotUpdate(`Tag`)
		return
	}

	updateSet3 := sqlite.NewColumn().Set(`name`, name).Set(`tag`, tag).Set(`sort`, sort)
	w3 := sqlite.NewWhere().AndEq(`rowid`, id)
	_, err = sqlite.Update(table, *updateSet3, *w3)
	if err != nil {
		msg = fmt.Sprintf(`Update bookmark tag error: %v`, err)
		return
	}

	ok = true
	return
}

func Delete(id int64) (msg string, ok bool) {
	sqlite := conf_sql.InitSqlite()

	count1, err := sqlite.CountSql(`select count(*) as c from bookmark where tag = (select tag from bookmark_tag where rowid = ?)`, id)
	if err != nil {
		msg = fmt.Sprintf(`Delete bookmark tag error: %v`, err)
		return
	}
	if count1 > 0 {
		msg = com_msg.HAS_BEEN_USED_NOT_DEL
		return
	}

	w2 := sqlite.NewWhere().AndEq(`rowid`, id)
	_, err = sqlite.Delete(table, *w2)
	if err != nil {
		msg = fmt.Sprintf(`Delete bookmark tag error: %v`, err)
		return
	}

	ok = true
	return
}
