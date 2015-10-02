package command

import (
	"fmt"
	//"github.com/mitchellh/cli"
	"os"
	"testing"
)

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
	cleanup(currentDir, t)
}
