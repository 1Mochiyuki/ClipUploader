package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sqweek/dialog"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ChooseFile() string {
	filename, err := dialog.File().Filter("Media files", "mp4", "mov", "mkv", "gif").Load()
	if err != nil {
		fmt.Println("user did not select file")
	}
	return filename
}

func (a *App) SaveStringPreference(pref string) {
	// impl save to file
}

func (a *App) SaveIntPreference(pref int) {
	// impl save to file
}

func (a *App) UploadViaLobfile(absPath string) {
	// impl upload to lobfile docs: https://lobfile.com/api/v3/docs/
}

func (a *App) UploadViaPomf(absPath string) {
	log.Printf("uploading: %s to", absPath)
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		panic(err)
	}
	if fileInfo.Size() > (1024 * 1024 * 1024) {
		msg := fmt.Sprintf("File: %s is over the 1 GB limit.\nSize: %d", fileInfo.Name(), fileInfo.Size())
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "File Size Limit",
			Message:       msg,
			DefaultButton: "OK",
		})
		return
	}
	files := map[string]string{"files[]:": absPath}
	data, err := json.Marshal(files)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
}

func (a *App) RemovePathFromFile(fileName string) string {
	var name string
	if strings.Contains(fileName, "\\") {
		parts := strings.Split(fileName, "\\")
		name = parts[len(parts)-1]
		fmt.Printf("new name (Go): %s\n", name)
		return name
	}
	if strings.Contains(fileName, "/") {
		parts := strings.Split(fileName, "/")
		name = parts[len(parts)-1]
		fmt.Printf("new name (Go): %s\n", name)

		return name
	}
	return ""
}
