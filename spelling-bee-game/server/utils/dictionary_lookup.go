package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type WordMap interface {
	IsWord(word string) (bool, string)
	IsPangram(word string) (bool, string)
	GetLetters() []string
}

type wordMap struct {
	validWords map[string]int
	pangrams   map[string]int
}

func (w wordMap) IsWord(word string) (bool, string) {
	//TODO implement me
	panic("implement me")
}

func (w wordMap) IsPangram(word string) (bool, string) {

}

func (w wordMap) GetLetters() []string {
	//TODO implement me
	panic("implement me")
}

func main() {
	if words, err := openJsonAsMap("../wordlists/wordlist.txt"); err != nil {
		fmt.Println(err)
		return
	}

	if pangrams, err := openJsonAsMap("../wordlists/pangrams.txt"); err != nil {
		fmt.Println(err)
		return
	}
}
