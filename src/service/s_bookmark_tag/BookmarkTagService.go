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
		query := sqlite.NewQuery()
		if len(dto.Keyword) > 0 {
			query.AndLike(`name`, dto.Keyword)
		}

		total, err = sqlite.Count(table, *query)
		if err != nil {
			return err
		}

		if total > 0 {
			query.Desc(`sort`).Asc(`rowid`).Page(dto.Page, dto.PageSize)
			err = sqlite.FindSlice(table, *query, &bookmarkTags, `*`, `rowid`)
		}
		return err
	})
	return
}

func All() (tags []model.BookmarkTag, err error) {
	sqlite := conf_sql.InitSqlite()
	query := sqlite.NewQuery().Desc(`sort`).Asc(`rowid`)
	err = sqlite.FindSlice(table, *query, &tags, `rowid`, `*`)
	return
}

func ById(id int64) (tag model.BookmarkTag, err error) {
	sqlite := conf_sql.InitSqlite()
	sqlite.FindById(table, id, &tag, `rowid`, `*`)
	return
}

func Save(data *model.BookmarkTag) (msg string, ok bool) {
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

	query := sqlite.NewQuery().AndEq(`name`, data.Name)
	count, err := sqlite.Count(table, *query)
	if err != nil {
		msg = fmt.Sprintf(`Save bookmark tag error: %v`, err)
		return
	}
	if count > 0 {
		msg = com_msg.ExistsNotAdd(`name`)
		return
	}
	_, err = sqlite.Insert(table, data)
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

	query1 := sqlite.NewQuery().AndEq(`name`, data.Name).AndNe(`rowid`, data.Id)
	count1, err := sqlite.Count(table, *query1)
	if err != nil {
		msg = fmt.Sprintf(`Update bookmark tag error: %v`, err)
		return
	}
	if count1 > 0 {
		msg = com_msg.ExistsNotUpdate(`name`)
		return
	}

	query2 := sqlite.NewQuery().AndEq(`name`, data.Name)
	count2, err := sqlite.Count(`bookmark`, *query2)
	if err != nil {
		msg = fmt.Sprintf(`Update bookmark tag error: %v`, err)
		return
	}
	if count2 > 0 {
		msg = com_msg.HasBeenUsedNotUpdate(`Bookmark tag`)
		return
	}

	_, err = sqlite.UpdateById(table, data, data.Id)
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

	_, err = sqlite.DeleteById(table, id)
	if err != nil {
		msg = fmt.Sprintf(`Delete bookmark tag error: %v`, err)
		return
	}

	msg = com_msg.DEL_SUCCESS
	ok = true
	return
}
