package c_random

import (
	"my-project/src/common"
	"my-project/src/utils/u_random"
	"strings"

	"github.com/gin-gonic/gin"
)

type paramsDto struct {
	Len       int `json:"len"`
	Number    int `json:"number"`
	Lower     int `json:"lower"`
	Upper     int `json:"upper"`
	Character int `json:"character"`
}

func Router(r *gin.Engine) {

	r.POST(`/api/randomString`, func(ctx *gin.Context) {
		params := paramsDto{}
		ctx.BindJSON(&params)

		len := params.Len
		if len <= 0 {
			common.SuccessJson(``, ``, ctx)
			return
		}
		if len > 100 {
			len = 100
		}

		var seed strings.Builder
		if params.Number == 1 {
			seed.WriteString(u_random.SEED_NUMBER)
		}
		if params.Lower == 1 {
			seed.WriteString(u_random.SEED_LOWER)
		}
		if params.Upper == 1 {
			seed.WriteString(u_random.SEED_UPPER)
		}
		if params.Character == 1 {
			seed.WriteString(u_random.SEED_CHARACTER)
		}
		if seed.Len() == 0 {
			common.SuccessJson(``, ``, ctx)
			return
		}

		common.SuccessJson(``, u_random.GenString(seed.String(), len), ctx)
	})
}
