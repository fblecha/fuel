package command

import (
	"fmt"
	//"github.com/mitchellh/cli"
	"os"
	"testing"
)

func TestNewHelp(t *testing.T) {
	cmd := new()
	text := cmd.Help()
	if len(text) < 1 {
		t.Fail()
	}
	cleanup(t)
}

func TestNewSynopsis(t *testing.T) {
	cmd := new()
	text := cmd.Synopsis()
	if len(text) < 1 {
		t.Fail()
	}
	cleanup(t)
}

func TestCreateNewProject(t *testing.T) {
	new()
	//test to see if these directories exist
	// runtest/config
	// runtest/content
	// runtest/style

	currentDir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	expectedDirs := []string{
		fmt.Sprintf("%s/runtest/config", currentDir),
		fmt.Sprintf("%s/runtest/content", currentDir),
		fmt.Sprintf("%s/runtest/style", currentDir),
		fmt.Sprintf("%s/runtest/views", currentDir),
		fmt.Sprintf("%s/runtest/views/partials", currentDir),
	}

	for _, expectedDir := range expectedDirs {
		if fileInfo, err := os.Lstat(expectedDir); err == nil {
			if !fileInfo.IsDir() {
				t.Fail()
			}
		} else {
			t.Error(err)
		}
	}
	cleanup(t)
}
