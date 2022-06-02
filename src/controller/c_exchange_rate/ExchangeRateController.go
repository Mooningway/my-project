package c_exchange_rate

import (
	"encoding/json"
	"my-project/src/common"
	"my-project/src/logger"
	"my-project/src/service/s_exchange_rate"
	"net/http"

	"github.com/gin-gonic/gin"
)

type convert struct {
	FromCode string `json:"fromCode"`
	ToCode   string `json:"toCode"`
	Amount   string `json:"amount"`
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, `exchange_rate.html`, nil)
}

func Router(r *gin.Engine) {

	// Get currency code list
	r.GET(`/api/exrate/code`, func(ctx *gin.Context) {
		codes, err := s_exchange_rate.Codes()
		if err != nil {
			logger.Print(`Get currency code error: %v`, err)
		}
		jsonBytes, _ := json.Marshal(codes)
		common.SuccessJson(``, string(jsonBytes), ctx)
	})

	// Update exchange rate for config
	r.PUT(`/api/exrate/rate/:code`, func(ctx *gin.Context) {
		code := ctx.Param(`code`)
		msg := s_exchange_rate.PullAndSaveRates(code)
		if len(msg) > 0 {
			common.ErrorJson(msg, nil, ctx)
		} else {
			common.SuccessJson(`Update data success`, nil, ctx)
		}
	})

	// Convert exchange rate
	r.POST(`/api/exrate/convert`, func(ctx *gin.Context) {
		data := convert{}
		ctx.ShouldBindJSON(&data)
		msg, ok := s_exchange_rate.Exchange(data.FromCode, data.ToCode, data.Amount)
		if ok {
			common.SuccessJson(msg, nil, ctx)
		} else {
			common.ErrorJson(msg, nil, ctx)
		}
	})

	// Get reates Of currency code
	r.GET(`/api/exrate/ratedata`, func(ctx *gin.Context) {
		rates, err := s_exchange_rate.RatesData()
		if err != nil {
			logger.Print(`Get rates data error: %v`, err)
		}
		jsonBytes, _ := json.Marshal(rates)
		common.SuccessJson(``, string(jsonBytes), ctx)
	})

	// Delete rate
	r.DELETE(`/api/exrate/rate/:code`, func(ctx *gin.Context) {
		code := ctx.Param(`code`)
		err := s_exchange_rate.DeleteRateByCode(code)
		if err != nil {
			logger.Print(`Delete rate data error: %v`, err)
		}
		common.SuccessJson(``, nil, ctx)
	})
}
