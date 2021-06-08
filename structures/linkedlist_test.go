package structures

import (
	"github.com/vielendanke/go-dstructures/structures/api"
)

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
