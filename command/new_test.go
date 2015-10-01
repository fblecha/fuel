package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"os"
	"testing"
)

func TestCreateNewProject(t *testing.T) {
	//use the haiku run command to make a new test blog
	ui := &cli.BasicUi{Writer: os.Stdout}
	cmd := NewCommand{
		Name: "new",
		Ui:   ui,
	}

	args := []string{
		"runtest",
	}
	cmd.Run(args)
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
	CleanupRunTestDirectory(currentDir, t)
}
