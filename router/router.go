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
	// HENNGE Mission 1 - Write a program which fulfills the requirements below
	router.GET("/missionHENNGE", controllers.MissionHENNGE)
	// HENNGE Mission 2 - HTTP Basic Authentication && TOTP
	router.GET("/authTotp", controllers.AuthTotp)

	return router
}
