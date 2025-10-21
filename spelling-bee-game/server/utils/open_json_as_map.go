package utils

import (
	"encoding/json"
	"os"
)

type WordMapType int

const (
	StringIntMap WordMapType = iota
	StringRuneSliceMap
)

func wordMapFactory(w WordMapType) any {
	switch w {
	case StringIntMap:
		return make(map[string]int)
	case StringRuneSliceMap:
		return make(map[string][]rune)
	default:
		return nil
	}
}

func openJsonAsMap(filename string, w WordMapType) (any, error) {
	file, err := os.Open(filename)
	if err != nil {

		return nil, err
	}
	defer file.Close()

	words := wordMapFactory(w)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&words)
	if err != nil {
		return nil, err
	}
	return words, nil
}
