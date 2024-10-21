package main

import (
	apphome "ClipUploader/go/app_home"
	"ClipUploader/go/types"
	"context"
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"github.com/sqweek/dialog"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	appHomeErr := types.CreateAppHome()
	if appHomeErr != nil {
		log.Error("shit went wrong", "error", appHomeErr)
		panic(appHomeErr)
	}
	prefFileErr := types.CreatePrefFileViper()
	if prefFileErr != nil {
		log.Error("shit went wrong", "error", prefFileErr)
		panic(prefFileErr)
	}
	createDbErr := apphome.CreateDatabase()
	if createDbErr != nil {
		log.Error("shit went wrong", "error", prefFileErr)
		panic(createDbErr)
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

func (a *App) HostList() map[string]types.Uploader {
	catbox := types.NewCatbox()
	lobfile := types.NewLobfile()
	pomf := types.NewPomf()
	return map[string]types.Uploader{
		catbox.Name:  catbox,
		lobfile.Name: lobfile,
		pomf.Name:    pomf,
		"Test":       catbox,
	}
}

func (a *App) GetHost() string {
	return viper.GetString("config.host")
}

func (a *App) GetTimeout() int {
	return viper.GetInt("config.timeout")
}

func (a *App) SaveTimeoutDuration(duration int) error {
	if duration <= 0 {
		return fmt.Errorf("duration must be greater than 0, received: %d", duration)
	}
	viper.Set("config.timeout", duration)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("balls")
		return fmt.Errorf("error saving timout: %v", err)
	}
	return nil
}

func (a *App) SaveHost(selectedHost string) error {
	if selectedHost == "" {
		return fmt.Errorf("host must not be empty")
	}
	viper.Set("config.host", selectedHost)
	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("error saving timeout :%v", err)
	}
	return nil
}

func RemovePathFromFile(fileName string) string {
	var name string
	if fileName == "" {
		return ""
	}

	separator := string(os.PathSeparator)
	if strings.Contains(fileName, separator) {
		parts := strings.Split(fileName, separator)
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
	return fileName
}

func (a *App) RemovePathFromFile(fileName string) string {
	var name string
	if fileName == "" {
		return ""
	}

	separator := string(os.PathSeparator)
	if strings.Contains(fileName, separator) {
		parts := strings.Split(fileName, separator)
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
	return fileName
}
