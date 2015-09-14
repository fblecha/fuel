package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"os"
	"strings"
	"path/filepath"
	"io/ioutil"
	"github.com/russross/blackfriday"
	//"text/template"
)

type RunCommand struct {
	Name string
	Ui   cli.Ui
}

func (c *RunCommand) Help() string {
	helpText := `
Usage: haiku run

Generate the curent haiku blog in blog/public

`
	return strings.TrimSpace(helpText)
}


func (c *RunCommand) Run(args []string) int {

	if appDir, err := AreWeInProjectDir(); err == nil {
		//create public dir
		err := createPublicDir(appDir)
		if err != nil {
			log.Fatal(err)
			return 1
		}
		//render all content
		err = renderAllContent(appDir)
		if err != nil {
			log.Fatal(err)
			return 1
		}
		return 0
	} else {
		log.Fatal(err)
		return 1
	}
}

func (c *RunCommand) Synopsis() string {
	return "processes all the content to create a new Haiku blog"
}

func createPublicDir(appDir string) error {
	publicDir := fmt.Sprintf("%s/public", appDir)
	return os.MkdirAll(publicDir, 0777)
}

func renderHaiku(path string) {
	fmt.Printf("%s \n", path )
	if input, err := ioutil.ReadFile(path); err != nil {
		output := blackfriday.MarkdownCommon(input)
		fmt.Println(output)
	} else {
		log.Fatal(err)
	}

}

func walkpath(path string, f os.FileInfo, err error) error {
	switch filepath.Ext(path) {
	case ".haiku":
		renderHaiku(path)
	}
	return nil
}

func renderAllContent(appDir string) error {
	//for all *.md files appDir/content/** render them into the public dir at same depth
	//eg. if it's content/blog/post1.md, content/dogs/blacklab.md
	//then they'd be rendered about appDir/public/blog/post1.html and appDir/public/dogs/blacklab.md
	filepath.Walk(appDir, walkpath)

	//TODO also update the persistance layer

	return nil
}
