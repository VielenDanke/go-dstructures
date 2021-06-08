package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeSet_Add(t *testing.T) {
	ts := prepareTreeSet()
	testData := "abcbca"

	ts.Add(cStr(testData))

	assert.Equal(t, prepareTreeSet().Size()+1, ts.Size())
	assert.True(t, ts.Contains(cStr(testData)))
	assert.False(t, ts.Contains(cStr(testData + "abc")))
}

func TestTreeSet_Contains(t *testing.T) {
	ts := prepareTreeSet()

	c2 := ts.Contains(cStr("abc2"))
	c3 := ts.Contains(cStr("abc3"))
	c4 := ts.Contains(cStr("abc4"))
	c100000 := ts.Contains(cStr("abc100000"))

	assert.True(t, c2)
	assert.True(t, c3)
	assert.True(t, c4)
	assert.False(t, c100000)
}

func TestTreeSet_Remove(t *testing.T) {
	ts := prepareTreeSet()

	r2 := ts.Remove(cStr("abc2"))
	r3 := ts.Remove(cStr("abc3"))
	r4 := ts.Remove(cStr("abc4"))
	r1000 := ts.Remove(cStr("abc1000"))

	c2 := ts.Contains(cStr("abc2"))
	c3 := ts.Contains(cStr("abc3"))
	c4 := ts.Contains(cStr("abc4"))

	assert.NotNil(t, r2)
	assert.NotNil(t, r3)
	assert.NotNil(t, r4)
	assert.Nil(t, r1000)
	assert.False(t, c2)
	assert.False(t, c3)
	assert.False(t, c4)
}

func TestTreeSet_Size(t *testing.T) {
	ts := prepareTreeSet()

	ts.Add(cStr("bla"))
	ts.Add(cStr("uuu"))
	ts.Add(cStr("abc3"))

	assert.Equal(t, prepareTreeSet().Size() + 2, ts.Size())
}

func TestTreeSet_ToArray(t *testing.T) {
	ts := prepareTreeSet()

	arr := ts.ToArray()

	assert.NotNil(t, arr)
	assert.Equal(t, prepareTreeSet().Size(), len(arr))
}