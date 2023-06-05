package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "github.com/crawlab-team/crawlab-go-sdk/entity"
	"github.com/gocolly/colly/v2"
)

// 爬蟲
func Climb(c *gin.Context) {
	// 在colly中使用 Collector 這類物件 來做事情
	collyv2 := colly.NewCollector(
		colly.Async(true),
  		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
	)

	// 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
	// collyv2.OnResponse(func(r *colly.Response) {
	// 	// 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
	// 	fmt.Println( string(r.Body) )
	// })
	// Find and visit all links
	collyv2.OnHTML("a[href]", func(e *colly.HTMLElement) {
		video_name := e.Text
		video_pic  := e.Attr("src")
		video_url  := e.Attr("href")
		// Print link
		fmt.Printf("[Link] name=%q | pic=%s | url=%s\n", video_name, video_pic, video_url)


		// linksStr := e.Attr("href")
		// linksStr = strings.Replace(linksStr, " ", "", -1) // 把空白以""取代掉
		// links := strings.Split(linksStr, "\n")            // 以換行符號(\n)做為分隔來切割字串

		// for _, link := range links {
		// 	c2 := colly.NewCollector()	// 這邊要在迴圈一開始再宣告一個 Collector，才不會與原本的混合到
		// 	c2.OnHTML(".qa-markdown", func(e2 *colly.HTMLElement) {
		// 		fmt.Println(e2.Text) // 印出 qa-markdown class中的文字（Go繁不及備載 文章的內文）
		// 	})
		// 	c2.Visit(link) // 找到<a>連結網址後，點進去訪問
		// }

		// e.Text 印出 <a> tag 裡面的文字，也就是文章標題
		// e.Attr("href") 則是找到 <a> tag裡面的 href元素

		// item := entity.Item {
		// 	"title": e.ChildText("a"),
		// 	"url":   e.ChildAttr("a", "href"),
		// }
		// fmt.Println(item)
	})

	collyv2.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	// Visit 要放最後
	collyv2.Visit("https://zh.xhamster3.com/categories/malaysian")
}
