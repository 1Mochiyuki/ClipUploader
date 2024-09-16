package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
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
	pathToPref := fmt.Sprintf("%s/Documents/clip-uploader/preferences.txt", homeDir)

	_, err = os.Stat(pathToPref)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Hosts() map[string]string {
	return map[string]string{
		"catbox":  "Catbox",
		"lobfile": "lobfile",
		"pomf":    "pomf",
		"other":   "Other",
	}
}

func SetDefaultPreferences(app *App) error {
	setTimeoutDurationErr := app.SaveTimeoutDuration(5000)
	if setTimeoutDurationErr != nil {
		return setTimeoutDurationErr
	}
	setHostErr := app.SaveHost(app.Hosts()["catbox"])
	if setHostErr != nil {
		return setHostErr
	}

	return nil
}

func createPreferenceFile() error {
	appHome := appHome()

	_, appHomeExists := os.Stat(appHome)
	if appHomeExists != nil {

		makeAppHomeErr := os.Mkdir(appHome, 0755)
		if makeAppHomeErr != nil {
			return fmt.Errorf("error creating app home directory: %s", makeAppHomeErr.Error())
		}
		log.Println("app home created, trying to write to file again")
		createPreferenceFile()
		return nil

	}
	log.Println("app directory created")
	prefFilePath := fmt.Sprintf("%s/preferences.txt", appHome)
	_, fileExistErr := os.Stat(prefFilePath)
	if fileExistErr != nil {

		_, err := os.Create(prefFilePath)
		if err != nil {
			return fmt.Errorf("error creating file: %s", err.Error())
		}

		log.Println("pref file created")
	}

	return nil
}

func (a *App) SaveTimeoutDuration(duration int, attempt ...int) error {
	if len(attempt) < 1 || len(attempt) == 0 {
		return fmt.Errorf("attempts must be length of 1, received length of %d", len(attempt))
	}
	attemptInt := attempt[0]
	if attemptInt >= 3 {
		return errors.New("max attempts reached")
	}
	log.Printf("attempt: %d", attemptInt)
	appHome := appHome()

	createFileErr := createPreferenceFile()
	if createFileErr != nil {
		return createFileErr
	}

	prefFilePath := fmt.Sprintf("%s/preferences.txt", appHome)
	file, openFileErr := os.OpenFile(prefFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if openFileErr != nil {
		return openFileErr
	}
	durationString := strconv.Itoa(duration)
	log.Printf("duration as string: %s", durationString)
	bytesWritten, err := fmt.Fprintf(file, "durationTimeout=%d\n", duration)
	if err != nil {
		return fmt.Errorf("error writing to file: %s", err.Error())
	}
	log.Printf("bytes written: %v", bytesWritten)

	return nil
}

func appHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	path := fmt.Sprintf("%s/Documents/clip-uploader", home)
	return path
}

func (a *App) SaveHost(selectedHost string, attempt ...int) error {
	if len(attempt) < 1 || len(attempt) == 0 {
		return fmt.Errorf("attempts must be length of 1, received length of %d", len(attempt))
	}
	attemptInt := attempt[0]
	if attemptInt >= 3 {
		return errors.New("max attempts reached")
	}
	log.Printf("attempt: %d", attemptInt)
	appHome := appHome()

	createFileErr := createPreferenceFile()
	if createFileErr != nil {
		return createFileErr
	}
	prefFilePath := fmt.Sprintf("%s/preferences.txt", appHome)
	_, fileExistErr := os.Stat(prefFilePath)
	if fileExistErr != nil {

		_, err := os.Create(prefFilePath)
		if err != nil {
			return fmt.Errorf("error creating file: %s", err.Error())
		}

		log.Println("pref file crated")
	}

	file, openFileErr := os.OpenFile(prefFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if openFileErr != nil {
		return openFileErr
	}

	log.Printf("duration as string: %s", selectedHost)
	bytesWritten, err := fmt.Fprintf(file, "selectedHost=%s\n", selectedHost)
	if err != nil {
		return fmt.Errorf("error writing to file: %s", err.Error())
	}
	log.Printf("bytes written: %v", bytesWritten)

	return nil
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
