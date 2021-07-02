package structures

import (
	"strconv"

	"github.com/vielendanke/go-dstructures/structures/api"
)

func prepareHashMap() api.Map {
	hm := NewHashMap(16)
	return populateMap(hm)
}

func prepareTreeMap() api.Map {
	tm, _ := NewTreeMap(16, func(leftKey interface{}, rightKey interface{}) bool {
		l := leftKey.(cInt)
		r := rightKey.(cInt)
		return r > l
	})
	return populateMap(tm)
}

func populateMap(m api.Map) api.Map {
	for i := 1; i < 21; i++ {
		m.Put(cInt(i), "a"+strconv.Itoa(i))
	}
	return m
}
