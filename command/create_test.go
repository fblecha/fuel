package command

import (
	"github.com/mitchellh/cli"
	"os"
	"testing"
)

func TestCreateHelp(t *testing.T) {
	ui := &cli.BasicUi{Writer: os.Stdout}
	//args := []string{""}
	cmd := CreateCommand{
		Name: "new",
		Ui:   ui,
	}
	text := cmd.Help()
	if len(text) < 1 {
		t.Fail()
	}
}

func TestCreateSynopsis(t *testing.T) {
	ui := &cli.BasicUi{Writer: os.Stdout}
	//args := []string{""}
	cmd := CreateCommand{
		Name: "new",
		Ui:   ui,
	}
	text := cmd.Synopsis()
	if len(text) < 1 {
		t.Fail()
	}
}

func TestCreateContent(t *testing.T) {

}
