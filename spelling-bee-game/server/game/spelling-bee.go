package game

import (
	"maps"
	"slices"
	"spelling-bee-game/server/utils"
	"strconv"
	"strings"
)

type game struct {
	score   int
	word    string
	letters map[rune]int
	centre  rune
	guessed map[string]int
}

func (g game) Score() int { return g.score }

func (g game) Word() string { return g.word }

func (g game) Letters() map[rune]int { return g.letters }

func (g game) PrintableLettersWithCentre() string {
	var letters []string
	for k := range maps.Keys(g.Letters()) {
		if k != g.Centre() {
			letters = append(letters, string(k))
		}
	}
	middle := len(letters) / 2
	letters = slices.Insert(letters, middle, "["+string(g.Centre())+"]")
	return strings.Join(letters, " ")
}

func (g game) Centre() rune { return g.centre }

func (g game) Guessed() map[string]int { return g.guessed }

func (g game) Submit(s string) (string, int) {
	// reject old guesses
	if _, ok := g.guessed[s]; ok {
		return "Already guessed.", g.Score()
	}

	// reject words of less than 4 letters
	i := len(s)
	if i < 4 {
		return "Word must be at least 4 letters!", g.Score()
	}

	// check if word
	if !utils.GetInstance().IsWord(s) {
		return "Not a valid word.", g.Score()
	}

	// perform tests iterating over letters in word
	// reject words missing centre letter or with letters not in word
	hasCentre := false
	var notIn []rune
	for _, letter := range s {
		if _, ok := g.Letters()[letter]; !ok {
			notIn = append(notIn, letter)
		}
		if letter == g.Centre() {
			hasCentre = true
		}
	}
	if len(notIn) != 0 {
		return "Letters " + strings.Trim(string(notIn), "[]") + " not in list.", g.Score()
	}

	if !hasCentre {
		return "Missing Centre Letter!", g.Score()
	}

	// the word is valid
	// add to previous guesses
	g.guessed[s] = 1

	// determine pangram
	isPangram := true
	for _, letter := range slices.Collect(maps.Keys(g.Letters())) {
		if !strings.ContainsRune(s, letter) {
			isPangram = false
			break
		}
	}
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
	return "Valid word scoring " + strconv.Itoa(i) + " points.", g.Score()
}

func New() Game {
	dict := utils.GetInstance()
	word, letters, centre := dict.GetWordAndLetters()
	return game{
		score:   0,
		word:    word,
		letters: letters,
		centre:  centre,
		guessed: map[string]int{},
	}
}
