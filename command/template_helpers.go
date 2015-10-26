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
	if len(partialTemplatePaths) < 1 {
		log.Printf("Expect partial templates to be len > 0")
		return collectorTemplate
	}
	for _, path := range partialTemplatePaths {
		input, err := ioutil.ReadFile(path)
		if err != nil {
			log.Print(err)
			continue
		}

		//HACK revisit using template.Must
		//collectorTemplate, err = collectorTemplate.Clone()
		if err != nil {
			log.Printf("Error was %s \n", err)
			return collectorTemplate
		}
		name := ConvertTemplateName(appDir, path)
		//fmt.Println(input)
		fmt.Printf("templateName = %s \n", name)
		collectorTemplate = template.Must(collectorTemplate.New(name).Parse(string(input)))
	}

	for _, tp := range collectorTemplate.Templates() {
		fmt.Printf("tp.Name = %s \n", tp.Name())
	}
	return collectorTemplate
}

func ConvertTemplateName(appDir string, path string) string {
	relPath := strings.Split(path, fmt.Sprintf("%s/views/", appDir))
	var result string
	if len(relPath) < 2 {
		log.Printf("relPath was less than 2, relPath was %s with len = %s \n", relPath, len(relPath))
	} else {
		result = relPath[1]
	}
	return result //HACK error prone
}

func FindPartialTemplates(appDir string) []string {
	//absDir, _ := filepath.Abs(appDir)
	//fmt.Printf("appDir = %s absVersion = %s \n", appDir, absDir)
	//appDirAbs = filepath.Abs(appDir)
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
