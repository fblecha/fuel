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
	log.Printf("creating %v \n", newDir)
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
	log.Printf("app root dir = %v\n", name)
	//mode := int(0777)
	return os.MkdirAll(name, 0777)
}

func createChildDirs(appDir string) {
	childDirs := [...]fuelDir{
		{appDir, "content"},        //the primary app location -- most new code will go in here
		{appDir, "config"},         //central source for the app config
		{appDir, "style"},          //central source for the app config
		{appDir, "views"},          //central source for the app config
		{appDir, "views/partials"}, //central source for the app config
	}
	for _, child := range childDirs {
		if err := child.Create(); err != nil {
			panic(err)
		}
	}
}

func CreateNewApp(rootDir string, name string) {
	appDir := fmt.Sprintf("%v/%v", rootDir, name)

	createAppRootDir(appDir)
	createChildDirs(appDir)
}

func (c *NewCommand) Run(args []string) int {
	if len(args) < 1 {
		fmt.Printf("%v", c.Help())
		return 1
	}
	//assume that we want to create app with name args[0] in the current working directory

	if currentDir, err := os.Getwd(); err == nil {
		CreateNewApp(currentDir, args[0])
		return 0
	} else {
		log.Fatal(err)
		return 1
	}
}

func (c *NewCommand) Synopsis() string {
	return "creates a new Fuel app"
}
