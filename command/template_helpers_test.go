package command

import (
	"testing"
	//"html/template"
	//"os"
	//"fmt"
	"path/filepath"
)

func TestConvertTemplateName(t *testing.T) {
	appDir, err := filepath.Abs("../example")
	if err != nil {
		t.Fatal(err)
	}
	path := "../example/views/dogs/layout.html"
	absPath, _ := filepath.Abs(path)
	if err != nil {
		t.Fatal(err)
	}
	expectedName := "dogs/layout.html"
	actualName := ConvertTemplateName(appDir, absPath)
	if actualName != expectedName {
		t.Fatalf("Expected '%s' but actually got '%s' ", expectedName, actualName)
	}
}

// func TestFindPartialTemplates(t *testing.T) {
//   wd, _ := os.Getwd()
//   os.Chdir("../example")
//   expectations := []string {
//     "views/dogs/menu2.partial.html",  //dog specific menu
//     "/views/menu.partial.html",  //site menu
//   }
//   dir, err := filepath.Abs("../example")
//   if err != nil {
//     t.Fatal(err)
//   }
//   results := FindPartialTemplates(dir)
//
//
//   tests := make(map[string]bool)
//
//
//
//   for _, val := range results {
//     tests[val] = true
//   }
//
//   for _, expected := range expectations {
//     if _, ok := tests[expected]; ok != true {
//       t.Fatalf("expected %s \n but it was not found in %q \n with input of %q \n", expected, tests, results)
//     }
//   }
//   os.Chdir(wd)
// }

// func TestLoadPartialTemplates(t *testing.T) {
//   tmpl := template.New("root")
//   partials = FindPartialTemplates("../example")
//   fmt.Printf("found partials = %q \n", partials)
//   tmpl = LoadPartialTemplates(partials, tmpl)
//   results := make(map[string]bool)
//   for _, tp := range tmpl.Templates() {
//     results[tp.Name()] = true
//   }
//   expectations := []string {
//     "views/dogs/menu.partial.html",  //dog specific menu
//     "views/menu.partial.html",  //site menu
//   }
//   for _, expected := range expectations {
//     if _, ok := results[expected]; ok != true {
//       t.Fatalf("expected %s in %q but it wasn't there", expected, results)
//     }
//   }
// }
