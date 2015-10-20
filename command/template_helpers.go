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
  originalDir, _ := os.Getwd()  //HACK - check error
  os.Chdir("../example")  //HACK - check error

  walker := func (path string, f os.FileInfo, err error) error {
  	switch filepath.Ext(path) {
  	case ".html":
      relPath, _ := filepath.Rel(appDir, path)
  		fmt.Println( relPath )
  		partials = append(partials, path)
  		return nil
  	}
  	return nil
  }

	root := fmt.Sprintf("%s/views", ".")
	//for appDir/views/partials, load all files in that directory into partials
	filepath.Walk(root, walker)
  os.Chdir(originalDir) //HACK - check error
  return partials
}
