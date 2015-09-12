package main

import (
	//"log"
	//"encoding/json"
	"fmt"
	"os"
)

func appErrorMessage() error {
	currentDir, _ := os.Getwd()
	return fmt.Errorf(`
Hi there!  You likely wanted to execute this command in a haiku project directory.
For example, if you ran:

$ cd ~/code
$ haiku new todo

then ~/code/todo is your haiku project dir.  It'll have a polka/app haiku/config,
and the rest of the haiku generated files.

This time you ran haiku in %v
`, currentDir)
}

func areWeInProjectDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	//we are in the app root dir if we have both a ./app and
	//a ./config in the current working dir
	checkDirs := [...]string{
		fmt.Sprintf("%v/content", currentDir),
		fmt.Sprintf("%v/config", currentDir), //expand if needed
	}
	for _, checkDir := range checkDirs {
		if _, err := os.Stat(checkDir); err != nil {
			return "", appErrorMessage()
		}
	}
	//if we made it here, all the checkDirs exist, which means we should be good
	return currentDir, nil
}
