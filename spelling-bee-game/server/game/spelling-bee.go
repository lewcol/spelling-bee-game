package game

import (
	"spelling-bee-game/server/utils"
	"strings"
)

type game struct {
	score   int
	word    string
	letters map[rune]int
	centre  rune
}

func (g game) Score() int { return g.score }

func (g game) Word() string { return g.word }

func (g game) Letters() map[rune]int { return g.letters }

func (g game) Centre() rune { return g.centre }

func (g game) Submit(s string) (string, int) {
	// reject words of less than 4 letters
	i := len(s)
	if i < 4 {
		return "Word be at least 4 letters!", g.Score()
	}

	// perform tests iterating over letters in word
	// reject words missing centre letter or with letters not in word
	// determine pangram
	hasCentre := false
	isPangram := true
	var notIn []rune
	for _, letter := range s {
		if _, ok := g.Letters()[letter]; !ok {
			notIn = append(notIn, letter)
		}
		if letter == g.Centre() {
			hasCentre = true
		}
		if _, ok := g.Letters()[letter]; !ok {
			isPangram = false
		}
	}
	if len(notIn) != 0 {
		return "Letters " + strings.Trim(string(notIn), "[]") + " not in list", g.Score()
	}

	if !hasCentre {
		return "Missing Centre Letter!", g.Score()
	}

	// the word is valid
	// update scores for valid words
	if i == 4 {
		g.score += 1
	} else {
		g.score += i
	}

	// apply pangram bonus
	if isPangram {
		g.score += 7
		return "Pangram!", g.Score()
	}

	// no pangram bonus
	return "", g.Score()
}

func (g game) New() Game {
	var dict utils.Dictionary
	dict = dict.GetInstance()
	word, letters, centre := dict.GetWordAndLetters()
	return game{
		score:   0,
		word:    word,
		letters: letters,
		centre:  centre,
	}
}
