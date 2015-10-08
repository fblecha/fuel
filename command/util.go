package command

import (
	"encoding/json"
	"errors"
	"path/filepath"
	"strings"
)

func isEmpty(s string) bool {
	return len(s) == 0
}

func parseJSON(jsonStr string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	jsonBytes := []byte(jsonStr)
	err := json.Unmarshal(jsonBytes, &dat)
	return dat, err
}

func getFilenameMinusExtension(path string) string {
	filename := filepath.Base(path)
	extensionIndex := strings.LastIndex(filename, ".")
	newFilename := filename[0:extensionIndex]
	return newFilename
}

func GetRelativePath(appDir string, path string) (string, error) {
	return "", errors.New("blah")
}

func PathToDirs(path string) []string {

	return strings.Split(path, "/")
}

func Reverse(things []string) []string {
	//TODO need to figure this out when I'm not on the train
	//	return []string(sort.Reverse(sort.StringSlice(things)))

	// //TODO more efficient way to do this
	newThings := make([]string, len(things))
	for i := range things {
		j := len(things) - i - 1
		newThings[i] = things[j]
	}
	return newThings
}

func LoadFileAsString(path string) (string, error) {
	return "", errors.New("blah")
}
