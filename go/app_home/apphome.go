package apphome

import (
	"ClipUploader/go/errs"
	"ClipUploader/go/types"
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

func IsDatabaseExist(path string) error {
	if path == "" {
		return errs.NewEmptyStringError()
	}
	_, err := os.Stat(path)
	return err
}

func CreateDatabase() error {
	appHome := types.AppHome()
	dbPath := fmt.Sprintf("%s/database.sqlite", appHome)

	err := IsDatabaseExist(dbPath)
	if errors.Is(err, errs.EmptyStringError{}) {
		panic(err)
	}
	if err != nil {
		_, appHomeExist := os.Stat(appHome)
		if appHomeExist != nil {
			types.CreateAppHome()
			log.Info("Retrying")
			return CreateDatabase()
		}
		_, err := os.Create(dbPath)
		if err != nil {
			panic(err)
		}
		log.Info("Database created")
		return nil
	}
	log.Info("Database exists")
	return nil
}
