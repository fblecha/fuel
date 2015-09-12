package main

import (
	"os"
	//"fmt"
	"github.com/fblecha/haiku/command"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available Polka commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{

		"new": func() (cli.Command, error) {
			return &command.NewCommand{
				Name: "new",
				Ui:   ui,
			}, nil
		},

	}

}
