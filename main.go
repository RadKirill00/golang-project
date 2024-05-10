package main

import (
	
	"log"

	"github.com/gocolly/colly"
)

func main() {

	logger := log.Default()

	c := colly.NewCollector()

	c.OnHTML("a[href]", func(r *colly.HTMLElement) {
			link := r.Attr("href")
			logger.Print("Link: ", link)
	})


		c.Visit("https://scrapeme.live/shop/")

		c.Wait()
}
