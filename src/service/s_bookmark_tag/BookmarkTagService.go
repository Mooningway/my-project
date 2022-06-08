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
			w.Desc(`sort`).Asc(`rowid`).Limit(dto.Page, dto.PageSize)
			err = sqlite.FindSlice(table, *w, &bookmarkTags, `*`, `rowid`)
		}
		return err
	})
	return
}

func All() (tags []model.BookmarkTag, err error) {
	sqlite := conf_sql.InitSqlite()
	w := sqlite.NewWhere().Desc(`sort`).Asc(`rowid`)
	err = sqlite.FindSlice(table, *w, &tags, `rowid`, `*`)
	return
}

func ById(id int64) (tag model.BookmarkTag, err error) {
	sqlite := conf_sql.InitSqlite()
	w := sqlite.NewWhere().AndEq(`rowid`, id)
	sqlite.FindOne(table, *w, &tag, `rowid`, `*`)
	return
}

func Save(data *model.BookmarkTag) (msg string, ok bool) {
	if data.Id != 0 {
		msg = com_msg.ADD_FAIL
	}
	data.Name = strings.TrimSpace(data.Name)
	if len(data.Name) == 0 {
		msg = com_msg.Required(`Name`)
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

func insert(data model.BookmarkTag) (msg string, ok bool) {
	sqlite := conf_sql.InitSqlite()

	w := sqlite.NewWhere().AndEq(`name`, data.Name)
	count, err := sqlite.Count(table, *w)
	if err != nil {
		msg = fmt.Sprintf(`Save bookmark tag error: %v`, err)
		return
	}
	if count > 0 {
		msg = com_msg.ExistsNotAdd(`name`)
		return
	}

	insertSet := sqlite.NewColumn().Set(`name`, data.Name).Set(`description`, data.Description).Set(`sort`, data.Sort)
	_, err = sqlite.Insert(table, *insertSet)
	if err != nil {
		msg = fmt.Sprintf(`Save bookmark tag error: %v`, err)
		return
	}

	msg = com_msg.ADD_SUCCESS
	ok = true
	return
}

func update(data model.BookmarkTag) (msg string, ok bool) {
	sqlite := conf_sql.InitSqlite()

	w1 := sqlite.NewWhere().AndEq(`name`, data.Name).AndNq(`rowid`, data.Id)
	count1, err := sqlite.Count(table, *w1)
	if err != nil {
		msg = fmt.Sprintf(`Update bookmark tag error: %v`, err)
		return
	}
	if count1 > 0 {
		msg = com_msg.ExistsNotUpdate(`name`)
		return
	}

	w2 := sqlite.NewWhere().AndEq(`name`, data.Name)
	count2, err := sqlite.Count(`bookmark`, *w2)
	if err != nil {
		msg = fmt.Sprintf(`Update bookmark tag error: %v`, err)
		return
	}
	if count2 > 0 {
		msg = com_msg.HasBeenUsedNotUpdate(`Bookmark tag`)
		return
	}

	updateSet3 := sqlite.NewColumn().Set(`name`, data.Name).Set(`description`, data.Description).Set(`sort`, data.Sort)
	w3 := sqlite.NewWhere().AndEq(`rowid`, data.Id)
	_, err = sqlite.Update(table, *updateSet3, *w3)
	if err != nil {
		msg = fmt.Sprintf(`Update bookmark tag error: %v`, err)
		return
	}

	msg = com_msg.UPD_SUCCESS
	ok = true
	return
}

func Delete(id int64) (msg string, ok bool) {
	sqlite := conf_sql.InitSqlite()

	count1, err := sqlite.CountSql(`select count(*) as c from bookmark where tag = (select name from bookmark_tag where rowid = ?)`, id)
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

	msg = com_msg.DEL_SUCCESS
	ok = true
	return
}
