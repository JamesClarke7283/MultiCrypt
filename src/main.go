package main

import (
	"flag"
	"os"
	"runtime/pprof"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/JamesClarke7283/MultiCrypt/src/backend"
	"github.com/JamesClarke7283/MultiCrypt/src/frontend"
	"github.com/JamesClarke7283/MultiCrypt/src/shared"
)

func main() {
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			shared.GetLogger().Fatalf("could not create CPU profile: %v", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			shared.GetLogger().Fatalf("could not start CPU profile: %v", err)
		}
		defer pprof.StopCPUProfile()
	}

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

	frontend.ApplySettings(a, w, config)

	w.ShowAndRun()
}