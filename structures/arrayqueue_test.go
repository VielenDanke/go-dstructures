package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayQueue_Enqueue(t *testing.T) {
	q := NewArrayQueue()

	q.Enqueue(cInt(1))
	q.Enqueue(cInt(2))
	q.Enqueue(cInt(3))

	assert.Equal(t, 3, q.Size())
}

func TestArrayQueue_Dequeue(t *testing.T) {
	q := prepareArrayQueue()

	elem1, isE1 := q.Dequeue()
	elem2, isE2 := q.Dequeue()
	elem3, isE3 := q.Dequeue()

	assert.True(t, isE1)
	assert.True(t, isE2)
	assert.True(t, isE3)
	assert.NotNil(t, elem1)
	assert.NotNil(t, elem2)
	assert.NotNil(t, elem3)
	assert.True(t, cInt(21).Equal(elem1))
	assert.True(t, cInt(20).Equal(elem2))
	assert.True(t, cInt(19).Equal(elem3))
}

func TestArrayQueue_Size(t *testing.T) {
	q := prepareArrayQueue()

	q.Enqueue(cInt(500))
	q.Enqueue(cInt(299))
	q.Dequeue()

	assert.Equal(t, prepareArrayQueue().Size()+1, q.Size())
}
