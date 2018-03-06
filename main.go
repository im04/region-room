package main

import (
	routerSet "region-room/router"
	"github.com/gin-gonic/gin"
	"region-room/config"
	"fmt"
)

type test struct {
	test string
}

func main() {
	router := gin.Default()
	routerSet.InitRouter(router)
	router.Run(":" + fmt.Sprintf("%d", config.ServerConfig.Port))
}
