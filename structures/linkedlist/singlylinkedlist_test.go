package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"github.com/vielendanke/go-dstructures/structures/api"
	"testing"
)

func TestSinglyLinkedList_Push(t *testing.T) {
	ll := prepareSinglyLinkedList()

	ll.Push(cInt(555))

	assert.True(t, ll.Contains(cInt(555)))
}

func TestSinglyLinkedList_Pop(t *testing.T) {
	ll := prepareSinglyLinkedList()

	val, isV := ll.Pop()

	assert.True(t, isV)
	assert.NotNil(t, val)
	assert.True(t, val.Equal(cInt(20)))
	assert.False(t, ll.Contains(cInt(20)))
}

func TestSinglyLinkedList_Contains(t *testing.T) {
	ll := prepareSinglyLinkedList()

	assert.True(t, ll.Contains(cInt(5)))
	assert.True(t, ll.Contains(cInt(6)))
	assert.True(t, ll.Contains(cInt(7)))
	assert.False(t, ll.Contains(cInt(33)))
	assert.False(t, ll.Contains(cInt(999)))
	assert.False(t, ll.Contains(cInt(51)))
}

func TestSinglyLinkedList_Get(t *testing.T) {
	ll := prepareSinglyLinkedList()

	g1, isG1 := ll.Get(0)
	g2, isG2 := ll.Get(7)
	g3, isG3 := ll.Get(15)
	g4, isG4 := ll.Get(555)

	assert.True(t, isG1)
	assert.True(t, isG2)
	assert.True(t, isG3)
	assert.False(t, isG4)
	assert.NotNil(t, g1)
	assert.NotNil(t, g2)
	assert.NotNil(t, g3)
	assert.Nil(t, g4)
}

func TestSinglyLinkedList_Set(t *testing.T) {
	ll := prepareSinglyLinkedList()

	ll.Set(0, cInt(500))

	assert.False(t, ll.Contains(cInt(1)))
	assert.True(t, ll.Contains(cInt(500)))
}

func TestSinglyLinkedList_Enqueue(t *testing.T) {
	ll := NewLinkedList()

	q := ll.(api.Queue)

	q.Enqueue(cInt(50))
	q.Enqueue(cInt(60))
	q.Enqueue(cInt(70))

	assert.True(t, ll.Contains(cInt(50)))
	assert.True(t, ll.Contains(cInt(60)))
	assert.True(t, ll.Contains(cInt(70)))
}

func TestSinglyLinkedList_Dequeue(t *testing.T) {
	ll := NewLinkedList()

	q := ll.(api.Queue)

	q.Enqueue(cInt(50))
	q.Enqueue(cInt(60))
	q.Enqueue(cInt(70))

	f1, _ := q.Dequeue()
	f2, _ := q.Dequeue()
	f3, _ := q.Dequeue()

	assert.True(t, f1.Equal(cInt(50)))
	assert.True(t, f2.Equal(cInt(60)))
	assert.True(t, f3.Equal(cInt(70)))
}

func TestSinglyLinkedList_Unshift(t *testing.T) {
	ll := prepareSinglyLinkedList()

	ll.Unshift(cInt(500))

	assert.True(t, ll.Contains(cInt(500)))
	assert.False(t, ll.Contains(cInt(501)))
}

func TestSinglyLinkedList_Shift(t *testing.T) {
	ll := prepareSinglyLinkedList()

	val, isV := ll.Shift()

	assert.True(t, isV)
	assert.True(t, cInt(1).Equal(val))
}

func TestSinglyLinkedList_Remove(t *testing.T) {
	ll := prepareSinglyLinkedList()

	r1, isR1 := ll.Remove(cInt(1))
	r2, isR2 := ll.Remove(cInt(14))
	r3, isR3 := ll.Remove(cInt(20))

	assert.True(t, isR1)
	assert.NotNil(t, r1)
	assert.False(t, ll.Contains(cInt(1)))
	assert.True(t, isR2)
	assert.NotNil(t, r2)
	assert.False(t, ll.Contains(cInt(14)))
	assert.True(t, isR3)
	assert.NotNil(t, r3)
	assert.False(t, ll.Contains(cInt(20)))
}

func TestSinglyLinkedList_RemoveIdx(t *testing.T) {
	ll := prepareSinglyLinkedList()

	r1, isR1 := ll.RemoveIdx(0)

	assert.True(t, isR1)
	assert.NotNil(t, r1)
	assert.False(t, ll.Contains(cInt(1)))
}

func TestSinglyLinkedList_Insert(t *testing.T) {
	ll := prepareSinglyLinkedList()

	is0 := ll.Insert(0, cInt(999))
	is10 := ll.Insert(10, cInt(998))
	g0, isG0 := ll.Get(0)
	g10, isG10 := ll.Get(10)

	assert.True(t, is0)
	assert.True(t, is10)
	assert.True(t, ll.Contains(cInt(999)))
	assert.True(t, ll.Contains(cInt(998)))
	assert.True(t, isG0)
	assert.True(t, g0.Equal(cInt(999)))
	assert.True(t, isG10)
	assert.True(t, g10.Equal(cInt(998)))
}

func TestSinglyLinkedList_Size(t *testing.T) {
	ll := prepareSinglyLinkedList()

	ll.Push(cInt(999))
	ll.Push(cInt(998))
	ll.Remove(cInt(999))

	assert.Equal(t, prepareSinglyLinkedList().Size()+1, ll.Size())
}
