package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	fileErr := createPreferencesFile()
	if fileErr != nil {
		panic(fileErr)
	}

	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Clip Uploader",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Logger: nil,
	})
	if err != nil {
		println("error: ", err.Error())
	}
}

func createPreferencesFile() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	prefDir := fmt.Sprintf("%s/Documents/preferences.toml", home)
	log.Printf("pref dir: %s", prefDir)

	file, err := os.Stat(prefDir)
	if err != nil {
		_, err := os.Create(prefDir)
		if err != nil {
			log.Println("what the fuck")
			return err
		}
		// impl creating pref file if it doesnt exist
	}
	log.Printf("file exists file size: %d", file.Size())
	return nil
}
