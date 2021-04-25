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

	for i := 1; i < 21; i++ {
		hm.Put(cInt(i), "a" + strconv.Itoa(i))
	}
	return hm
}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
