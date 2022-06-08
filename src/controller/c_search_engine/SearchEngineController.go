package c_search_engine

import (
	"my-project/src/common"
	"my-project/src/logger"
	"my-project/src/service/s_search_engine"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.GET(`/api/searchengine/all`, func(ctx *gin.Context) {
		searchEngines, err := s_search_engine.All()
		if err != nil {
			logger.Print(`Get all search engines error: %v`, err)
		}
		common.SuccessJson(``, searchEngines, ctx)
	})
}
