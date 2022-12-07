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
	// Kafka
	router.GET("/kafkaProducer", controllers.KafkaProducer)
	router.GET("/kafkaConsumer", controllers.KafkaConsumer)
	// HTTP Basic Authentication && TOTP
	router.GET("/authTotp", controllers.AuthTotp)

	return router
}
