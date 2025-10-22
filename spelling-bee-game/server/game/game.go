package game

type Game interface {
	Score() int
	Word() string
	Letters() map[rune]int
	Centre() rune
	Submit(s string) (string, int)
	New() Game
}
