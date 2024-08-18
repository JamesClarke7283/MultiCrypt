package frontend

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/JamesClarke7283/MultiCrypt/src/shared"
)

func ShowSettingsDialog(a fyne.App, w fyne.Window, config *shared.Config) {
	themeSelect := widget.NewSelect([]string{"system", "light", "dark"}, func(selected string) {
		config.Appearance.SelectedTheme = selected
		shared.SaveConfig(config)
		ApplySettings(a, w, config)
	})
	themeSelect.SetSelected(config.Appearance.SelectedTheme)

	fontSizeSlider := widget.NewSlider(10, 24)
	fontSizeSlider.SetValue(float64(config.Appearance.FontSize))
	fontSizeLabel := widget.NewLabel(fmt.Sprintf("Font Size: %d", config.Appearance.FontSize))

	fontSizeSlider.OnChanged = func(value float64) {
		newSize := int(value)
		config.Appearance.FontSize = newSize
		fontSizeLabel.SetText(fmt.Sprintf("Font Size: %d", newSize))
		shared.SaveConfig(config)
		ApplySettings(a, w, config)
	}

	content := container.NewVBox(
		widget.NewLabel("Theme:"),
		themeSelect,
		fontSizeLabel,
		fontSizeSlider,
	)

	dialog.ShowCustom("Settings", "Close", content, w)
}

func ApplySettings(a fyne.App, w fyne.Window, config *shared.Config) {
	// Apply theme
	switch config.Appearance.SelectedTheme {
	case "light":
		a.Settings().SetTheme(theme.LightTheme())
	case "dark":
		a.Settings().SetTheme(theme.DarkTheme())
	default:
		a.Settings().SetTheme(theme.DefaultTheme())
	}

	// Apply font size
	// Note: Fyne doesn't have a built-in way to change font size globally.
	// You might need to implement a custom theme to achieve this.
	// For now, we'll just log the change.
	shared.GetLogger().Infof("Font size changed to %d", config.Appearance.FontSize)

	// Refresh the window to apply changes
	w.Content().Refresh()
}
