package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"os"
	"strings"
	"path/filepath"
	"io/ioutil"
	"regexp"
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

	if appDir, err := AreWeInProjectDir(); err == nil {
		if input, err := ioutil.ReadFile(path); err == nil {
			output := blackfriday.MarkdownCommon(input)

			re := regexp.MustCompile("/content/(.+)")
			//locs := re.FindStringIndex(path)
			//oldPath := path[locs[0]:locs[1]]
			fmt.Printf("submatch = %q \n", re.FindStringSubmatch(path) )
			//fmt.Printf("from: %s \n to: %s \n\n", oldPath, newPath )
			if matches := re.FindStringSubmatch(path); matches != nil && len(matches)==2 {
				oldPath := matches[1]
				//tmpPath is the expected location before some manipulation around the filename, e.g. convert
				// blah.haiku to blah.html
				tmpPath := fmt.Sprintf("%s/public/%s", appDir, oldPath)

				//convert from .haiku to .html
				filename := filepath.Base(tmpPath)
				extensionIndex := strings.LastIndex(filename, ".")
				newFilename := filename[0:extensionIndex]
				//finally make the new dir and the new path (for the file to be created
				newDir := filepath.Dir(tmpPath)
				newPath := fmt.Sprintf("%s/%s.html", newDir, newFilename)

				//make the new dir in public
				os.MkdirAll(newDir, 0777)
				//output is a []byte -- write it to a file
				err = ioutil.WriteFile(newPath, output, 0644)
			}




			fmt.Println(output)
		} else {
			log.Fatal(err)
		}
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
