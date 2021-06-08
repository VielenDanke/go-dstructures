package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue_Enqueue(t *testing.T) {
	q := NewPriorityQueue(priorityCalc)

	q.Enqueue(cInt(1))
	q.Enqueue(cInt(2))
	q.Enqueue(cInt(3))

	assert.Equal(t, 3, q.Size())
}

func TestPriorityQueue_Dequeue(t *testing.T) {
	q := preparePriorityQueue()

	elem1, isE1 := q.Dequeue()
	elem2, isE2 := q.Dequeue()
	elem3, isE3 := q.Dequeue()

	assert.True(t, isE1)
	assert.True(t, isE2)
	assert.True(t, isE3)
	assert.NotNil(t, elem1)
	assert.NotNil(t, elem2)
	assert.NotNil(t, elem3)
	assert.True(t, cInt(1).Equal(elem1))
	assert.True(t, cInt(2).Equal(elem2))
	assert.True(t, cInt(3).Equal(elem3))
}

func TestPriorityQueue_Size(t *testing.T) {
	q := preparePriorityQueue()

	q.Enqueue(cInt(325))
	q.Enqueue(cInt(421))
	q.Dequeue()

	assert.Equal(t, preparePriorityQueue().Size()+1, q.Size())
}
