package conf_static

// Static resource configuration
import (
	"my-project/src/controller/c_bookmark"
	"my-project/src/controller/c_bookmark_tag"
	"my-project/src/controller/c_exchange_rate"

	"github.com/gin-gonic/gin"
)

func Config(engine *gin.Engine) {
	engine.Delims(`{{{`, `}}}`)

	engine.LoadHTMLGlob(`template/**/*`)
	engine.Static(`static`, `static`)

	engine.GET(`/exrate`, c_exchange_rate.Index)
	engine.GET(`/bookmark/tag`, c_bookmark_tag.Index)
	engine.GET(`/bookmark`, c_bookmark.Index)
}
