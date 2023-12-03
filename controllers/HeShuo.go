package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 1.如下程式碼不增加與刪減行數(不含 package 與 import)情況下請修正輸出"預期結果"，並且說明原因。
/*
	預期結果 :
	{Max 10}
	{Max 20}

	原因 : 取址後用指標取得記憶體存放的內容, 並修改原始內容。
*/
func HeShuo1(c *gin.Context) {
	student := people{"Max", 10}
	fmt.Println(student)
	modify(&student)
	fmt.Println(student)
}

func modify(p *people) {
	p.age = p.age + 10
}

type people struct {
	name string
	age  int
}

// 2.請修正以下程式輸出預期結果，限定修正其中兩行，不可新增行數。
/*
	預期結果 :
	2
	2
	2
	2
	2

	原因 : num 會 ++，所以再重新賦值。
*/
func HeShuo2(c *gin.Context) {
	num := 1 // 不可修改
	for i := 0; i < 5; i++ {
		num = 1
		func() {
			num++            // 不可修改
			fmt.Println(num) // 不可修改
		}()
	}
	time.Sleep(2 * time.Second)
}

// 3.如下程式碼請排除3個錯誤(不含 package 與 import)並且 能夠印出"預期結果"，並且說明原因。
/*
	預期結果 :
	hello
	SayHi

	錯誤1 : fmt.Println(job) -> job 是 ActFun func()。
	錯誤2 : 只打印出 hello -> 調整容量 make(chan ActFun, 1)。
	錯誤3 : 。
*/
type ActFun func() interface{}

var MessageQueue chan ActFun

func HeShuo3(c *gin.Context) {
	MessageQueue = make(chan ActFun, 2)
	go sendMessageQueue()
	enMessageQueue(Hello, MessageQueue)
	enMessageQueue(SayHi, MessageQueue)
}

func enMessageQueue(action ActFun, mq chan<- ActFun) bool {
	select {
	case mq <- action:
		return true
	default:
		return false
	}
}

func sendMessageQueue() {
	for job := range MessageQueue {
		fmt.Println(job())
	}
}

func Hello() interface{} {
	return "hello"
}

func SayHi() interface{} {
	return "SayHi"
}

// 4.以下golang程式會發生什麼問題？要如何解決？
/*
	報錯 : fatal error: concurrent map read and map write

	解決方式 : map 為引用類型，高並發時對 map 並發寫會產生競爭，不管是否同一個 key 都會報錯，讀不會有問題。
	1.使用 channel
	2.使用 sync.map
	3.如果必須定義全域 map，需要加鎖
	優化1.可以採用 cow 策略，read 不加鎖，每次修改 copy 修改，再賦值
	優化2.可以採用分片鎖，減少鎖的粒度
*/
var GoMap = make(map[int]string)
var lock sync.Mutex

func MapRead() {
	i := 0
	for {
		fmt.Println(GoMap[i])
		i++

		if i > 10 {
			break
		}
	}
}

func MapWrite() {
	i := 0
	for {
		// 加鎖期間其他協程會進入阻塞狀態直到解鎖
		lock.Lock()
		GoMap[i] = fmt.Sprintf("%d", i)
		lock.Unlock()

		i++
		if i > 10 {
			break
		}
	}
}

func HeShuo4(c *gin.Context) {
	runtime.GOMAXPROCS(2)
	go MapWrite()
	go MapRead()
	time.Sleep(time.Duration(20) * time.Millisecond)
}

// 5.請使用golang在命令列上畫出一個等腰三角形, 高度為變數(2+n), 底為變數(3+2*n), n為整數且n >= 0。
/*
	範例如下:
	*
	***
	*****
*/
func HeShuo5(c *gin.Context) {
	n := 1
	height := 2 + n

	for i := 0; i < height; i++ {
		stars := strings.Repeat("*", 1+i*2)
		fmt.Println(stars)
	}
}

// 6.使用 golang 建立一個 RESTful API server，使用 8080 port，開啟一個 api服 務可以依 id 取得 MySQL 資料庫中
// 該筆資料內容，並使用 json 格式回傳。
/*
	資料庫格式 :
	id name createtime
	1 Jack 2020-02-02 02:02:02
	2 Alice 2020-02-02 02:02:12

	API格式 :
	http://{server_url}:8080/{id}

	回傳格式 :
	{ "id": 1, "name": "Jack", "createtime": "2020-02-02 02:02:02" }
*/
/*
CREATE TABLE IF NOT EXISTS `api_test` (
  `id` int(10) unsigned NOT NULL,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `api_test` (`id`, `name`, `createtime`) VALUES
	(1, 'Jack', '2020-02-01 18:02:02'),
	(2, 'Alice', '2020-02-01 18:02:12');
*/
var Db *sqlx.DB

type ApiTest struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Createtime string `db:"createtime"`
}

func initDB() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang_test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
}

func HeShuo6(c *gin.Context) {
	initDB()

	var apiTest []ApiTest
	id := c.Param("id")
	err := Db.Select(&apiTest, "select id, name, createtime from api_test where id=?", id)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	c.JSON(200, apiTest)
}

// 7.使用 golang 呼叫上述 API 服務印出回傳結果，每3秒呼叫一次，執行檔格式 {exe} {id}。
/*
	"範例" :
	./client 2
	id: 2
	name: Alice
	createtime: 2020-02-02 02:02:12

	go build -o client.exe

	if len(os.Args) < 2 {
		fmt.Println("Missing command line argument")
		return
	}
	id := os.Args[1]
*/
func HeShuo7(c *gin.Context) {
	// exe := c.Param("exe")
	id := c.Param("id")

	for {
		resp, err := http.Get("http://localhost:8080/heShuo6/" + id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var apiTest []ApiTest
		err = json.Unmarshal([]byte(string(body)), &apiTest)
		if err != nil {
			fmt.Println("Failed to unmarshal response:", err)
			return
		}
		for _, item := range apiTest {
			fmt.Printf("id: %d\nname: %s\ncreatetime: %s\n", item.Id, item.Name, item.Createtime)
		}

		time.Sleep(3 * time.Second)
	}
}
