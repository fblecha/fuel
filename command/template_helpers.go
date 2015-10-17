package command

import (
  "fmt"
  "os"
  "path/filepath"
	//"regexp"
	//"strings"
	//"bufio"
	//"bytes"
	//"github.com/termie/go-shutil"
	"html/template"
)

func LoadPartialTemplates(partialTemplatePaths []string, parent *template.Template) *template.Template {
  // for i, path := range partialTemplatePaths {
  //
  // }
  return parent
}

func FindPartialTemplates(appDir string) []string {
  var partials []string

  walker := func (path string, f os.FileInfo, err error) error {
  	switch filepath.Ext(path) {
  	case ".html":
  		fmt.Println(path)
  		partials = append(partials, path)
  		return nil
  	}
  	return nil
  }

	root := fmt.Sprintf("%s/views/partials", appDir)
	//for appDir/views/partials, load all files in that directory into partials
	filepath.Walk(root, walker)
  return partials
}
