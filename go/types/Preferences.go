package types

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Preferences struct {
	UploadHost      string
	TimeoutDuration int
}

const FILE_NAME = "preferences.toml"

func CreateAppHome() error {
	_, createHomeErr := os.Stat(AppHome())
	if createHomeErr != nil {
		err := os.MkdirAll(AppHome(), os.ModeDir)
		if err != nil {
			return err
		}
		log.Info("APP HOME CREATED")
	}
	log.Info("APP HOME EXISTS")
	return nil
}

func CreatePrefFileViper() error {
	viper.SetConfigName("preferences")
	viper.SetConfigType("toml")
	viper.AddConfigPath(AppHome())
	err := viper.ReadInConfig()
	if err != nil {
		log.Info("in here")
		viper.SetDefault("config.host", "Catbox")
		viper.SetDefault("config.timeout", 5000)

		writeErr := viper.SafeWriteConfig()
		if writeErr != nil {
			panic(writeErr)
		}
		log.Info("CONFIG FILE CREATION SUCCESS")
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("CONFIG FILE CHANGED", "msg", e.Name)
	})
	viper.WatchConfig()

	log.Info("CONFIG FILE WATCH SUCCESS")
	return nil
}

func AppHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	path := fmt.Sprintf("%s%sDocuments%sclip-uploader/", home, string(os.PathSeparator), string(os.PathSeparator))
	return path
}
