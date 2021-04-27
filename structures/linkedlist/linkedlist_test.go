package linkedlist

import (
	"github.com/vielendanke/go-dstructures/structures/api"
	"os"
	"testing"
)

type cInt int

func (c cInt) Equal(p interface{}) bool {
	return c.Hash() == p.(cInt).Hash()
}

func (c cInt) Hash() int {
	return int(c)
}

func prepareLinkedList() (ll api.LinkedList) {
	ll = NewLinkedList()
	fillList(ll)
	return
}

func prepareSinglyLinkedList() (ll api.LinkedList) {
	ll = NewSinglyLinkedList()
	fillList(ll)
	return
}

func fillList(ll api.LinkedList) {
	for i := 1; i < 21; i++ {
		ll.Push(cInt(i))
	}
}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
