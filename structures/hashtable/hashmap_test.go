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

	key5, val5 := hm.Remove(cInt(5))
	key6, val6 := hm.Remove(cInt(6))
	key7, val7 := hm.Remove(cInt(7))

	assert.NotNil(t, key5)
	assert.NotNil(t, key6)
	assert.NotNil(t, key7)
	assert.NotNil(t, val5)
	assert.NotNil(t, val6)
	assert.NotNil(t, val7)
	assert.False(t, hm.Contains(key5))
	assert.False(t, hm.Contains(key6))
	assert.False(t, hm.Contains(key7))
	assert.Equal(t, prepareHashMap().Size() - 3, hm.Size())
}

func TestHashMap_PutNil(t *testing.T) {
	hm := prepareHashMap()

	hm.Put(nil, "abc")

	v := hm.Get(nil)

	assert.Nil(t, v)
}

func TestHashMap_Size(t *testing.T) {
	hm := prepareHashMap()

	hm.Put(cInt(503), "new value")
	hm.Put(cInt(555), "good value")

	assert.Equal(t, prepareHashMap().Size()+2, hm.Size())
}
