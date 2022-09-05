package c_rsa

import (
	"my-project/src/common"
	"my-project/src/utils/u_rsa"

	"github.com/gin-gonic/gin"
)

type rsaDto struct {
	Bits           int    `json:"bits"`
	Privatekey     string `json:"privatekey"`
	Publickey      string `json:"publickey"`
	Text           string `json:"text"`
	Pkcs           string `json:"pkcs"`
	OutputEncoding string `json:"outputEncoding"`
}

func Router(r *gin.Engine) {
	r.POST(`/api/rsa/x509/key`, func(ctx *gin.Context) {
		data := rsaDto{}
		ctx.ShouldBindJSON(&data)

		result := make(map[string]string)
		privateKey, publicKey, err := u_rsa.GenerateKeyX509(data.Bits, data.Pkcs)
		if err != nil {
			common.ErrorJson(err.Error(), result, ctx)
		} else {
			result[`privateKey`] = privateKey
			result[`publicKey`] = publicKey
			common.SuccessJson(`success`, result, ctx)
		}
	})

	r.POST(`/api/rsa/x509/encrypt`, func(ctx *gin.Context) {
		data := rsaDto{}
		ctx.ShouldBindJSON(&data)
		result, err := u_rsa.EncryptX509(data.Text, data.Publickey, data.OutputEncoding)
		if err != nil {
			common.ErrorJson(err.Error(), result, ctx)
		} else {
			common.SuccessJson(`success`, result, ctx)

		}
	})

	r.POST(`/api/rsa/x509/decrypt`, func(ctx *gin.Context) {
		data := rsaDto{}
		ctx.ShouldBindJSON(&data)
		result, err := u_rsa.DecryptX509(data.Text, data.Privatekey, data.Pkcs, data.OutputEncoding)
		if err != nil {
			common.ErrorJson(err.Error(), result, ctx)
		} else {
			common.SuccessJson(`success`, result, ctx)

		}
	})
}
