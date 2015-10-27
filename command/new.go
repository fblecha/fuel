package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"os"
	"strings"
	//"text/template"
)

type NewCommand struct {
	Name string
	Ui   cli.Ui
}

type fuelDir struct {
	RootDir string //root dir absolute path e.g. /Users/fb3/code/fuel_blog
	DirName string //relative name to RootDir, e.g. app, test, config
}

func (d *fuelDir) Create() error {
	newDir := fmt.Sprintf("%v/%v", d.RootDir, d.DirName)
	return os.MkdirAll(newDir, 0777)
}

func (c *NewCommand) Help() string {
	helpText := `
Usage: fuel new blogname

Generate a new Fuel blog

`
	return strings.TrimSpace(helpText)
}

//Create a new app in the absolute location specified by name
//eg it should be ""/Users/fb3/code/todo" not "todo"
func createAppRootDir(name string) error {
	return os.MkdirAll(name, 0777)
}

//Create all the expected child dirs for a new fuel installation
func createChildDirs(appDir string) error {
	childDirs := [...]fuelDir{
		{appDir, "content"}, //the primary app location -- most new code will go in here
		{appDir, "config"},  //central source for the app config
		{appDir, "style"},
		{appDir, "views"},
		{appDir, "views/partials"},
	}
	for _, child := range childDirs {
		if err := child.Create(); err != nil {
			return err
		}
	}
	//no problems, so return a nil error
	return nil
}

func CreateNewApp(rootDir string, name string) error {
	appDir := fmt.Sprintf("%v/%v", rootDir, name)
	if err := createAppRootDir(appDir); err != nil {
		return err
	}
	if err := createChildDirs(appDir); err != nil {
		return err
	}
	return nil
}

func (c *NewCommand) Run(args []string) int {
	if len(args) < 1 {
		fmt.Printf("%v", c.Help())
		return 1
	}
	//assume that we want to create app with name args[0] in the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return 1
	}
	err = CreateNewApp(currentDir, args[0])
	if err == nil {
		return 0
	} else {
		log.Fatal(err)
		return 1
	}
}

func (c *NewCommand) Synopsis() string {
	return "creates a new Fuel app"
}
