package model

type BookmarkDto struct {
	Keyword  string `json:"keyword"`
	Tag      string `json:"tag"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
}

type Bookmark struct {
	Id          int    `json:"rowid"`
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Sort        int    `json:"sort"`
}

type BookmarkTag struct {
	Id          int64  `json:"rowid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Sort        int    `json:"sort"`
}
