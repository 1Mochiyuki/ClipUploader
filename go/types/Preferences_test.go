package types

import (
	"testing"
)

func TestPreferenceFileCreation(t *testing.T) {
	err := CreatePrefFileViper()
	want := err == nil
	if !want {
		t.Fatalf("Create Preferences File failed: %v", err)
	}
}

func TestAppHomeDirCreation(t *testing.T) {
	err := CreateAppHome()
	want := err == nil
	if !want {
		t.Fatalf("Create Preferences File failed: %v", err)
	}
}
