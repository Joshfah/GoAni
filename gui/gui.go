package gui

import (
	"fmt"
	"strconv"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
	"joshfah.com/scraper/scraping"
)

func Gui() {
	body := core.NewBody()
	core.NewText(body).SetText("Enter Aniworld-URL here: ")
	URLFrame := core.NewFrame(body)
	URLText := core.NewTextField(URLFrame).SetPlaceholder("for example: https://aniworld.to/anime/stream/a-silent-voice")
	URLText.Styler(func(s *styles.Style) {
		s.Min.Set(units.Pw(100), units.Ph(20))
	})

	SearchButton := core.NewButton(body).SetText("Search")
	SearchButton.OnClick(func(e events.Event) {
		URLs := scraping.ScrapeForSeasons(URLText.Text())
		seasons := scraping.ScrapeForEpisodes(URLs)
		redirects := scraping.ScrapeForURLs(seasons)
		fmt.Println(redirects)
		core.NewSwitch(body).SetType(core.SwitchCheckbox)
		body.Update()

		// Store episodeSwitches for each season
		seasonEpisodeSwitches := make(map[int][]*core.Switch)

		var prevNum int
		for n := range seasons {
			SeasonsAndEpisodesFrame := core.NewFrame(body)
			SeasonsAndEpisodesFrame.Styler(func(s *styles.Style) {
				s.Display = styles.Grid
				s.Columns = 1
				s.Max.X.Em(7.5)
			})

			EpisodesFrame := core.NewFrame(body)
			EpisodesFrame.Styler(func(s *styles.Style) {
				s.Display = styles.Grid
				s.Columns = 1
				s.Max.X.Em(10)
			})

			episodeSwitch := core.NewSwitch(EpisodesFrame).SetType(core.SwitchCheckbox)
			episodeSwitch.Styler(func(s *styles.Style) {
				s.CenterAll()
				s.Grow.Set(1, 1)
			})

			// Add the episodeSwitch to the corresponding season's list
			seasonEpisodeSwitches[seasons[n].Count] = append(seasonEpisodeSwitches[seasons[n].Count], episodeSwitch)

			var seasonText string
			seasonText = "Season " + strconv.Itoa(seasons[n].Count)

			if seasons[n].IsMovie {
				seasonText = "Movie/s"
			}

			if seasons[n].IsMovie == false && seasons[0].IsMovie {
				seasonText = "Season " + strconv.Itoa(seasons[n].Count-1)
			}

			if n == 0 || seasons[n].Count != prevNum {
				seasonsSwitch := core.NewSwitch(SeasonsAndEpisodesFrame).SetType(core.SwitchCheckbox).SetText(seasonText)
				seasonsSwitch.Styler(func(s *styles.Style) {
					s.CenterAll()
					s.Grow.Set(1, 1)
				})

				// Add an OnChange event handler to the seasonSwitch
				seasonsSwitch.OnChange(func(e events.Event) {
					// Toggle all episodeSwitches in this season
					for _, eps := range seasonEpisodeSwitches[seasons[n].Count] {
						eps.SetChecked(true)
						EpisodesFrame.Update()
					}
				})
			}
			prevNum = seasons[n].Count

			SeasonsAndEpisodesFrame.Update()
			EpisodesFrame.Update()
		}
	})

	body.RunMainWindow()
}
