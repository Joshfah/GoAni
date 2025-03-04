package gui

import (
	"fmt"
	"strconv"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
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

	directoryChoosing := core.NewFrame(body)
	//directoryButton := core.NewFileButton(directoryChoosing).SetText("Choose Directory") //Text somehow doesn't show :(
	core.NewText(directoryChoosing).SetText("Choose Directory").Styler(func(s *styles.Style) {
		s.Grow.Set(1, 1)
		s.CenterAll()
	})

	directoryChoosing.Styler(func(s *styles.Style) {
		s.Min.Set(units.Pw(100), units.Ph(20))
	})

	var turnedOnSwitches []*core.Switch

	SearchButton := core.NewButton(body).SetText("Search")
	SearchButton.OnClick(func(e events.Event) {
		URLs := scraping.ScrapeForSeasons(URLText.Text())
		seasons := scraping.ScrapeForEpisodes(URLs)
		fmt.Println(seasons)
		//redirects := scraping.ScrapeForURLs(seasons)
		var allSwitches []*core.Switch
		body.Update()

		seasonEpisodeSwitches := make(map[int][]*core.Switch)
		masterSwitch := core.NewSwitch(body).SetText("Toggle all").SetType(core.SwitchCheckbox)
		body.Update()

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
			allSwitches = append(allSwitches, episodeSwitch)
			episodeSwitch.Styler(func(s *styles.Style) {
				s.CenterAll()
				s.Grow.Set(1, 1)
			})

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

				seasonsSwitch.OnChange(func(e events.Event) {
					for _, eps := range seasonEpisodeSwitches[seasons[n].Count] {
						if seasonsSwitch.IsChecked() {
							eps.SetChecked(true)
						} else {
							eps.SetChecked(false)
						}
					}
					EpisodesFrame.Update()
				})
			}

			for n := range allSwitches {
				if allSwitches[n] == episodeSwitch {
					allSwitches[n].SetText("Episode " + strconv.Itoa(n+1))
				}
			}

			prevNum = seasons[n].Count

			masterSwitch.OnChange(func(e events.Event) {
				for n := range allSwitches {
					if masterSwitch.IsChecked() {
						allSwitches[n].SetChecked(true)
					} else {
						allSwitches[n].SetChecked(false)
					}
				}
			})

			core.UpdateAll()
		}

		downButton := core.NewButton(body).SetIcon(icons.Download)
		downButton.Styler(func(s *styles.Style) {
			s.Justify.Content = styles.End
		})
		downButton.OnClick(func(e events.Event) {
			for n := range allSwitches {
				if allSwitches[n].IsChecked() {
					turnedOnSwitches = append(turnedOnSwitches, allSwitches[n])
					fmt.Println(URLs[n])
				}
			}
		})

	})

	body.Update()

	body.RunMainWindow()
}
