package c_base64

import (
	"fmt"
	"my-project/src/common"
	"my-project/src/logger"
	"my-project/src/utils/encryption/u_base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, `base64.html`, nil)
}

type base64Dto struct {
	Source string `json:"source"`
}

func Router(r *gin.Engine) {

	r.POST(`/api/base64/encode`, func(ctx *gin.Context) {
		data := base64Dto{}
		ctx.ShouldBindJSON(&data)
		fmt.Println(data)
		common.SuccessJson(``, u_base64.Encode([]byte(data.Source)), ctx)
	})

	r.POST(`/api/base64/decode`, func(ctx *gin.Context) {
		data := base64Dto{}
		ctx.ShouldBindJSON(&data)
		common.SuccessJson(``, u_base64.Decode([]byte(data.Source)), ctx)
	})

	r.POST(`/api/base64/image`, func(ctx *gin.Context) {
		file, fileHeader, err := ctx.Request.FormFile(`file`)
		if err != nil {
			logger.Print(`Upload file error: %v`, err)
			common.SuccessJson(``, ``, ctx)
			return
		}
		result, err := u_base64.EncodeFile(file, fileHeader)
		if err != nil {
			logger.Print(`Upload file error: %v`, err)
			common.SuccessJson(``, ``, ctx)
			return
		}
		common.SuccessJson(``, result, ctx)
	})

}
