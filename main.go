//fuel generates a static html site from a markdown dialect, while also putting metadata into a datastore.
package main

import (
	"github.com/fblecha/fuel/command"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	handleCommandOutsideOfProjectDir(args)

	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("fuel"),
	}

	exitStatus, err := cli.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}

func handleCommandOutsideOfProjectDir(args []string) {
	if len(args) > 0 {
		switch args[0] {
		case "new":
			//do nothing in the case of the "new" command
		default:
			if _, err := command.AreWeInProjectDir(); err != nil {
				log.Println(err)
			}
		}
	}
}
