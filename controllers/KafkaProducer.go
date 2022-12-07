package controllers

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"time"
)

// Producer server.properties 只配了 9092
func KafkaProducer(c *gin.Context) {
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
	config.Version = sarama.V2_1_0_0

	fmt.Println("start make asyncProducer")
	// 使用配置,新建一個異步生產者
	asyncProducer, err := sarama.NewAsyncProducer([]string{"127.0.0.1:9092", "127.0.0.1:9093", "127.0.0.1:9094"}, config)
	if err != nil {
		fmt.Printf("NewAsyncProducer | err=%v \n", err)
		return
	}
	defer asyncProducer.AsyncClose()
	// 循環判斷哪個通道發送過來數據.
	fmt.Println("start goroutine")
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case suc := <-p.Successes():
				encode, _ := suc.Value.Encode()
				fmt.Printf("KafkaProducer success | topic=%s | value=%s \n", suc.Topic, string(encode))
			case err := <-p.Errors():
				// 這裡報錯可以發送到 Telegram 來警示
				if err != nil {
					fmt.Printf("KafkaProducer | err=%v \n", err)
				} else {
					fmt.Printf("KafkaProducer is nil")
				}
			}
		}
	}(asyncProducer)

	fmt.Println("asyncProducer.Input")
	var value string
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		time11 := time.Now()
		value = "this is a message for test " + time11.Format("15:04:05")
		// 注意 這裡的 msg 必須得是新構建的變量 不然你會發現發送過去的消息內容都是一樣的 因為批次發送消息的關係
		msg := &sarama.ProducerMessage{
			Topic: "test",
			// 將字符串轉化為字節數組
			Value: sarama.ByteEncoder(value),
		}
		// 使用通道發送
		asyncProducer.Input() <- msg
	}
}
