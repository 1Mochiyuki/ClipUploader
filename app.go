package main

import (
	"context"
	"fmt"
	"strings"

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
}

func (a *App) ChooseFile() string {
	filename, err := dialog.File().Filter("Media files", "mp4", "mov", "mkv", "gif", "txt").Load()
	if err != nil {
		fmt.Println("user did not select file")
	}
	return filename
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
