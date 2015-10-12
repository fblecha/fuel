package command

import (
	"errors"
	"fmt"
	"github.com/russross/blackfriday"
	"github.com/mitchellh/cli"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	//"bufio"
	"bytes"
	"github.com/termie/go-shutil"
	"html/template"
	//"github.com/microcosm-cc/bluemonday"
)

type RunCommand struct {
	Name string
	Ui   cli.Ui
}

func (c *RunCommand) Help() string {
	helpText := `
Usage: fuel run

Generate the curent fuel blog in blog/public

`
	return strings.TrimSpace(helpText)
}

func (c *RunCommand) Run(args []string) int {
	//TODO refactor this into a sequence of functions that are applied via a loop ?
	appDir, err := AreWeInProjectDir()
	if err != nil {
		log.Fatal(err)
		return 1
	}
	//create public dir
	if err := createPublicDir(appDir); err != nil {
		log.Fatal(err)
		return 1
	}
	//render all content
	if err := renderAllContent(appDir); err != nil {
		log.Fatal(err)
		return 1
	}
	//copy the style directory over to /public
	if err := copyStyleDirToPublic(appDir); err != nil {
		log.Fatal(err)
		return 1
	}
	//both of the commands worked, we're good to go
	return 0
}

func (c *RunCommand) Synopsis() string {
	return "process all the content to create a new fuel blog"
}

func createPublicDir(appDir string) error {
	publicDir := fmt.Sprintf("%s/public", appDir)
	return os.MkdirAll(publicDir, 0777)
}

