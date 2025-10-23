package utils

import (
	"maps"
	"slices"
)

type StringIntMap struct {
	data map[string]int
}

func (sim StringIntMap) GetKeys() []string {
	return slices.Collect(maps.Keys(sim.data))
}

func (sim StringIntMap) GetValue(s string) interface{} {
	return sim.data[s]
}

func (sim StringIntMap) KeyOk(s string) bool {
	_, ok := sim.data[s]
	return ok
}
