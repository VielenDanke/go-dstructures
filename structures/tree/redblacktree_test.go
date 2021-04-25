package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRbTree_String(t *testing.T) {
	rb := prepareRbTreeMap()

	assert.NotEmpty(t, fmt.Sprintf("%v", rb))
	assert.Contains(t, fmt.Sprintf("%v", rb), "1:qwe")
	assert.Contains(t, fmt.Sprintf("%v", rb), "2:asd")
	assert.Contains(t, fmt.Sprintf("%v", rb), "3:zxc")
}

func TestRbTree_Size(t *testing.T) {
	rb := prepareRbTreeMap()

	rb.Put(cInt(21), "abc")
	rb.Put(cInt(22), "abc")
	rb.Put(cInt(23), "abc")
	rb.Put(cInt(24), "abc")
	rb.Put(cInt(25), "abc")

	assert.Equal(t, 25, rb.Size())
}

func TestRbTree_Put(t *testing.T) {
	rb := prepareRbTreeMap()

	assert.Equal(t, 20, rb.Size())
	assert.True(t, rb.Contains(cInt(11)))
	assert.True(t, rb.Contains(cInt(1)))
	assert.True(t, rb.Contains(cInt(19)))
	assert.False(t, rb.Contains(cInt(21)))
}

func TestRbTree_Get(t *testing.T) {
	rb := prepareRbTreeMap()

	val15 := rb.Get(cInt(15))
	val11 := rb.Get(cInt(11))
	val21 := rb.Get(cInt(21))

	assert.NotNil(t, val15)
	assert.NotNil(t, val11)
	assert.Nil(t, val21)

}

func TestRbTree_Remove(t *testing.T) {
	rb := prepareRbTreeMap()

	key, val := rb.Remove(cInt(15))

	assert.NotNil(t, key)
	assert.NotNil(t, val)
	assert.False(t, rb.Contains(cInt(15)))
	assert.Equal(t, 19, rb.Size())
}

func TestRbTree_PutNil(t *testing.T) {
	rb := prepareRbTreeMap()

	rb.Put(nil, "abc")

	v := rb.Get(nil)

	assert.Nil(t, v)
}