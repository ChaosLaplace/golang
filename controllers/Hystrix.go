package controllers

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

// 熔斷
func Hystrix(c *gin.Context) {
	config := hystrix.CommandConfig{
		Timeout:                2000, // 執行 command 超時時間
		MaxConcurrentRequests:  8,    // command 最大併發量
		SleepWindow:            2000, // 熔斷再次啟動時間。單位毫秒
		ErrorPercentThreshold:  30,   // 錯誤率 請求數量 >= RequestVolumeThreshold & 錯誤率達到設定的 % 則啟動
		RequestVolumeThreshold: 5,    // 熔斷是否打開先看這個 請求閥值(單個窗口 10 秒內統計數量) 設置至少該請求數才計算 錯誤率(此處為 5)
	}
	hystrix.ConfigureCommand("test", config)
	cbs, _, _ := hystrix.GetCircuit("test")

	defer hystrix.Flush()

	for i := 0; i <= 9; i++ {
		start1 := time.Now()
		hystrix.Do("test", RunFunc, func(e error) error {
			fmt.Println("[Hystrix] 服务器错误 触发 fallbackFunc 调用函数执行失败 : ", e)
			return nil
		})
		fmt.Println("[Hystrix] 请求次数:", i+1, ";用时:", time.Now().Sub(start1), ";熔断器开启状态:", cbs.IsOpen(), "请求是否允许：", cbs.AllowRequest())
		time.Sleep(time.Second)
	}
}

func RunFunc() error {
	rand.Seed(time.Now().Unix())
	if rand.Intn(10) > 5 {
		fmt.Println("[Hystrix] RunFunc 执行失败")
		return errors.New("[Hystrix] RunFunc ERROR")
	}
	fmt.Println("[Hystrix] RunFunc 执行成功")
	return nil
}
