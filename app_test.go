package main

import (
	"os"
	"strings"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/assert"
)

func TestRemovePathFromFile(t *testing.T) {
	filepath := `C:\Users\yuki\Videos\Clips\Replay 2024-10-13 23-34-26.mp4`
	msg := RemovePathFromFile(filepath)
	parts := strings.Split(filepath, string(os.PathSeparator))
	wantedString := parts[len(parts)-1]
	want := assert.Equal(t, wantedString, msg)
	if !want {
		t.Fatalf("Did not get correect name back:\nWant: %v\nReceived: %v", wantedString, msg)
	}
}

func TestRemovePathUnix(t *testing.T) {
	filepath := `C:/Users/yuki/Videos/Clips/Replay 2024-10-13 23-34-26.mp4`
	msg := RemovePathFromFile(filepath)
	parts := strings.Split(filepath, "/")
	wantedString := parts[len(parts)-1]
	want := assert.Equal(t, wantedString, msg)
	if !want {
		t.Fatalf("Did not get correect name back:\nWant: %v\nReceived: %v", wantedString, msg)
	}
}

func TestRemovePathEmptyName(t *testing.T) {
	filepath := ""
	msg := RemovePathFromFile(filepath)
	want := assert.Equal(t, filepath, msg)
	if !want {
		t.Fatal("Did not receive empty string")
	}

	log.Info("Test Passed")
}
