package main

import (
	"C2E-Wails/types"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func testUpload(host types.Uploader) {
	host.Upload()
}

func main() {
	// Create an instance of the app structure
	catbox := types.NewCatbox()
	pomf := types.NewPomf()
	lobfile := types.NewLobfile()
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
			catbox,
			pomf,
			lobfile,
		},
		Logger: nil,
	})
	if err != nil {
		println("error: ", err.Error())
	}
}
