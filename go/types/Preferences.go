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

func CreatePreferenceFile() error {
	appHome := AppHome()

	_, appHomeExists := os.Stat(appHome)
	if appHomeExists != nil {
		log.Info("app home doesnt exist, creating now")

		makeAppHomeErr := os.Mkdir(appHome, 0755)
		if makeAppHomeErr != nil {
			return fmt.Errorf("error creating app home directory: %s", makeAppHomeErr.Error())
		}
		log.Info("app home created, trying to write to file again")
		CreatePreferenceFile()
		return nil

	} else {
		log.Info("app home already exists")
	}
	prefFilePath := fmt.Sprintf("%s/%s", appHome, FILE_NAME)
	_, fileExistErr := os.Stat(prefFilePath)
	if fileExistErr != nil {

		_, err := os.Create(prefFilePath)
		if err != nil {
			return fmt.Errorf("error creating file: %s", err.Error())
		}

		log.Info("pref file created")
		return nil
	}
	log.Info("preferences file already exists")

	return nil
}

func CreatePrefFileViper() error {
	viper.SetConfigName("preferences")
	viper.SetConfigType("toml")
	viper.AddConfigPath(AppHome())
	err := viper.ReadInConfig()
	if err != nil {
		viper.Set("config.host", "Catbox")
		viper.Set("config.timeout", 5000)
		writeErr := viper.SafeWriteConfig()
		if writeErr != nil {
			panic(writeErr)
		}
		log.Info("CONFIG CREATION SUCCESS")

	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("CONFIG CHANGED", "msg", e.Name)
	})
	viper.WatchConfig()

	log.Info("CONFIG WATCH SUCCESS")
	return nil
}

func AppHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	path := fmt.Sprintf("%s/Documents/clip-uploader", home)
	return path
}
