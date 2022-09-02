package c_request

import (
	"fmt"
	"my-project/src/common"

	"github.com/gin-gonic/gin"
)

type requestDto struct {
	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Data    map[string]string `json:"data"`
	Params  map[string]string `json:"params"`
}

func Router(r *gin.Engine) {
	r.POST(`/api/request`, func(ctx *gin.Context) {
		data := requestDto{}
		ctx.ShouldBindJSON(&data)

		fmt.Println(ctx)

		common.SuccessJson(``, ``, ctx)
	})
}
