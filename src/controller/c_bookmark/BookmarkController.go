package c_bookmark

import (
	"my-project/src/common"
	"my-project/src/logger"
	"my-project/src/model"
	"my-project/src/service/s_bookmark"
	"my-project/src/utils/u_string"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, `bookmark.html`, nil)
}

func Router(r *gin.Engine) {

	r.POST(`/api/bookmark/page`, func(ctx *gin.Context) {
		params := model.BookmarkDto{}
		ctx.ShouldBindJSON(&params)
		bookmarks, total, err := s_bookmark.Page(params)
		if err != nil {
			logger.Print(`Get bookmark error: %v`, err)
		}
		common.ErrorPageJson(``, params.Page, params.PageSize, total, bookmarks, ctx)
	})

	r.GET(`/api/bookmark/:id`, func(ctx *gin.Context) {
		id := ctx.Param(`id`)
		idInt64, _ := u_string.Int64(id)
		bookmark, err := s_bookmark.ById(idInt64)
		if err != nil {
			logger.Print(`Get bookmark error: %v`, err)
		}
		common.SuccessJson(``, bookmark, ctx)
	})

	// Insert ro Update bookmark
	r.POST(`/api/bookmark`, func(ctx *gin.Context) {
		params := model.Bookmark{}
		ctx.ShouldBindJSON(&params)
		msg, ok := s_bookmark.Save(&params)
		if ok {
			common.SuccessJson(msg, nil, ctx)
		} else {
			common.ErrorJson(msg, nil, ctx)
		}
	})

	// Delete bookmark tag
	r.DELETE(`/api/bookmark/:id`, func(ctx *gin.Context) {
		id := ctx.Param(`id`)
		idInt64, _ := u_string.Int64(id)
		msg, ok := s_bookmark.Delete(idInt64)
		if ok {
			common.SuccessJson(msg, nil, ctx)
		} else {
			common.ErrorJson(msg, nil, ctx)
		}
	})

	r.GET(`/api/bookmark/bytag/:tag`, func(ctx *gin.Context) {
		tag := ctx.Param("tag")
		bookmarks, err := s_bookmark.ByTag(tag)
		if err != nil {
			logger.Print(`Get bookmarks by tag error: %v`, err)
		}
		common.SuccessJson(``, bookmarks, ctx)
	})
}
