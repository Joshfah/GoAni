package gui

import "cogentcore.org/core/core"

func Gui() {
	body := core.NewBody()
	core.NewText(body)
	body.RunMainWindow()
}
