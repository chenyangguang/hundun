package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
	"strings"
)

func main() {
	fName := "hr-tencent.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"position", "category", "hire_num", "address", "publish_day", "link"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("#position tbody tr", func(e *colly.HTMLElement) {
		itemClass := e.Attr("class")

		if strings.EqualFold(itemClass, "even") || strings.EqualFold(itemClass, "odd") {
			writer.Write([]string{
				e.ChildText("td:nth-child(1)"),
				e.ChildText("td:nth-child(2)"),
				e.ChildText("td:nth-child(3)"),
				e.ChildText("td:nth-child(4)"),
				e.ChildText("td:nth-child(5)"),
				e.ChildAttr("a", "href"),
			})
		}
	})

	c.Visit("https://hr.tencent.com/position.php?keywords=&lid=0&start=0#a")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
