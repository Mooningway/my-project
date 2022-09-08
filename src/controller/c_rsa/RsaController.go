package c_rsa

import (
	"my-project/src/common"
	"my-project/src/utils/u_rsa"

	"github.com/gin-gonic/gin"
)

type rsaDto struct {
	Privatekey string `json:"privatekey"`
	Publickey  string `json:"publickey"`
	Text       string `json:"text"`
	Mode       string `json:"mode"`
	Pkcs       string `json:"pkcs"`
	TextOutput string `json:"textOutput"`
	Hash       string `json:"hash"`
	Label      string `json:"label"`
}

type keyDto struct {
	Bits      int    `json:"bits"`
	KeyFormat string `json:"keyFormat"`
}

func Router(r *gin.Engine) {
	r.POST(`/api/rsa/key`, func(ctx *gin.Context) {
		data := keyDto{}
		ctx.ShouldBindJSON(&data)

		result := make(map[string]string)
		privateKey, publicKey, err := u_rsa.GenerateKey(data.Bits, data.KeyFormat)
		if err != nil {
			common.ErrorJson(err.Error(), result, ctx)
		} else {
			result[`privateKey`] = privateKey
			result[`publicKey`] = publicKey
			common.SuccessJson(`success`, result, ctx)
		}
	})

	r.POST(`/api/rsa/encrypt`, func(ctx *gin.Context) {
		data := rsaDto{}
		ctx.ShouldBindJSON(&data)

		result := ``
		var err error
		if data.Mode == `PKCS1v15` {
			result, err = u_rsa.EncryptPKCS1v15(data.Text, data.Publickey, data.TextOutput)
		} else if data.Mode == `OAEP` {
			h, err1 := u_rsa.NewHash(data.Hash)
			if err1 == nil {
				var label []byte
				if data.Label == `` {
					label = nil
				} else {
					label = []byte(data.Label)
				}
				result, err = u_rsa.EncryptOaep(data.Text, data.Publickey, data.TextOutput, h, label)
			} else {
				err = err1
			}
		}
		if err != nil {
			common.ErrorJson(err.Error(), result, ctx)
		} else {
			common.SuccessJson(`success`, result, ctx)
		}
	})

	r.POST(`/api/rsa/decrypt`, func(ctx *gin.Context) {
		data := rsaDto{}
		ctx.ShouldBindJSON(&data)

		result := ``
		var err error
		if data.Mode == `PKCS1v15` {
			result, err = u_rsa.DecryptPKCS1v15(data.Text, data.Privatekey, data.TextOutput, data.Pkcs)
		} else if data.Mode == `OAEP` {
			h, err1 := u_rsa.NewHash(data.Hash)
			if err1 == nil {
				var label []byte
				if data.Label == `` {
					label = nil
				} else {
					label = []byte(data.Label)
				}
				result, err = u_rsa.DecryptOaep(data.Text, data.Privatekey, data.TextOutput, data.Pkcs, h, label)
			} else {
				err = err1
			}
		}
		if err != nil {
			common.ErrorJson(err.Error(), result, ctx)
		} else {
			common.SuccessJson(`success`, result, ctx)

		}
	})
}
