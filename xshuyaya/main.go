package main

import (
	"fmt"
	"math/rand"

	"github.com/gocolly/colly"
)

var baseUrl = "https://www.xshuyaya.com"
var startUrl = baseUrl + "/read/5/52735.html"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// 爬取《斗罗大陆》免费小说
func main() {
	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{DomainGlob: "*xshuyaya.*", Parallelism: 2, RandomDelay: 5})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
	})

	//  爬小说
	c.OnHTML("body.read_body", func(e *colly.HTMLElement) {
		content := e.ChildText("#content")
		fmt.Println(content)
	})
	c.OnHTML(".readbutton", func(h *colly.HTMLElement) {
		prev := h.ChildAttr(".readup a", "href")
		root := h.ChildAttr(".readturn a", "href")
		next := h.ChildAttr(".readdown a", "href")
		fmt.Println(prev, root, next)
		c.Visit(next)
	})

	c.Visit(startUrl)
}
