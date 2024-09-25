package types

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

type Preferences struct {
	UploadHost      string
	TimeoutDuration int
}

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
	prefFilePath := fmt.Sprintf("%s/preferences.txt", appHome)
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

func AppHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	path := fmt.Sprintf("%s/Documents/clip-uploader", home)
	return path
}
