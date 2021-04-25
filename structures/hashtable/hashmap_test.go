package hashtable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestHashMap_Put(t *testing.T) {
	hm := prepareHashMap()

	v5 := hm.Get(cInt(5))
	v6 := hm.Get(cInt(6))
	v7 := hm.Get(cInt(7))

	assert.NotNil(t, v5)
	assert.NotNil(t, v6)
	assert.NotNil(t, v7)
}

func TestHashMap_Contains(t *testing.T) {
	hm := prepareHashMap()

	c5 := hm.Contains(cInt(5))
	c6 := hm.Contains(cInt(6))
	c7 := hm.Contains(cInt(7))

	assert.True(t, c5)
	assert.True(t, c6)
	assert.True(t, c7)
}

func TestHashMap_Get(t *testing.T) {
	hm := prepareHashMap()

	g5 := hm.Get(cInt(5))
	g6 := hm.Get(cInt(6))
	g7 := hm.Get(cInt(7))

	assert.NotNil(t, g5)
	assert.NotNil(t, g6)
	assert.NotNil(t, g7)
}

func TestHashMap_Remove(t *testing.T) {
	hm := prepareHashMap()

	key, val := hm.Remove(cInt(5))
	gv := hm.Get(cInt(5))

	assert.NotNil(t, key)
	assert.NotNil(t, val)
	assert.Nil(t, gv)
}
