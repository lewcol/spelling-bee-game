package utils

type WordMap interface {
	GetKeys() []string
	GetValue(s string) interface{}
	KeyOk(s string) bool
}
