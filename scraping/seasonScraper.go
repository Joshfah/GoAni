package scraping

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func ScrapeForSeasons(URL string) (seasonURLs []string) {
	var seasons []string

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnHTML("li", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		if strings.HasPrefix(link, "/anime/stream") && strings.Contains(link, "episode") == false {
			season := "https://aniworld.to" + link
			seasons = append(seasons, season)
		}
	})

	c.Visit(URL)
	fmt.Println(seasons)
	return seasons
}
