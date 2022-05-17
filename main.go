package main

import (
	"log"
	"my-project/src/config/conf_router"
	"my-project/src/config/conf_static"
	"my-project/src/init/init_db"

	"github.com/gin-gonic/gin"
)

func main() {
	err := init_db.Init()
	if err != nil {
		log.Println(err)
		return
	}

	engine := gin.Default()

	conf_static.Config(engine)
	conf_router.Config(engine)

	engine.Run(`:8088`)
}
