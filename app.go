package main

import (
	"C2E-Wails/go/types"
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/sqweek/dialog"
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
	err := types.CreatePreferenceFile()
	if err != nil {
		log.Fatalf("err: %s", err)
		return
	}
	log.Info("SUCCESS")
}

func (a *App) ChooseFile() string {
	filename, err := dialog.File().Filter("Media files", "mp4", "mov", "mkv", "gif").Load()
	if err != nil {
		fmt.Println("user did not select file")
	}
	return filename
}

func (a *App) Hosts() map[string]types.Uploader {
	catbox := types.NewCatbox()
	lobfile := types.NewLobfile()
	pomf := types.NewPomf()
	return map[string]types.Uploader{
		catbox.Name:  catbox,
		lobfile.Name: lobfile,
		pomf.Name:    pomf,
	}
}

func SetDefaultPreferences(app *App) error {
	setTimeoutDurationErr := app.SaveTimeoutDuration(5000)
	if setTimeoutDurationErr != nil {
		return setTimeoutDurationErr
	}
	setHostErr := app.SaveHost(app.Hosts()["lobfile"])
	if setHostErr != nil {
		return setHostErr
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
	log.Info("attempt: ", "attempt", attemptInt)
	appHome := types.AppHome()

	createFileErr := types.CreatePreferenceFile()
	if createFileErr != nil {
		return createFileErr
	}

	prefFilePath := fmt.Sprintf("%s/preferences.txt", appHome)
	file, openFileErr := os.OpenFile(prefFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if openFileErr != nil {
		return openFileErr
	}
	durationString := strconv.Itoa(duration)
	log.Info("duration as string: ", "duration", durationString)
	bytesWritten, err := fmt.Fprintf(file, "durationTimeout=%d\n", duration)
	if err != nil {
		return fmt.Errorf("error writing to file: %s", err.Error())
	}
	log.Info("bytes written: %v", bytesWritten)

	return nil
}

func (a *App) SaveHost(selectedHost types.Uploader, attempt ...int) error {
	if len(attempt) < 1 || len(attempt) == 0 {
		return fmt.Errorf("attempts must be length of 1, received length of %d", len(attempt))
	}
	attemptInt := attempt[0]
	if attemptInt >= 3 {
		return errors.New("max attempts reached")
	}
	log.Info("attempt: %d", attemptInt)
	appHome := types.AppHome()

	createFileErr := types.CreatePreferenceFile()
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

		log.Info("log file created")
	}

	file, openFileErr := os.OpenFile(prefFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if openFileErr != nil {
		return openFileErr
	}

	log.Info("duration as string: ", "duration", selectedHost)
	switch currHost := selectedHost.(type) {
	case *types.Catbox:

		log.Info("host: %s", currHost.Name)
		_, err := fmt.Fprintf(file, "selectedHost=%s\n", selectedHost)
		if err != nil {
			return fmt.Errorf("error writing to file: %s", err.Error())
		}

	case *types.Lobfile:

		log.Info("host: %s", currHost.Name)
		_, err := fmt.Fprintf(file, "selectedHost=%s\n", selectedHost)
		if err != nil {
			return fmt.Errorf("error writing to file: %s", err.Error())
		}

	case *types.Pomf:

		log.Info("host: %s", currHost.Name)
		_, err := fmt.Fprintf(file, "selectedHost=%s\n", selectedHost)
		if err != nil {
			return fmt.Errorf("error writing to file: %s", err.Error())
		}

	default:
		log.Info("what the fuck")

	}

	return nil
}

func (a *App) RemovePathFromFile(fileName string) string {
	var name string
	if strings.Contains(fileName, "\\") {
		parts := strings.Split(fileName, "\\")
		name = parts[len(parts)-1]
		log.Info("new name (Go): ", "name", name)
		return name
	}
	if strings.Contains(fileName, "/") {
		parts := strings.Split(fileName, "/")
		name = parts[len(parts)-1]
		log.Info("new name (Go): ", "name", name)

		return name
	}
	return ""
}
