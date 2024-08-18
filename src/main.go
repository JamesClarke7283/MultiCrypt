package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/JamesClarke7283/MultiCrypt/src/backend"
	"github.com/JamesClarke7283/MultiCrypt/src/frontend"
	"github.com/JamesClarke7283/MultiCrypt/src/shared"
)

func main() {
	logger := shared.GetLogger()
	logger.Info("Starting MultiCrypt application")

	config, err := shared.LoadConfig()
	if err != nil {
		logger.Fatalf("Failed to load config: %v", err)
	}

	a := app.New()
	w := a.NewWindow("MultiCrypt")

	content := frontend.CreateMainContent(a, w, config, backend.EncryptAES256, backend.DecryptAES256)

	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))

	// Apply settings before showing the window
	frontend.ApplySettings(a, w, config)

	w.ShowAndRun()
}
