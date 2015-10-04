package command

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestRunHelp(t *testing.T) {
	new()
	cdRuntest(t)
	cmd := run()
	text := cmd.Help()
	if len(text) < 1 {
		t.Fail()
	}
	cdBack(t)
	cleanup(t)
}

func TestRunSynopsis(t *testing.T) {
	new()
	cdRuntest(t)
	cmd := run()
	text := cmd.Synopsis()
	if len(text) < 1 {
		t.Fail()
	}
	cdBack(t)
	cleanup(t)
}

func TestSplitJsonAndMarkdown(t *testing.T) {
	if jsonMap, _, err := SplitJsonAndMarkdown("../example/content/dogs/labrador_retriever.haiku"); err != nil {
		t.Error(err)
	} else {
		if jsonMap["breed"] != "Labrador Retriever" {
			t.Fail()
		}
	}
}

func TestSeparaterButNoJson(t *testing.T) {
	jsonMap, _, err := SplitJsonAndMarkdown("../example/content/dogs/jack_russel_terrier.haiku")
	if err != nil {
		t.Error(err)
	}
	//there shouldn't be any items in the json map since we didn't have any json
	if len(jsonMap) != 0 {
		t.Fail()
	}
}

func TestNoSeparaterNoJson(t *testing.T) {
	jsonMap, markdown, err := SplitJsonAndMarkdown("../example/content/dogs/airedale_terrier.haiku")
	if err != nil {
		t.Error(err)
	}
	//there shouldn't be any items in the json map since we didn't have any json
	if len(jsonMap) == 546 { //maybe a bit harsh, but I know that 546 is the right number for the airedale file
		t.Fail()
	}
	if len(markdown) < 1 {
		t.Fail()
	}
}

func TestCopyStyleDirectory(t *testing.T) {
	new()
	cdRuntest(t)
	run()

	currentDir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	styleDir := fmt.Sprintf("%s/style", currentDir)
	//see if the contents from example/style all exist in example/public/style
	if err := filepath.Walk(styleDir, WalkStyleDirectory); err != nil {
		t.Error(err)
	}
	cleanup(t)
}

func WalkStyleDirectory(path string, info os.FileInfo, err error) error {
	//assume path is something like workingdirectory/style/somelevel/something
	//then return error if runtest/public/style doesn't contain /somelevel/something for all style content
	re := regexp.MustCompile("/runtest/style/")
	//pull out whatever/style from the path; let the remaining be called relativeContentPath
	relativeContentPath := re.ReplaceAll([]byte(path), []byte("/runtest/public/style/"))
	fmt.Printf("path = %s \n relativeContentPath = %s \n", path, relativeContentPath)
	//check for runtest/style/relativeContentPath; if not there return a new Error
	if _, err := os.Lstat(string(relativeContentPath)); err != nil {
		return err
	}
	return nil
}

func TestTemplateParseAndInsertBasic(t *testing.T) {
	content := "Winter"
	htmlTemplate := `
<html>
<head>
</head>
<body>
{{ .Content }} is coming
</body>
</html>
`
	result, err := ParseAndInsert(content, htmlTemplate)
	if err != nil {
		t.Error(err)
	}
	//re := regexp.MustCompile( content )
	fmt.Printf("result = %s \n Contains = %s \n", result, strings.Contains(result, content))

	if !strings.Contains(result, content) { //!re.Match([]byte(result)) {

		t.Fatal("result did not contain expected content")
	}
}
