package main

import (
	"my-project/src/config/conf_router"
	"my-project/src/init/init_db"

	"github.com/gin-gonic/gin"
)

func main() {
	ok := init_db.Init()
	if !ok {
		return
	}

	engine := gin.Default()

	conf_router.Config(engine)

	engine.Run(`:8088`)
}
