package gui

import (
	"fmt"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
)

func Gui() {
	body := core.NewBody()
	core.NewText(body).SetText("Enter URL for Anime here: ")
	URLText := core.NewTextField(body).SetPlaceholder("https://aniworld.to/...")
	core.NewButton(body).SetText("Search").OnClick(func(e events.Event) {
		fmt.Println(URLText.Text())
	})
	body.RunMainWindow()
}
