package set

import (
	"github.com/vielendanke/go-dstructures/structures/api"
	"os"
	"strconv"
	"testing"
)

type cStr string

func (c cStr) Equal(p interface{}) bool {
	incStr := p.(cStr)
	return string(c) == string(incStr)
}

func (c cStr) Hash() int {
	i := 0
	for _, v := range c {
		i += int(v)
	}
	return i
}

func prepareHashSet() api.Set {
	ts := NewHashSet()
	fillWithData(ts)
	return ts
}

func prepareTreeSet() api.Set {
	ts, _ := NewTreeSet(func(left interface{}, right interface{}) bool {
		l := left.(cStr)
		r := right.(cStr)
		return len(string(l)) > len(string(r))
	})
	fillWithData(ts)
	return ts
}

func fillWithData(s api.Set) {
	for i := 1; i < 21; i++ {
		is := strconv.Itoa(i)
		s.Add(cStr("abc" + is))
	}
}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
