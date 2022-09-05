package c_aes

import (
	"my-project/src/common"
	"my-project/src/utils/encryption/u_aes"

	"github.com/gin-gonic/gin"
)

type aesDto struct {
	Key            string `json:"key"`
	Text           string `json:"text"`
	Mode           string `json:"mode"`
	Nonce          string `json:"nonce"`
	Padding        string `json:"padding"`
	OutputEncoding string `json:"outputEncoding"`
}

func Router(r *gin.Engine) {
	r.POST(`/api/aes/gcm/encrypt`, func(ctx *gin.Context) {
		data := aesDto{}
		ctx.ShouldBindJSON(&data)
		result, err := u_aes.AesGcmEncrypt(data.Key, data.Text, data.Nonce, data.OutputEncoding)
		if err != nil {
			common.ErrorJson(err.Error(), result, ctx)
		} else {
			common.SuccessJson(`success`, result, ctx)
		}
	})

	r.POST(`/api/aes/gcm/decrypt`, func(ctx *gin.Context) {
		data := aesDto{}
		ctx.ShouldBindJSON(&data)
		result, err := u_aes.AesGcmDecrypt(data.Key, data.Text, data.Nonce, data.OutputEncoding)
		if err != nil {
			common.ErrorJson(err.Error(), result, ctx)
		} else {
			common.SuccessJson(`success`, result, ctx)
		}
	})

	r.POST(`/api/aes/cbc/encrypt`, func(ctx *gin.Context) {
		data := aesDto{}
		ctx.ShouldBindJSON(&data)
		result, err := u_aes.AesCbcEncrypt(data.Key, data.Text, data.Nonce, data.OutputEncoding)
		if err != nil {
			common.ErrorJson(err.Error(), result, ctx)
		} else {
			common.SuccessJson(`success`, result, ctx)
		}
	})

	r.POST(`/api/aes/cbc/decrypt`, func(ctx *gin.Context) {
		data := aesDto{}
		ctx.ShouldBindJSON(&data)
		result, err := u_aes.AesCbcDecrypt(data.Key, data.Text, data.Nonce, data.OutputEncoding)
		if err != nil {
			common.ErrorJson(err.Error(), result, ctx)
		} else {
			common.SuccessJson(`success`, result, ctx)
		}
	})
}
