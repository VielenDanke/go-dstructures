package hashtable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeMap_Put(t *testing.T) {
	tm := prepareTreeMap()

	tm.Put(cInt(50), "fifty")
	tm.Put(cInt(51), "fifty")
	tm.Put(cInt(52), "fifty")

	assert.True(t, tm.Contains(cInt(50)))
	assert.True(t, tm.Contains(cInt(51)))
	assert.True(t, tm.Contains(cInt(52)))
	assert.False(t, tm.Contains(cInt(53)))
}

func TestTreeMap_Get(t *testing.T) {
	tm := prepareTreeMap()

	v17 := tm.Get(cInt(17))
	v18 := tm.Get(cInt(18))
	v19 := tm.Get(cInt(19))
	v51 := tm.Get(cInt(51))

	assert.Nil(t, v51)
	assert.NotNil(t, v17)
	assert.NotNil(t, v18)
	assert.NotNil(t, v19)
}

func TestTreeMap_Contains(t *testing.T) {
	tm := prepareTreeMap()

	b9 := tm.Contains(cInt(9))
	b10 := tm.Contains(cInt(10))
	b11 := tm.Contains(cInt(11))
	b32 := tm.Contains(cInt(32))

	assert.True(t, b9)
	assert.True(t, b10)
	assert.True(t, b11)
	assert.False(t, b32)
}

func TestTreeMap_Remove(t *testing.T) {
	tm := prepareTreeMap()

	r1, rv1 := tm.Remove(cInt(3))
	r2, rv2 := tm.Remove(cInt(4))
	r3, rv3 := tm.Remove(cInt(5))

	assert.NotNil(t, r1)
	assert.NotNil(t, rv1)
	assert.NotNil(t, r2)
	assert.NotNil(t, rv2)
	assert.NotNil(t, r3)
	assert.NotNil(t, rv3)

	assert.False(t, tm.Contains(cInt(3)))
	assert.False(t, tm.Contains(cInt(4)))
	assert.False(t, tm.Contains(cInt(5)))
	assert.Equal(t, prepareTreeMap().Size() - 3, tm.Size())
}

func TestTreeMap_Size(t *testing.T) {
	tm := prepareTreeMap()

	tm.Put(cInt(501), "five hundred and one")
	tm.Put(cInt(502), "five hundred and two")
	tm.Put(cInt(503), "five hundred and three")

	assert.Equal(t, prepareTreeMap().Size()+3, tm.Size())
}