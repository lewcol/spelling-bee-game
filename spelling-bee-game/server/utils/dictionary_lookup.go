package utils

import (
	"fmt"
	"sync"
)

type Dictionary interface {
	IsWord(word string) bool
	IsPangram(word string) bool
	GetInstance() *dictionary
}

type dictionary struct {
	words    *map[string]int
	pangrams *map[string][]rune
}

var once sync.Once
var instance *dictionary

func (d dictionary) IsWord(word string) bool {
	//TODO implement me
	panic("implement me")
}

func (d dictionary) IsPangram(word string) bool {
	//TODO implement me
	panic("implement me")
}

func (d dictionary) GetInstance() *dictionary {
	if instance == nil {
		once.Do(func() {
			words, err := openJsonAsMap("../wordlists/words_dictionary.json", StringIntMap)
			if err != nil {
				fmt.Println(err)
				instance = nil
			}

			pangrams, err := openJsonAsMap("../wordlists/pangrams.json", StringRuneSliceMap)
			if err != nil {
				fmt.Println(err)
				instance = nil
			}

			instance = &dictionary{
				words:    words.(*map[string]int),
				pangrams: pangrams.(*map[string][]rune),
			}
		})
	}
	return instance
}
