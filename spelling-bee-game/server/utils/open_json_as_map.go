package utils

import (
	"encoding/json"
	"os"
)

func openJsonAsMap(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	words := map[string]int{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&words)
	if err != nil {
		return nil, err
	}
	return words, nil
}
