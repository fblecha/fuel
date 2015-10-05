package command

import (
	"encoding/json"
)

func isEmpty(s string) bool {
	return len(s) == 0
}

func parseJSON(jsonStr string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	jsonBytes := []byte(jsonStr)
	err := json.Unmarshal(jsonBytes, &dat);
	return dat, err
}
