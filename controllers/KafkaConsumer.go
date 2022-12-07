package controllers

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"sync"
)

// Consumer server.properties 只配了 9092
func KafkaConsumer(c *gin.Context) {
	address := []string{"127.0.0.1:9092", "127.0.0.1:9093", "127.0.0.1:9094"}
	topic := "test"
	wg := &sync.WaitGroup{}
	wg.Add(2)
	// 廣播式消費：消費者1
	go clusterConsumer(wg, address, topic, "group-1")
	// 廣播式消費：消費者2
	go clusterConsumer(wg, address, topic, "group-2")

	wg.Wait()
}

// 支持 address cluster的消費者
func clusterConsumer(wg *sync.WaitGroup, address []string, topic, groupId string) {
	defer wg.Done()

	config := sarama.NewConfig()
	config.Producer.Return.Errors = true // 是否等待成功和失敗後的響應,只有上面的 RequiredAcks 設置不是 NoReponse 這裡才有用
	config.Version = sarama.V2_1_0_0     // 設置使用的kafka版本,如果低於V0_10_0_0版本,消息中的timestrap沒有作用.需要消費和生產同時配置
	// init consumer
	consumer, err := sarama.NewConsumer(address, config)
	if err != nil {
		fmt.Printf("NewConsumer | groupId=%s | err=%v", groupId, err)
		return
	}
	defer consumer.Close()
	// 根據消費者獲取指定的主題分區的消費者, offset 為偏移量, sarama.OffsetOldest: 為從頭開始消費, sarama.OffsetNewest 為從最新的偏移量開始消費, 0: 即獲當前已消費的偏移量
	partition := 0
	consumePartition, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("consumePartition | topic=%s | partition=%v | OffsetNewest=%d | err=%v", topic, partition, sarama.OffsetNewest, err)
		return
	}
	defer consumePartition.Close()
	// trap SIGINT to trigger a shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// 循環等待接收信息
	for {
		select {
		case msg := <-consumePartition.Messages():
			msgStr := string(msg.Value)
			fmt.Printf("KafkaConsumer=%s \n", msgStr)
		case err := <-consumePartition.Errors():
			if err != nil {
				fmt.Printf("KafkaConsumer | err=%v \n", err)
			} else {
				fmt.Printf("KafkaConsumer is nil")
			}
		}
	}
}
