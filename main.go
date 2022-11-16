package main

import (
	"Heroku/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 三種模式 DebugMode ReleaseMode TestMode
	gin.SetMode(gin.DebugMode)
	// 路由
	router := router.InitRouter()

	router.Run()
}
