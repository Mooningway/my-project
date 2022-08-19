package c_md5

import (
	"my-project/src/common"
	"my-project/src/utils/encryption/u_md5"

	"github.com/gin-gonic/gin"
)

type md5Dto struct {
	Source string `json:"source"`
}

func Router(r *gin.Engine) {
	r.POST(`/api/md5`, func(ctx *gin.Context) {
		data := md5Dto{}
		ctx.ShouldBindJSON(&data)
		common.SuccessJson(``, u_md5.HexString(data.Source), ctx)
	})
}
