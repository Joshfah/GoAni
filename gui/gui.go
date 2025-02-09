package gui

import (
	"fmt"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
	"joshfah.com/scraper/httpreq"
	"joshfah.com/scraper/scraping"
)

func Gui() {

	body := core.NewBody()
	core.NewText(body).SetText("Enter Aniworld-URL here: ")
	URLFrame := core.NewFrame(body)
	URLText := core.NewTextField(URLFrame).SetPlaceholder("for example: https://aniworld.to/anime/stream/a-silent-voice") //.SetType(core.TextFieldOutlined)
	URLText.Styler(func(s *styles.Style) {
		s.Min.Set(units.Pw(100), units.Ph(20))
	})

	SearchButton := core.NewButton(body).SetText("Search")
	SearchButton.OnClick(func(e events.Event) {
		URLs := scraping.ScrapeForSeasons(URLText.Text())
		seasons := scraping.ScrapeForEpisodes(URLs)
		redirects := scraping.ScrapeForURLs(seasons)
		httpreq.VoeDownload(redirects[1])
		SeasonsAndEpisodesFrame := core.NewFrame(body)
		fmt.Println(seasons)
		fmt.Println(redirects)
		for n := range seasons {
			curCount := 0
			fmt.Println(seasons[n].Count)
			if seasons[n].Count == curCount {
				curCount++
			} else {
				core.NewSwitch(SeasonsAndEpisodesFrame)
				SeasonsAndEpisodesFrame.Update()
			}
		}
	})

	body.RunMainWindow()
}
