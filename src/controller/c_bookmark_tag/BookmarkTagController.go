package c_bookmark_tag

import (
	"my-project/src/common"
	"my-project/src/logger"
	"my-project/src/model"
	"my-project/src/service/s_bookmark_tag"
	"my-project/src/utils/u_string"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, `bookmark_tag.html`, nil)
}

func Router(r *gin.Engine) {

	r.POST(`/api/bookmark/tag/page`, func(ctx *gin.Context) {
		params := model.BookmarkDto{}
		ctx.ShouldBindJSON(&params)
		bookmarkTags, total, err := s_bookmark_tag.Page(params)
		if err != nil {
			logger.Print(`Get bookmark tags error: %v`, err)
		}
		common.ErrorPageJson(``, params.Page, params.PageSize, total, bookmarkTags, ctx)
	})

	r.GET(`/api/bookmark/tag`, func(ctx *gin.Context) {
		bookmarkTags, err := s_bookmark_tag.All()
		if err != nil {
			logger.Print(`Get all bookmark tags error: %v`, err)
		}
		common.SuccessJson(``, bookmarkTags, ctx)
	})

	r.GET(`/api/bookmark/tag/:id`, func(ctx *gin.Context) {
		id := ctx.Param(`id`)
		idInt64, _ := u_string.Int64(id)
		bookmarkTag, err := s_bookmark_tag.ById(idInt64)
		if err != nil {
			logger.Print(`Get bookmark tag error: %v`, err)
		}
		common.SuccessJson(``, bookmarkTag, ctx)
	})

	// Insert bookmark tag
	r.POST(`/api/bookmark/tag`, func(ctx *gin.Context) {
		params := model.BookmarkTag{}
		ctx.ShouldBindJSON(&params)
		msg, ok := s_bookmark_tag.Insert(params)
		if ok {
			common.SuccessJson(msg, nil, ctx)
		} else {
			common.ErrorJson(msg, nil, ctx)
		}
	})

	// Update bookmark tag
	r.PUT(`/api/bookmark/tag`, func(ctx *gin.Context) {
		params := model.BookmarkTag{}
		ctx.ShouldBindJSON(&params)
		msg, ok := s_bookmark_tag.Update(params)
		if ok {
			common.SuccessJson(msg, nil, ctx)
		} else {
			common.ErrorJson(msg, nil, ctx)
		}
	})

	// Delete bookmark tag
	r.DELETE(`/api/bookmark/tag/:id`, func(ctx *gin.Context) {
		id := ctx.Param(`id`)
		idInt64, _ := u_string.Int64(id)
		msg, ok := s_bookmark_tag.Delete(idInt64)
		if ok {
			common.SuccessJson(msg, nil, ctx)
		} else {
			common.ErrorJson(msg, nil, ctx)
		}
	})
}
