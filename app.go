package main

import (
	"C2E-Wails/types"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

func preferencesFileExists() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	pathToPref := fmt.Sprintf("%s/Documents/clip-uploader/preferences.toml", homeDir)

	_, err = os.Stat(pathToPref)
	if err != nil {
		return err
	}
	return nil
}

type PreferencesFile struct {
	Path        string
	Preferences types.Preferences
}

func GetPreferencesFile() (*PreferencesFile, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	pf := PreferencesFile{
		Path: fmt.Sprintf("%s/Documents/clip-uploader/preferences.toml", home),
	}

	log.Printf("set path: %s", pf.Path)
	return &pf, nil
}

func (p *PreferencesFile) File() *os.File {
	if p.Exists() != nil {
		p.Create()
	}
	file, err := os.Open(p.Path)
	if err != nil {
		panic(err)
	}

	return file
}

func preferencesPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	path := fmt.Sprintf("%s/Documents/clip-uploader/preferences.toml", home)
	return path
}

func (p *PreferencesFile) Create() error {
	if p.Exists() == nil {
		return nil
	}
	if p.Path == "" {
		p.Path = preferencesPath()
	}
	prefDir := fmt.Sprintf("%s/Documents/preferences.toml", p.Path)
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

func (p *PreferencesFile) Exists() error {
	if p.Path != "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		p.Path = fmt.Sprintf("%s/Documents/clip-uploader/preferences.toml", home)
	}
	_, err := os.Stat(p.Path)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) SaveTimeoutDuration(duration int) {
	pref, err := GetPreferencesFile()
	if err != nil {
		panic(err)
	}
	if pref.Exists() != nil {
		pref.Create()
	}

	durationString := strconv.Itoa(duration)
	_, err = pref.File().WriteString(durationString)
	if err != nil {
		panic(err)
	}
	pref.Preferences.TimeoutDuration = duration
}

func (a *App) SaveHost(selectedHost string) {
	pref, err := GetPreferencesFile()
	if err != nil {
		panic(err)
	}
	if pref.Exists() != nil {
		pref.Create()
	}

	_, err = pref.File().WriteString(selectedHost)
	if err != nil {
		panic(err)
	}
	pref.Preferences.UploadHost = selectedHost
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
