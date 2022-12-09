package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "github.com/mohong122/ip2region/binding/golang/ip2region"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"strings"
	"time"
)

var (
	// region   *ip2region.Ip2Region
	searcher *xdb.Searcher
)

// ip2region 版本升級到 2.11.0
func Ip2region(c *gin.Context) {
	InitIP()

	var ip = "1.2.3.4"
	queryIp(ip)
}

// 緩存整個 xdb 數據
func InitIP() {
	// // IP 庫 ip2region.db 方式
	// var dbPath = "./ip2region.db"
	// region, err = ip2region.New(dbPath)
	// if err != nil {
	// 	return err
	// }

	// 1、從 dbPath 加載整個 xdb 到內存
	var dbPath = "./ip2region.xdb"
	cBuff, err := xdb.LoadContentFromFile(dbPath)
	if err != nil {
		fmt.Printf("failed to load content from `%s`: %s\n", dbPath, err)
	}
	// 2、用全局的 cBuff 創建完全基於內存的查詢對象。
	searcher, err = xdb.NewWithBuffer(cBuff)
	if err != nil {
		fmt.Printf("failed to create searcher with content: %s\n", err)
	}
	// 備註：並發使用，用整個 xdb 緩存創建的 searcher 對象可以安全用於並發。
}

// 查詢 xdb 數據
func queryIp(ip string) {
	var tStart = time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
	}
	fmt.Printf("{region: %s, took: %s}\n\n", region, time.Since(tStart))

	regionArr := strings.Split(region, "|")
	if len(regionArr) > 0 {
		fmt.Printf("{regionArr: %+v, took: %s}\n\n", regionArr, time.Since(tStart))
		fmt.Printf("Country=%s \n\n", regionArr[0])
		fmt.Printf("Province=%s \n\n", regionArr[1])
		fmt.Printf("City=%s \n\n", regionArr[2])
		fmt.Printf("County=%s \n\n", regionArr[3])
		fmt.Printf("ISP=%s \n\n", regionArr[4])
	}
}
