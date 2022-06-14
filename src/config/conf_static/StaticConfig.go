package conf_static

// Static resource configuration
import (
	"my-project/src/controller/c_base64"
	"my-project/src/controller/c_bookmark"
	"my-project/src/controller/c_bookmark_tag"
	"my-project/src/controller/c_exchange_rate"
	"my-project/src/controller/c_index"
	"my-project/src/controller/c_md5"

	"github.com/gin-gonic/gin"
)

func Config(engine *gin.Engine) {
	engine.Delims(`{{{`, `}}}`)

	engine.LoadHTMLGlob(`template/**/*`)
	engine.Static(`static`, `static`)

	engine.GET(`/`, c_index.Index)
	engine.GET(`/exrate`, c_exchange_rate.Index)
	engine.GET(`/bookmark/tag`, c_bookmark_tag.Index)
	engine.GET(`/bookmark`, c_bookmark.Index)
	engine.GET(`/md5`, c_md5.Index)
	engine.GET(`/base64`, c_base64.Index)
}
