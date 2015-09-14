package main

import (
	"os"
	//"fmt"
	"github.com/fblecha/haiku/command"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available haiku commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{

		"create": func() (cli.Command, error) {
			return &command.CreateCommand{
				Name: "create",
				Ui:   ui,
			}, nil
		},

		"new": func() (cli.Command, error) {
			return &command.NewCommand{
				Name: "new",
				Ui:   ui,
			}, nil
		},

		"run": func() (cli.Command, error) {
			return &command.RunCommand{
				Name: "run",
				Ui:   ui,
			}, nil
		},

	}



}
