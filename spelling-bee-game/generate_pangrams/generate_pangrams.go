// Utility to generate pangrams.json

package main

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

func getLettersIfPangram(w string) (map[rune]int, bool) {
	uniqueLetters := map[rune]int{}

	for _, c := range w {
		if c != 's' {
			uniqueLetters[c] = 1
		} else {
			return nil, false
		}
	}

	if len(uniqueLetters) != 7 {
		return nil, false
	}

	return uniqueLetters, true
}

func main() {
	words, err := openJsonAsMap("../server/wordlists/words_dictionary.json")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("../server/wordlists/pangrams.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pangrams := map[string]map[rune]int{}
	for k := range words {
		if letters, ok := getLettersIfPangram(k); ok {
			pangrams[k] = letters
		}
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pangrams); err != nil {
		panic(err)
	}
}
