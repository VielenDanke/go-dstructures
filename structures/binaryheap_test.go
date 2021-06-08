package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryHeap_Insert(t *testing.T) {
	bh := prepareBinaryHeap()

	bh.Insert(503)
	bh.Insert(292)
	bh.Insert(111)

	assert.Equal(t, prepareBinaryHeap().Size()+3, bh.Size())
}

func TestBinaryHeap_RemoveMax(t *testing.T) {
	bh := prepareBinaryHeap()

	max1, isMax1 := bh.RemoveMax()
	max2, isMax2 := bh.RemoveMax()
	max3, isMax3 := bh.RemoveMax()
	assert.True(t, isMax1)
	assert.True(t, isMax2)
	assert.True(t, isMax3)
	assert.Equal(t, 20, max1)
	assert.Equal(t, 19, max2)
	assert.Equal(t, 18, max3)
}

func TestBinaryHeap_Contains(t *testing.T) {
	bh := prepareBinaryHeap()

	c1 := bh.Contains(18)
	c2 := bh.Contains(11)
	c3 := bh.Contains(3)
	c4 := bh.Contains(201)

	assert.True(t, c1)
	assert.True(t, c2)
	assert.True(t, c3)
	assert.False(t, c4)
}

func TestBinaryHeap_Size(t *testing.T) {
	bh := prepareBinaryHeap()

	bh.Insert(501)
	bh.Insert(303)
	bh.RemoveMax()

	assert.Equal(t, prepareBinaryHeap().Size()+1, bh.Size())
}
