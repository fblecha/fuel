package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"os"
	"strings"
)

type CreateCommand struct {
	Name string
	Ui   cli.Ui
}

func (c *CreateCommand) Help() string {
	helpText := `
Usage: fuel create path/name

Generate a fuel file for name at path.  E.g.

$ cd ~/my_train_site
$ fuel create trains/e-line

will create the file ~fb3/my_train_site/trains/e-line.md

`
	return strings.TrimSpace(helpText)
}

func (c *CreateCommand) Run(args []string) int {
	if len(args) < 2 {
		fmt.Printf("%v", c.Help())
		return 1
	}
	//assume that we want to create app with name args[0] in the current working directory

	appDir, err := AreWeInProjectDir()
	if err != nil {
		log.Fatal(err)
		return 1
	}

	path := args[0]
	contentName := args[1]
	if err := createContent(appDir, path, contentName); err != nil {
		log.Fatal(err)
		return 1
	}
	//else we're all good, so let's return 0
	return 0
}

func (c *CreateCommand) Synopsis() string {
	return "creates a new markdown content file"
}

func createContent(appDir string, path string, contentName string) error {
	contentDir := fmt.Sprintf("%s/content/%s", appDir, path)
	filePath := fmt.Sprintf("%s/%s.%s", contentDir, contentName, EXT)
	fmt.Printf("Creating %s \n", filePath)
	//make contentDir
	os.MkdirAll(contentDir, 0777)
	//make the file
	file, err := os.Create(filePath)
	defer file.Close()
	return err
}