func renderFuel(path string) error {

	if appDir, err := AreWeInProjectDir(); err == nil {
		if jsonMap, markdownStr, err := SplitJsonAndMarkdown(path); err == nil {
			storeJSON(jsonMap)
			renderMarkdown(appDir, path, markdownStr)
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}

func storeJSON(json map[string]interface{}) error {
	//fmt.Printf("%s \n", json)
	return nil
}

func loadHTML(appDir string, path string) (string, error) {
	//given that path is in ./public/something/maybe/content.md

	//do most exact matching to least exact matching
	// if we have ./views/blah/foo/bar/greenfrog.md
	// we'd try ./views/blah/foo/bar/greenfrog.layout.html
	// then ./views/blah/foo/bar/layout.html
	// then ./views/blah/foo/layout.html
	// then ./views/blah/layout.html
	// then ./views/layout.html

	//first need to make it relative
	relativePath, err := GetRelativePath(appDir, path)

	//fmt.Printf("relativePath = %s appDir = %s path=%s \n", relativePath, appDir, path)


	if err != nil {
		return "", err
	}
	dirs := PathToDirs(relativePath)         //gives back most general to most specific
	//fmt.Printf("dirs = %q \n", dirs)
	dirs = Reverse(dirs)                     //now in most specific to general order
	targets := addContentTargetsToDirs(dirs) //now each dir has a target of something/layout.html
	//fmt.Printf("targets = %q \n", targets)

	//target is a form of layout.html file that we'll use as the template
	result, err := findBestMatch(targets)
	//fmt.Printf("result = %s \n", result)
	return result, err
}

func addContentTargetsToDirs(dirs []string) []string {
	results := make([]string, len(dirs))
	for i := range dirs {
		results[i] = fmt.Sprintf("%s/layout.html", dirs[i])
	}
	return results
}

//For all the targets, find the most specific match.  When found, return the string that corresponds to that template.
func findBestMatch(targets []string) (string, error) {
	noMatchError := errors.New("no match found")

	if len(targets) > 0 {
		//can we load target[0]?
		path := targets[0]

		//fmt.Printf("path = %s \n", path)

		//currentDir, _ := os.Getwd()

		if file, err := ioutil.ReadFile(path); err == nil {
			return string(file), nil
		} else { //if not, then let's see if we can find it in targets[1:]
			//fmt.Println(err)
			if len(targets) > 1 {
				return findBestMatch(targets[1:])
			} else {
				return "", noMatchError
			}
		}
	} else {
		return "", noMatchError
	}
}

func renderMarkdown(appDir string, path string, markdownContent string) {
	renderer, extensions := configureBlackFriday(path)
	content := blackfriday.Markdown([]byte(markdownContent), renderer, extensions)

	re2 := regexp.MustCompile("/content/")

	src := []byte(path)
	replacement := []byte("/views/")
	htmlPath := re2.ReplaceAll(src, replacement)

	template, _ := loadHTML(appDir, string(htmlPath) )
	//fmt.Printf("template = %s \n", template)

	result, err := ParseAndInsert(string(content), template)
	if err != nil {
		log.Fatal(err)
	}

	//log.Printf("result = %s \n", result)


	//html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	re := regexp.MustCompile("/content/(.+)")
	if matches := re.FindStringSubmatch(path); matches != nil && len(matches) == 2 {
		oldPath := matches[1]
		//tmpPath is the expected location before some manipulation around the filename, e.g. convert
		// blah.fuel to blah.html
		tmpPath := fmt.Sprintf("%s/public/%s", appDir, oldPath)
		//convert from .fuel to .html
		newDir, newPath := convertFromFuelToHTML(tmpPath)
		//make the new dir in public
		os.MkdirAll(newDir, 0777)
		//output is a []byte -- write it to a file

		ioutil.WriteFile(newPath, []byte(result), 0644)
	}
}

func convertFromFuelToHTML(tmpPath string) (string, string) {
	newFilename := getFilenameMinusExtension(tmpPath)
	//finally make the new dir and the new path (for the file to be created
	newDir := filepath.Dir(tmpPath)
	newPath := fmt.Sprintf("%s/%s.html", newDir, newFilename)
	return newDir, newPath
}

func walkpath(path string, f os.FileInfo, err error) error {
	switch filepath.Ext(path) {
	case EXT:
		return renderFuel(path)
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

func configureBlackFriday(path string) (blackfriday.Renderer, int) {
	htmlFlags := 0 //blackfriday.HTML_COMPLETE_
	htmlFlags |= blackfriday.HTML_SKIP_HTML
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS

	title := getFilenameMinusExtension(path)
	css := ""
	renderer := blackfriday.HtmlRenderer(htmlFlags, title, css)

	extensions := 0
	extensions |= blackfriday.EXTENSION_LAX_HTML_BLOCKS

	// extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	// extensions |= blackfriday.EXTENSION_TABLES
	// extensions |= blackfriday.EXTENSION_FENCED_CODE
	// extensions |= blackfriday.EXTENSION_AUTOLINK
	// extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	// extensions |= blackfriday.EXTENSION_SPACE_HEADERS

	return renderer, extensions
}

func SplitJsonAndMarkdown(filename string) (map[string]interface{}, string, error) {
	var results [2]string
	if str, err := ioutil.ReadFile(filename); err == nil {
		for i, rune := range bytes.Split(str, []byte{'~', '~', '~'}) { //split by "~~~"
			//fmt.Printf("Counter %d :  %s\n", i, string(rune))
			results[i] = string(rune)
		}

		//fmt.Printf("results = %v and len = %s and results[1] = %v \n", results, len(results), len(results[1])==0 )

		if isEmpty(results[1]) {
			/*
					Is the 2nd element of the array empty? If so, we likely didn't parse out a separater.
				we likely have a situation like [current] (see below), so we need to convert it.
				current:
					results[0] = markdown string
					results[1] = empty
				need:
					results[0] = "{}"
					results[1] = markdown string
			*/
			results[1] = results[0]
			results[0] = "{}"
		}
		if len(results) == 2 {
			//I have markdown but not json, so let's just make an empty json to process
			if isEmpty(results[0]) {
				results[0] = "{}"
			}
		}

		dat, _ := parseJSON(results[0])
		return dat, results[1], nil
	} else {
		return nil, "", err
	}
}

func copyStyleDirToPublic(appDir string) error {
	//first we have to remove the old /public/style directory
	dest := fmt.Sprintf("%s/public/style", appDir)
	if err := os.RemoveAll(dest); err != nil {
		return err
	}
	//now we copy over the new /style directory
	src := fmt.Sprintf("%s/style", appDir)
	return shutil.CopyTree(src, dest, nil)
}

func ParseAndInsert(content string, htmlTemplate string) (string, error) {
	var data = make(map[string]interface{})
	data["Content"] = template.HTML(content)

	t := template.New("t")
	t, err := t.Parse(  htmlTemplate )
	if err != nil {
		return "", err
	}

	var b bytes.Buffer

	if err := t.Execute(&b, data ); err != nil {
		return "", err
	}
	return b.String(), nil
}
