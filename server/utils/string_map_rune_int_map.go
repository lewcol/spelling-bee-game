package utils

import (
	"maps"
	"slices"
)

type StringMapRuneIntMap struct {
	data map[string]map[rune]int
}

func (smr StringMapRuneIntMap) GetKeys() []string {
	return slices.Collect(maps.Keys(smr.data))
}

func (smr StringMapRuneIntMap) GetValue(s string) interface{} {
	return smr.data[s]
}

func (smr StringMapRuneIntMap) KeyOk(s string) bool {
	_, ok := smr.data[s]
	return ok
}
