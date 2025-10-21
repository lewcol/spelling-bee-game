// Utility to generate pangrams.json

package main

import (
	"encoding/json"
	"maps"
	"os"
	"slices"
)

/*
func is_pangram(w string) bool {
	var unique_letters map[int32]int
	for _, c := range w {
		if _, ok := unique_letters[c]; ok {
			return false
		}
	}
	return true
}*/

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

func get_letters_if_pangram(w string) ([]rune, bool) {
	unique_letters := map[int32]int{}

	for _, c := range w {
		if c != 's' {
			unique_letters[c] = 1
		} else {
			return nil, false
		}
	}

	if len(unique_letters) != 7 {
		return nil, false
	}

	return slices.Collect(maps.Keys(unique_letters)), true
}

func main() {
	words, err := openJsonAsMap("../spelling-bee-game/server/wordlists/words_dictionary.json")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("../spelling-bee-game/server/wordlists/pangrams.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pangrams := map[string][]rune{}
	for k := range words {
		if letters, ok := get_letters_if_pangram(k); ok {
			pangrams[k] = letters
		}
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pangrams); err != nil {
		panic(err)
	}
}
