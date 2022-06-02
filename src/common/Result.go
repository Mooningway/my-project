package common

// Result for gin
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type result struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type resultPage struct {
	Code      string      `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Page      int64       `json:"page"`
	PageSize  int64       `json:"pageSize"`
	Total     int64       `json:"total"`
	FirstPage int64       `jons:"firstPage"`
	PrevPage  int64       `jons:"prevPage"`
	NextPage  int64       `jons:"nextPage"`
	LastPage  int64       `jons:"lastPage"`
}

func SuccessJson(msg string, data interface{}, ctx *gin.Context) {
	data1 := data
	if data1 == nil {
		data1 = make(map[string]interface{})
	}
	ctx.JSON(http.StatusOK, result{Code: `200`, Data: data1, Msg: msg})
}

func ErrorJson(msg string, data interface{}, ctx *gin.Context) {
	data1 := data
	if data1 == nil {
		data1 = make(map[string]interface{})
	}
	ctx.JSON(http.StatusOK, result{Code: `500`, Data: data1, Msg: msg})
}

func SuccessPageJson(msg string, page, pageSize, total int64, data interface{}, ctx *gin.Context) {
	data1 := data
	if data1 == nil {
		data1 = make(map[string]interface{})
	}
	lastPage, prevPage, nextPage := initPage(page, pageSize, total)
	ctx.JSON(http.StatusOK, resultPage{Code: `200`, Data: data1, Msg: msg,
		Page: page, PageSize: page, Total: total, FirstPage: 1, PrevPage: prevPage, NextPage: nextPage, LastPage: lastPage})
}

func ErrorPageJson(msg string, page, pageSize, total int64, data interface{}, ctx *gin.Context) {
	data1 := data
	if data1 == nil {
		data1 = make(map[string]interface{})
	}
	lastPage, prevPage, nextPage := initPage(page, pageSize, total)
	ctx.JSON(http.StatusOK, resultPage{Code: `500`, Data: data1, Msg: msg,
		Page: page, PageSize: page, Total: total, FirstPage: 1, PrevPage: prevPage, NextPage: nextPage, LastPage: lastPage})
}

func initPage(page, pageSize, total int64) (lastPage, prevPage, nextPage int64) {
	if pageSize <= 0 {
		lastPage = total / 10
	} else {
		lastPage = total / pageSize
	}
	if pageSize != 0 && total%pageSize != 0 {
		lastPage++
	}
	if lastPage <= 0 {
		lastPage = 1
	}
	prevPage = page - 1
	if prevPage < 1 {
		prevPage = int64(1)
	}
	nextPage = page + 1
	if nextPage > lastPage {
		nextPage = lastPage
	}
	return
}
