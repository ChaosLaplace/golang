package router

import (
	"Heroku/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 熔斷
	router.GET("/hystrix", controllers.Hystrix)
	// gRPC
	router.GET("/grpcServer", controllers.GrpcServer)
	router.GET("/grpcClient", controllers.GrpcClient)

	return router
}
