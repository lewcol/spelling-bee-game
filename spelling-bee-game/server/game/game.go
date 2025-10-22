package game

type Game interface {
	Score() int
	Word() string
	Letters() map[rune]int
	Centre() rune
	Guessed() map[string]int
	Submit(s string) (string, int)
}
