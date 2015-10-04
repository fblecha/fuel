package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"os"
	"testing"
)

func run() RunCommand {
	ui := &cli.BasicUi{Writer: os.Stdout}
	args := []string{""}
	runCmd := RunCommand{
		Name: "run",
		Ui:   ui,
	}
	runCmd.Run(args)
	return runCmd
}

func new() NewCommand {
	//use the haiku run command to make a new test blog
	ui := &cli.BasicUi{Writer: os.Stdout}
	args := []string{"runtest"}
	newCmd := NewCommand{
		Name: "new",
		Ui:   ui,
	}
	newCmd.Run(args)
	return newCmd
}

func cleanup(t *testing.T) {
	currentDir, _ := os.Getwd()
	//cleanup by removeing all the directories we created (hopefully)
	dir := fmt.Sprintf("%s/runtest", currentDir)
	if err := os.RemoveAll(dir); err != nil {
		t.Error(err)
	}
}

func cdRuntest(t *testing.T) {
	wd, _ := os.Getwd()
	testdir := fmt.Sprintf("%s/runtest", wd)
	if err := os.Chdir(testdir); err != nil {
		t.Error(err)
	}
}

func cdBack(t *testing.T) {
	wd, _ := os.Getwd()
	testdir := fmt.Sprintf("%s/..", wd)
	if err := os.Chdir(testdir); err != nil {
		t.Error(err)
	}
}
