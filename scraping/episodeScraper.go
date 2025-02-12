package scraping

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Episode struct {
	Name, URL string
	Count     int
}

type Season struct {
	Episodes []Episode
	Count    int
	IsMovie  bool
}

func ScrapeForEpisodes(URLs []string) (s []Season) {
	var seasons []Season
	//	var episodes []Episode

	for n := range URLs {
		c := colly.NewCollector()
		episode := Episode{}
		season := Season{}

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting: ", r.URL)
		})

		c.OnHTML("td.seasonEpisodeTitle", func(e *colly.HTMLElement) {
			link := e.ChildAttr("a", "href")
			URL := "https://aniworld.to" + link
			episode.URL = URL
			season.IsMovie = true
			for i := 0; i <= 9; i++ {
				suffix := strconv.Itoa(i)
				if strings.HasSuffix(URLs[n], suffix) {
					season.IsMovie = false
				}
			}
			season.Count = n + 1
			seasons = append(seasons, season)
		})

		c.OnHTML("a", func(t *colly.HTMLElement) {
			if strings.HasPrefix(t.Attr("href"), "/anime/stream") {
				episodeName := t.ChildText("span")
				episode.Name = episodeName
			}
		})
		season.Episodes = append(season.Episodes, episode)

		c.Visit(URLs[n])
	}
	return seasons
}
