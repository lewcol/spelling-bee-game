package utils

import (
	"fmt"
	"maps"
	"math/rand/v2"
	"slices"
	"sync"
)

var (
	once     sync.Once
	instance *dictionary
)

type Dictionary interface {
	IsWord(word string) bool
	IsPangram(word string) bool
	GetInstance() *dictionary
	GetWordAndLetters() (string, map[rune]int, rune)
}

type dictionary struct {
	words    *map[string]int
	pangrams *map[string][]rune
}

func (d dictionary) IsWord(word string) bool {
	_, ok := (*d.words)[word]
	return ok
}

func (d dictionary) IsPangram(word string) bool {
	_, ok := (*d.pangrams)[word]
	return ok
}

func (d dictionary) GetWordAndLetters() (string, map[rune]int, rune) {
	wordSlice := slices.Collect(maps.Keys(*d.pangrams))
	length := len(wordSlice)
	word := wordSlice[rand.IntN(length)]
	letters := make(map[rune]int)
	for _, letter := range word {
		letters[letter] = 1
	}
	centre := rune(word[rand.IntN(len(word))])
	return word, letters, centre
}

func (d dictionary) GetInstance() *dictionary {
	once.Do(func() {
		words, err := openJsonAsMap("../wordlists/words_dictionary.json", StringIntMap)
		if err != nil {
			fmt.Println("Error loading dictionary: ", err)
			return
		}

		pangrams, err := openJsonAsMap("../wordlists/pangrams.json", RuneIntMap)
		if err != nil {
			fmt.Println("Error loading dictionary: ", err)
			return
		}

		instance = &dictionary{
			words:    words.(*map[string]int),
			pangrams: pangrams.(*map[string][]rune),
		}
	})
	return instance
}
