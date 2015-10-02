package command

import (
  "os"
	"github.com/mitchellh/cli"
  "testing"
  "fmt"
)


func run() {
	ui := &cli.BasicUi{Writer: os.Stdout}
	args := []string{""}
	runCmd := RunCommand{
		Name: "run",
		Ui:   ui,
	}
	runCmd.Run(args)
}

func new() {
	//use the haiku run command to make a new test blog
	ui := &cli.BasicUi{Writer: os.Stdout}
	args := []string{"runtest"}
	newCmd := NewCommand{
		Name: "new",
		Ui:   ui,
	}
	newCmd.Run(args)
}

func cleanup(currentDir string, t *testing.T) {
	//cleanup by removeing all the directories we created (hopefully)
	dir := fmt.Sprintf("%s/runtest", currentDir)
	if err := os.RemoveAll(dir); err != nil {
		t.Error(err)
	}
}
