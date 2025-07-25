package themekit

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Palet warna Mercury
var (
	mercury50  = color.RGBA{R: 254, G: 254, B: 254, A: 255}
	mercury100 = color.RGBA{R: 251, G: 251, B: 251, A: 255}
	mercury200 = color.RGBA{R: 247, G: 247, B: 247, A: 255}
	mercury300 = color.RGBA{R: 243, G: 243, B: 243, A: 255}
	mercury400 = color.RGBA{R: 232, G: 232, B: 231, A: 255}
	mercury500 = color.RGBA{R: 224, G: 224, B: 223, A: 255}
	mercury600 = color.RGBA{R: 207, G: 207, B: 207, A: 255}
	mercury700 = color.RGBA{R: 185, G: 185, B: 184, A: 255}
	mercury800 = color.RGBA{R: 164, G: 164, B: 163, A: 255}
	mercury900 = color.RGBA{R: 122, G: 122, B: 121, A: 255}
	mercury950 = color.RGBA{R: 69, G: 69, B: 69, A: 255}
)

// Implementasi tema Fyne menggunakan warna Mercury
type MercuryTheme struct{}

func (m *MercuryTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return mercury400
	case theme.ColorNameMenuBackground:
		return mercury400
	case theme.ColorNameForeground:
		return mercury950
	case theme.ColorNameDisabled:
		return mercury700
	case theme.ColorNameButton:
		return mercury600
	case theme.ColorNameInputBorder:
		return mercury600
	case theme.ColorNameInputBackground:
		return mercury300
	case theme.ColorNameSeparator:
		return mercury400
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (m *MercuryTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m *MercuryTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m *MercuryTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
