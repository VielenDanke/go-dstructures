package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vielendanke/go-dstructures/structures/api"
)

func TestLinkedList_Push(t *testing.T) {
	ll := prepareLinkedList()

	ll.Push(cInt(555))

	assert.True(t, ll.Contains(cInt(555)))
}

func TestLinkedList_Pop(t *testing.T) {
	ll := prepareLinkedList()

	val, isV := ll.Pop()

	assert.True(t, isV)
	assert.NotNil(t, val)
	assert.True(t, val.Equal(cInt(20)))
	assert.False(t, ll.Contains(cInt(20)))
}

func TestLinkedList_Contains(t *testing.T) {
	ll := prepareLinkedList()

	assert.True(t, ll.Contains(cInt(5)))
	assert.True(t, ll.Contains(cInt(6)))
	assert.True(t, ll.Contains(cInt(7)))
	assert.False(t, ll.Contains(cInt(33)))
	assert.False(t, ll.Contains(cInt(999)))
	assert.False(t, ll.Contains(cInt(51)))
}

func TestLinkedList_Get(t *testing.T) {
	ll := prepareLinkedList()

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

func TestLinkedList_Set(t *testing.T) {
	ll := prepareLinkedList()

	ll.Set(0, cInt(500))

	assert.False(t, ll.Contains(cInt(1)))
	assert.True(t, ll.Contains(cInt(500)))
}

func TestLinkedList_Enqueue(t *testing.T) {
	ll := NewLinkedList()

	q := ll.(api.Queue)

	q.Enqueue(cInt(50))
	q.Enqueue(cInt(60))
	q.Enqueue(cInt(70))

	assert.True(t, ll.Contains(cInt(50)))
	assert.True(t, ll.Contains(cInt(60)))
	assert.True(t, ll.Contains(cInt(70)))
}

func TestLinkedList_Dequeue(t *testing.T) {
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

func TestLinkedList_Unshift(t *testing.T) {
	ll := prepareLinkedList()

	ll.Unshift(cInt(500))

	assert.True(t, ll.Contains(cInt(500)))
	assert.False(t, ll.Contains(cInt(501)))
}

func TestLinkedList_Shift(t *testing.T) {
	ll := prepareLinkedList()

	val, isV := ll.Shift()

	assert.True(t, isV)
	assert.True(t, cInt(1).Equal(val))
}

func TestLinkedList_Remove(t *testing.T) {
	ll := prepareLinkedList()

	r1, isR1 := ll.Remove(cInt(1))

	assert.True(t, isR1)
	assert.NotNil(t, r1)
	assert.False(t, ll.Contains(cInt(1)))
}

func TestLinkedList_RemoveIdx(t *testing.T) {
	ll := prepareLinkedList()

	r1, isR1 := ll.RemoveIdx(0)

	assert.True(t, isR1)
	assert.NotNil(t, r1)
	assert.False(t, ll.Contains(cInt(1)))
}

func TestLinkedList_Insert(t *testing.T) {
	ll := prepareLinkedList()

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

func TestLinkedList_Size(t *testing.T) {
	ll := prepareLinkedList()

	ll.Push(cInt(999))
	ll.Push(cInt(998))
	ll.Remove(cInt(999))

	assert.Equal(t, prepareLinkedList().Size()+1, ll.Size())
}
