package scraping

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ScrapeForURLs(seasons []Season) (redURLS []string) {

	var redURLs []string

	for n := range seasons {
		for i := range seasons[n].Episodes {
			c := colly.NewCollector()

			c.OnRequest(func(r *colly.Request) {
				fmt.Println("Visiting: ", r.URL)
			})

			c.OnHTML("a.watchEpisode", func(e *colly.HTMLElement) {
				class := e.ChildAttr("i", "class")
				if class == "icon VOE" {
					href := e.Attr("href")
					fmt.Println(href)
					RedURL := "https://aniworld.to" + href
					redURLs = append(redURLs, RedURL)
				}
			})
			c.Visit(seasons[n].Episodes[i].URL)
		}
	}
	return redURLs
}
