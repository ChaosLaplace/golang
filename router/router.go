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
	// ip2region 版本升級到 2.11.0
	router.GET("/ip2region", controllers.Ip2region)
	// 爬蟲
	router.GET("/climb", controllers.Climb)

	// HENNGE 面試題 1 - Write a program which fulfills the requirements below
	router.GET("/missionHENNGE", controllers.MissionHENNGE)
	// HENNGE 面試題 2 - HTTP Basic Authentication && TOTP
	router.GET("/authTotp", controllers.AuthTotp)

	// 禾碩資訊 面試題
	router.GET("/heShuo1", controllers.HeShuo1)
	router.GET("/heShuo2", controllers.HeShuo2)
	router.GET("/heShuo3", controllers.HeShuo3)
	router.GET("/heShuo4", controllers.HeShuo4)
	router.GET("/heShuo5", controllers.HeShuo5)
	router.GET("/heShuo6/:id", controllers.HeShuo6)
	router.GET("/heShuo7/:exe/:id", controllers.HeShuo7)

	return router
}
