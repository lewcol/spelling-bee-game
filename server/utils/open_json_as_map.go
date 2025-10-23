package utils

import (
	"encoding/json"
	"os"
)

type WordMapType int

const (
	StringIntMapType WordMapType = iota
	StringMapRuneIntMapType
)

func wordMapFactory(w WordMapType) WordMap {
	switch w {
	case StringIntMapType:
		return &StringIntMap{data: map[string]int{}}
	case StringMapRuneIntMapType:
		return &StringMapRuneIntMap{data: map[string]map[rune]int{}}
	default:
		return nil
	}
}

func openJsonAsMap(filename string, w WordMapType) (WordMap, error) {
	file, err := os.Open(filename)
	if err != nil {

		return nil, err
	}
	defer file.Close()

	words := wordMapFactory(w)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(words)
	if err != nil {
		return nil, err
	}
	return words, nil
}
