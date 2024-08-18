package shared

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MultiCryptTheme struct{}

var _ fyne.Theme = (*MultiCryptTheme)(nil)

func (m MultiCryptTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if variant == theme.VariantDark {
		switch name {
		case theme.ColorNameBackground:
			return color.NRGBA{R: 28, G: 28, B: 30, A: 255}
		case theme.ColorNameForeground:
			return color.NRGBA{R: 220, G: 220, B: 220, A: 255}
		case theme.ColorNamePrimary:
			return color.NRGBA{R: 0, G: 122, B: 255, A: 255}
		case theme.ColorNameFocus:
			return color.NRGBA{R: 0, G: 122, B: 255, A: 255}
		}
	} else {
		switch name {
		case theme.ColorNameBackground:
			return color.NRGBA{R: 240, G: 240, B: 240, A: 255}
		case theme.ColorNameForeground:
			return color.NRGBA{R: 20, G: 20, B: 20, A: 255}
		case theme.ColorNamePrimary:
			return color.NRGBA{R: 0, G: 122, B: 255, A: 255}
		case theme.ColorNameFocus:
			return color.NRGBA{R: 0, G: 122, B: 255, A: 255}
		}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (m MultiCryptTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m MultiCryptTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m MultiCryptTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
