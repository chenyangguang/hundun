package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
	"strings"
)

var (
	donain    = "https://app02.szmqs.gov.cn"
	start_url = donain + "/0501W/Default.aspx?page=1"
)

func main() {
	fName := "sz-house-summary.csv"
	file, err := os.Create(fName)
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
	})

	c.OnHTML(".xbaiPage_next", func(h *colly.HTMLElement) {
		t := donain + h.Attr("href")
		log.Printf(t)
		c.Visit(t)
	})

	c.Visit(start_url)

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
