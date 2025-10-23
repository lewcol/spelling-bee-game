package utils

import (
	"encoding/json"
	"os"
)

type WordMapType int

const (
	StringIntMap WordMapType = iota
	StringMapRuneIntMap
)

type WordMap interface {
	Exists(key string) bool
	GetValueFromKey(key string) (interface{}, bool)
	GetKeys() []string
	GetType() string
}

func wordMapFactory(w WordMapType) any {
	switch w {
	case StringIntMap:
		return &map[string]int{}
	case StringMapRuneIntMap:
		return &map[string]map[rune]int{}
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
	err = decoder.Decode(words)
	if err != nil {
		return nil, err
	}
	return words, nil
}
