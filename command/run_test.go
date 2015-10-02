package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {

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
	//use the haiku run command to make a new test blog
	ui := &cli.BasicUi{Writer: os.Stdout}
	args := []string{"runtest"}
	newCmd := NewCommand{
		Name: "new",
		Ui:   ui,
	}
	newCmd.Run(args)

	wd, _ := os.Getwd()
	testdir := fmt.Sprintf("%s/runtest", wd)
	if err := os.Chdir(testdir); err != nil {
		t.Error(err)
	}

	runCmd := RunCommand{
		Name: "run",
		Ui:   ui,
	}
	runCmd.Run(args)

	currentDir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	styleDir := fmt.Sprintf("%s/style", currentDir)
	//see if the contents from example/style all exist in example/public/style
	if err := filepath.Walk(styleDir, WalkStyleDirectory); err != nil {
		t.Error(err)
	}
	CleanupRunTestDirectory(currentDir, t)
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

func CleanupRunTestDirectory(currentDir string, t *testing.T) {
	//cleanup by removeing all the directories we created (hopefully)
	dir := fmt.Sprintf("%s/runtest", currentDir)
	if err := os.RemoveAll(dir); err != nil {
		t.Error(err)
	}
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
	fmt.Printf("result = %s \n Contains = %s \n", result, strings.Contains(result,content) )

	if !strings.Contains(result, content) { //!re.Match([]byte(result)) {

		t.Fatal("result did not contain expected content")
	}
}
