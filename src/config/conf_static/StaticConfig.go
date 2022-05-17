package conf_static

// Static resource configuration
import (
	"my-project/src/controller/c_exchange_rate"

	"github.com/gin-gonic/gin"
)

func Config(engine *gin.Engine) {
	engine.LoadHTMLGlob(`template/**/*`)
	engine.Static(`static`, `static`)

	engine.GET(`/exrate`, c_exchange_rate.Index)
}
