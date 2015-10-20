package command

import (
  "testing"
  "html/template"
  "os"
  "fmt"
)


func TestFindPartialTemplates(t *testing.T) {
  expectations := []string {
    "views/dogs/menu2.partial.html",  //dog specific menu
    "views/menu.partial.html",  //site menu
  }
  wd, _ := os.Getwd()
  fmt.Printf("working directory = %s \n", wd)
  dir := "../example"
  results := FindPartialTemplates(dir)
  tests := make(map[string]bool)
  for _, val := range results {
    tests[val] = true
  }

  for _, expected := range expectations {
    if _, ok := tests[expected]; ok != true {
      t.Fatalf("expected %s but it was not found in %q with input of %q", expected, tests, results)
    }
  }
}

func XTestLoadPartialTemplates(t *testing.T) {
  tmpl := template.New("root")
  partials = FindPartialTemplates("../example")
  tmpl = LoadPartialTemplates(partials, tmpl)
  results := make(map[string]bool)
  for _, tp := range tmpl.Templates() {
    results[tp.Name()] = true
  }
  expectations := []string {
    "dogs/menu.partial.html",  //dog specific menu
    "menu.partial.html",  //site menu
  }
  for _, expected := range expectations {
    if _, ok := results[expected]; ok != true {
      t.Fatalf("expected %s in %q but it wasn't there", expected, results)
    }
  }
}
