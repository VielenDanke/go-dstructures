package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Push(t *testing.T) {
	ll := prepareLinkedList()

	removed, isRemoved := ll.Remove(cInt(5))

	assert.True(t, isRemoved)
	assert.NotNil(t, removed)
	assert.False(t, ll.Contains(cInt(5)))
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

func TestLinkedList_Enqueue(t *testing.T) {

}

func TestLinkedList_Dequeue(t *testing.T) {

}