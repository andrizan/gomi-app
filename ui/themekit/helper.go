package themekit

import (
	"image/color"

	"fyne.io/fyne/v2"
)
func ThemeColor(name fyne.ThemeColorName) color.Color {
	app := fyne.CurrentApp()
	return app.Settings().Theme().Color(name, app.Settings().ThemeVariant())
}
