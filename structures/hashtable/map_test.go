package hashtable

import (
	"github.com/vielendanke/go-dstructures/structures/api"
	"os"
	"strconv"
	"testing"
)

type cInt int

func (c cInt) Equal(p interface{}) bool {
	return c.Hash() == p.(cInt).Hash()
}

func (c cInt) Hash() int {
	return int(c)
}

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
		m.Put(cInt(i), "a" + strconv.Itoa(i))
	}
	return m
}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
