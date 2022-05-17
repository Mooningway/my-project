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

func SuccessJson(msg string, data interface{}, ctx *gin.Context) {
	if data == nil {
		ctx.JSON(http.StatusOK, result{Code: `200`, Data: make(map[string]interface{}), Msg: msg})
	} else {
		ctx.JSON(http.StatusOK, result{Code: `200`, Data: data, Msg: msg})
	}
}

func ErrorJson(msg string, data interface{}, ctx *gin.Context) {
	if data == nil {
		ctx.JSON(http.StatusOK, result{Code: `500`, Data: make(map[string]interface{}), Msg: msg})
	} else {
		ctx.JSON(http.StatusOK, result{Code: `500`, Data: data, Msg: msg})
	}
}
