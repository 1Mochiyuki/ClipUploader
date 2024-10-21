package main

import (
	"ClipUploader/go/db"
	"ClipUploader/go/types"
	"database/sql"
	"embed"

	"github.com/charmbracelet/log"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	catbox := types.NewCatbox()
	pomf := types.NewPomf()
	lobfile := types.NewLobfile()

	log.Info("hello")

	app := NewApp()
	openConn, openConnErr := sql.Open("sqlite3", "database.sqlite")
	if openConnErr != nil {
		panic(openConnErr)
	}
	queries := db.New(openConn)

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
			queries,
		},
	})
	if err != nil {
		println("error: ", err.Error())
	}
}
