package model

type SearchEngine struct {
	Id     int    `json:"rowid"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Search string `json:"search"`
	Sort   int    `json:"sort"`
}
