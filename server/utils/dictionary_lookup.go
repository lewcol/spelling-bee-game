package utils

import (
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
	GetWordAndLetters() (string, map[rune]int, rune)
}

type dictionary struct {
	words    WordMap
	pangrams WordMap
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

func GetInstance() Dictionary {
	once.Do(func() {
		words, err := openJsonAsMap("./wordlists/words_dictionary.json", StringIntMapType)
		if err != nil {
			panic("Error loading dictionary: " + err.Error())
			return
		}

		pangrams, err := openJsonAsMap("./wordlists/pangrams.json", StringMapRuneIntMapType)
		if err != nil {
			panic("Error loading dictionary: " + err.Error())
			return
		}

		instance = &dictionary{
			words:    words,
			pangrams: pangrams,
		}
	})
	return instance
}
