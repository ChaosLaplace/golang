package controllers

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Producer
func Producer(c *gin.Context) {
	config := sarama.NewConfig()
	// 等待服務器所有副本都保存成功後的響應
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 隨機向partition發送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失敗後的響應,只有上面的RequireAcks設置不是NoReponse這裡才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	// 設置使用的kafka版本,如果低於V0_10_0_0版本,消息中的timestrap沒有作用.需要消費和生產同時配置
	// 注意，版本設置不對的話，kafka會返回很奇怪的錯誤，並且無法成功發送消息
	config.Version = sarama.V0_10_0_1

	fmt.Println("start make producer")
	// 使用配置,新建一個異步生產者
	producer, e := sarama.NewAsyncProducer([]string{"127.0.0.1:6667", "127.0.0.1:6667", "127.0.0.1:6667"}, config)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer producer.AsyncClose()

	// 循環判斷哪個通道發送過來數據.
	fmt.Println("start goroutine")
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case <-p.Successes():
				// fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
			case fail := <-p.Errors():
				fmt.Println("err: ", fail.Err)
			}
		}
	}(producer)

	var value string
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		time11 := time.Now()
		value = "this is a message 0606 " + time11.Format("15:04:05")

		// 發送的消息,主題。
		// 注意：這裡的msg必須得是新構建的變量，不然你會發現發送過去的消息內容都是一樣的，因為批次發送消息的關係。
		msg := &sarama.ProducerMessage{
			Topic: "0606_test",
		}

		// 將字符串轉化為字節數組
		msg.Value = sarama.ByteEncoder(value)
		// fmt.Println(value)

		// 使用通道發送
		producer.Input() <- msg
	}
}
