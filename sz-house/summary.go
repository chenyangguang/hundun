package main

import (
	"encoding/csv"
	//"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gocolly/colly"
	"strings"
)

var (
	donain       = "https://app02.szmqs.gov.cn"
	start_url    = donain + "/0501W/Default.aspx?page=1"
	property_url = donain + "/0501W/Iframe/LicItemIframe.aspx?licId="
	property_dir = "./property"
)

func main() {
	if _, err := os.Stat(property_dir); os.IsNotExist(err) {
		os.Mkdir(property_dir, 0777)
	}
	fName := "sz-house-summary.csv"
	f := filepath.Join(property_dir, fName)
	file, err := os.Create(f)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"company", "property", "location", "area", "id"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("#caTable tbody tr", func(e *colly.HTMLElement) {
		itemClass := e.Attr("class")
		evenClass := "tab_body bd0"
		oddClass := "tab_body bd1"
		if strings.EqualFold(itemClass, evenClass) || strings.EqualFold(itemClass, oddClass) {
			writer.Write([]string{
				e.ChildText("td:nth-child(1)"),
				e.ChildText("td:nth-child(2)"),
				e.ChildText("td:nth-child(3)"),
				e.ChildText("td:nth-child(4)"),
				e.ChildAttr("td:nth-child(5) a", "val"),
			})
		}
		go property_crawler(e)
	})

	c.OnHTML(".xbaiPage_next", func(h *colly.HTMLElement) {
		t := donain + h.Attr("href")
		log.Printf(t)
		c.Visit(t)
	})

	c.Visit(start_url)

	log.Printf("Scraping finished, check file %q for results\n", fName)
}

// 楼盘划分继续分页爬
func property_crawler(p *colly.HTMLElement) {
	name := p.ChildText("td:nth-child(2)")

	if name == "" {
		return
	}
	fName := filepath.Join(property_dir, name+".csv")
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"房号", "建筑面积", "建筑面积单价", "套内面积", "套内面积单价", "总售价", "销售状态", "备注"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("#caTable tbody tr", func(e *colly.HTMLElement) {
		itemClass := e.Attr("class")
		evenClass := "tab_body"
		oddClass := "tab_body bd1"
		if strings.EqualFold(itemClass, evenClass) || strings.EqualFold(itemClass, oddClass) {
			writer.Write([]string{
				e.ChildText("td:nth-child(1)"),
				e.ChildText("td:nth-child(2)"),
				e.ChildText("td:nth-child(3)"),
				e.ChildText("td:nth-child(4)"),
				e.ChildText("td:nth-child(5)"),
				e.ChildText("td:nth-child(6)"),
				e.ChildText("td:nth-child(7)"),
				e.ChildText("td:nth-child(8)"),
			})
		}
	})

	c.OnHTML(".xbaiPage_next", func(n *colly.HTMLElement) {
		t := donain + n.Attr("href")
		log.Printf(t)
		c.Visit(t)
	})

	property_start := property_url + p.ChildAttr("td:nth-child(5) a", "val")
	c.Visit(property_start)

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
