package types

import (
	"fmt"
	"log"
	"os"
)

type Preferences struct {
	UploadHost      string
	TimeoutDuration int
}

func CreatePreferenceFile() error {
	appHome := AppHome()

	_, appHomeExists := os.Stat(appHome)
	if appHomeExists != nil {

		makeAppHomeErr := os.Mkdir(appHome, 0755)
		if makeAppHomeErr != nil {
			return fmt.Errorf("error creating app home directory: %s", makeAppHomeErr.Error())
		}
		log.Println("app home created, trying to write to file again")
		CreatePreferenceFile()
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

func AppHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	path := fmt.Sprintf("%s/Documents/clip-uploader", home)
	return path
}
