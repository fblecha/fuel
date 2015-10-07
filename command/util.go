package command

import (
	"encoding/json"
	"errors"
	"path/filepath"
	//"sort"
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

func PathToDirs(relativePath string) []string {
	blah := []string{"blah"}
	return blah
}

func Reverse(things []string) []string {
//	return []string(sort.Reverse(sort.StringSlice(things)))

	// //TODO more efficient way to do this
	newThings := make([]string, len(things) )
	for i := range things {
		j := len(things)-i
		newThings[i] = things[j]
	}
	return newThings
}

func LoadFileAsString(path string) (string, error) {
	return "", errors.New("blah")
}
