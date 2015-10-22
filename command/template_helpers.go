package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	//"regexp"
	//"strings"
	//"bufio"
	//"bytes"
	//"github.com/termie/go-shutil"
	"html/template"
	"strings"
)

func LoadPartialTemplates(appDir string, partialTemplatePaths []string, collectorTemplate *template.Template) *template.Template {
	for _, path := range partialTemplatePaths {
		input, err := ioutil.ReadFile(path)
		if err != nil {
			log.Print(err)
			continue
		}
		//HACK revisit using template.Must
		collectorTemplate = template.Must(collectorTemplate.Clone())
		name := ConvertTemplateName(appDir, path)
		collectorTemplate = template.Must(collectorTemplate.New(name).Parse(string(input)))
	}
	return collectorTemplate
}

func ConvertTemplateName(appDir string, path string) string {
	relPath := strings.Split(path, fmt.Sprintf("%s/views/", appDir))
	result := relPath[1]
	return result //HACK error prone
}

func FindPartialTemplates(appDir string) []string {
	partials := make([]string, 0)

	walker := func(path string, f os.FileInfo, err error) error {
		switch filepath.Ext(path) {
		case ".html":
			//relPath, _ := filepath.Rel(appDir, path)
			//fmt.Println( relPath )
			//absPath, _ := filepath.Abs(relPath)
			partials = append(partials, path)
			return nil
		}
		return nil
	}
	root := fmt.Sprintf("%s/views", appDir)
	//for appDir/views/partials, load all files in that directory into partials
	filepath.Walk(root, walker)
	return partials
}
