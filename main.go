package main

import (
	"joshfah.com/scraper/httpreq"
)

func main() {
	/*seasonURLs := scraping.ScrapeForSeasons("https://aniworld.to/anime/stream/solo-leveling")
	seasons := scraping.ScrapeForEpisodes(seasonURLs)
	redURLs := scraping.ScrapeForURLs(seasons)

	fmt.Println(redURLs) */

	red := httpreq.GetVoeRedirect("https://aniworld.to/redirect/3192028")
	httpreq.Getm3u8(red)
}
