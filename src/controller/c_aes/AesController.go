package c_aes

import (
	"errors"
	"my-project/src/common"
	"my-project/src/logger"
	"my-project/src/utils/encryption/u_aes"
	"strings"

	"github.com/gin-gonic/gin"
)

type aesDto struct {
	Key            string `json:"key"`
	Text           string `json:"text"`
	Mode           string `json:"mode"`
	Nonce          string `json:"nonce"`
	Padding        string `json:"padding"`
	OutputEncoding string `json:"outputEncoding"`
	Operation      string `json:"operation"`
}

func Router(r *gin.Engine) {
	r.POST(`/api/aes`, func(ctx *gin.Context) {
		// Currently only gcm
		data := aesDto{}
		ctx.ShouldBindJSON(&data)

		var result string
		var err error

		if strings.ToLower(data.Operation) == `encrypt` {
			// Encrypt
			result, err = u_aes.AesGcmEncrypt(data.Key, data.Text, data.Nonce, data.OutputEncoding)
			if err != nil {
				logger.Print(`Aes encrypt error: %v`, err)
			}
		} else if strings.ToLower(data.Operation) == `decrypt` {
			// Decrypt
			result, err = u_aes.AesGcmDecrypt(data.Key, data.Text, data.Nonce, data.OutputEncoding)
			if err != nil {
				logger.Print(`Aes decrypt error: %v`, err)
			}
		} else {
			err = errors.New(`error`)
		}

		if err != nil {
			common.SuccessJson(err.Error(), result, ctx)
		} else {
			common.SuccessJson(data.Operation+` success`, result, ctx)
		}
	})
}
