package model

type ExrateCode struct {
	Id   int    `json:"rowid"`
	Name string `json:"name"`
	Code string `json:"code"`
	Sort int    `json:"sort"`
}

type ExrateRate struct {
	Id       int    `json:"rowid"`
	DateUnix int64  `json:"date_unix"`
	Code     string `json:"code"`
	Sort     int    `json:"sort"`
	Rates    string `json:"rates" type:"BLOB"`
}
