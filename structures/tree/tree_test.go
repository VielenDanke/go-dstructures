package tree

import (
	"github.com/vielendanke/go-dstructures/structures/api"
	"os"
	"testing"
)

type cInt int

func (c cInt) Equal(v interface{}) bool {
	return c == v.(cInt)
}

func (c cInt) Hash() int {
	return int(c)
}

func prepareRbTreeMap() api.Map {
	rb := NewRBTree(func(leftKey interface{}, rightKey interface{}) bool {
		l := leftKey.(cInt)
		r := rightKey.(cInt)
		return l > r
	})
	rb.Put(cInt(1), "qwe")
	rb.Put(cInt(2), "asd")
	rb.Put(cInt(3), "zxc")
	rb.Put(cInt(4), "rty")
	rb.Put(cInt(5), "fgh")
	rb.Put(cInt(6), "qwe")
	rb.Put(cInt(7), "asd")
	rb.Put(cInt(8), "zxc")
	rb.Put(cInt(9), "rty")
	rb.Put(cInt(10), "fgh")
	rb.Put(cInt(11), "qwe")
	rb.Put(cInt(12), "asd")
	rb.Put(cInt(13), "zxc")
	rb.Put(cInt(14), "rty")
	rb.Put(cInt(15), "fgh")
	rb.Put(cInt(16), "qwe")
	rb.Put(cInt(17), "asd")
	rb.Put(cInt(18), "zxc")
	rb.Put(cInt(19), "rty")
	rb.Put(cInt(20), "fgh")

	return rb
}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
