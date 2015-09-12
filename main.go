package main

import (
	"github.com/mitchellh/cli"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	var err error

	handleCommandOutsideOfProjectDir(args)

	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("polka"),
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
			if _, err := areWeInProjectDir(); err == nil {
				//otherwise, make sure we have the AppConfig setup
			} else {
				log.Println(err)
			}
		}
	}
}
