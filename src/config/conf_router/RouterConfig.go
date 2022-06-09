package conf_router

// Router configuration
import (
	"my-project/src/controller/c_bookmark"
	"my-project/src/controller/c_bookmark_tag"
	"my-project/src/controller/c_exchange_rate"
	"my-project/src/controller/c_md5"
	"my-project/src/controller/c_search_engine"

	"github.com/gin-gonic/gin"
)

func Config(router *gin.Engine) {
	c_search_engine.Router(router)
	c_exchange_rate.Router(router)
	c_bookmark.Router(router)
	c_bookmark_tag.Router(router)
	c_md5.Router(router)
}
