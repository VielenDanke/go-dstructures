package linkedlist

import (
	"fmt"
	"github.com/vielendanke/go-dstructures/structures/api"
)

type singleNode struct {
	val  interface{}
	next *singleNode
}

type singlyLinkedList struct {
	head   *singleNode
	tail   *singleNode
	length int
}

func NewSinglyLinkedList() api.LinkedList {
	return &singlyLinkedList{}
}

func (ll *singlyLinkedList) Enqueue(val interface{}) {
	ll.Push(val)
}

func (ll *singlyLinkedList) Dequeue() (interface{}, bool) {
	return ll.Shift()
}

func (ll *singlyLinkedList) Push(val interface{}) {
	n := &singleNode{val: val}
	if ll.head == nil {
		ll.head = n
		ll.tail = ll.head
	} else {
		ll.tail.next = n
		ll.tail = ll.tail.next
	}
	ll.length++
}

func (ll *singlyLinkedList) Pop() (interface{}, bool) {
	n, isFound := ll.popNode()
	if !isFound {
		return nil, false
	}
	return n.val, true
}

func (ll *singlyLinkedList) Shift() (interface{}, bool) {
	n, isFound := ll.shiftNode()
	if !isFound {
		return nil, false
	}
	return n.val, true
}

func (ll *singlyLinkedList) Unshift(val interface{}) {
	if ll.length == 0 {
		ll.Push(val)
	} else {
		n := &singleNode{val: val}
		tempHead := ll.head
		ll.head = n
		ll.head.next = tempHead
	}
	ll.length++
}

func (ll *singlyLinkedList) Get(idx int) (interface{}, bool) {
	n, isFound := ll.getNode(idx)
	if !isFound {
		return nil, false
	}
	return n.val, true
}

func (ll *singlyLinkedList) Set(idx int, val interface{}) (ok bool) {
	fNode, ok := ll.getNode(idx)
	if !ok {
		return ok
	}
	fNode.val = val
	return ok
}

func (ll *singlyLinkedList) Insert(idx int, val interface{}) (ok bool) {
	if idx < 0 || idx > ll.length {
		return
	}
	if idx == ll.length-1 {
		ll.Push(val)
		ok = true
		ll.length++
		return
	}
	if idx == 0 {
		ll.Unshift(val)
		ok = true
		ll.length++
		return
	}
	prev, isFound := ll.getNode(idx - 1)
	if !isFound {
		ok = false
		return
	}
	n := &singleNode{val: val}
	temp := prev.next
	prev.next = n
	n.next = temp
	ok = true
	ll.length++
	return
}

func (ll *singlyLinkedList) Remove(idx int) (interface{}, bool) {
	n, isRemoved := ll.removeNode(idx)
	if !isRemoved {
		return nil, false
	}
	return n.val, true
}

func (ll *singlyLinkedList) Reverse() {
	var next, prev *singleNode
	curr := ll.head
	ll.head = ll.tail
	ll.tail = curr
	for i := 0; i < ll.length; i++ {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
}

func (ll *singlyLinkedList) Size() int {
	return ll.length
}

func (ll *singlyLinkedList) String() string {
	arr := make([]interface{}, 0)
	current := ll.head
	for current != nil {
		arr = append(arr, current.val)
		current = current.next
	}
	return fmt.Sprintf("%v", arr)
}

func (ll *singlyLinkedList) popNode() (toFind *singleNode, isFound bool) {
	if ll.length == 0 {
		return toFind, isFound
	}
	if ll.length == 1 {
		toFind = ll.head
		ll.head = nil
		ll.tail = nil
	} else {
		current := ll.head
		for i := 0; i < ll.length; i++ {
			if i == ll.length-2 {
				toFind = current.next
				current.next = nil
				ll.tail = current
				break
			}
			current = current.next
		}
	}
	ll.length--
	isFound = true
	return toFind, isFound
}

func (ll *singlyLinkedList) shiftNode() (toFind *singleNode, isFound bool) {
	if ll.length == 0 {
		return toFind, isFound
	}
	toFind = ll.head
	ll.head = ll.head.next
	isFound = true
	ll.length--
	if ll.length == 0 {
		ll.tail = nil
	}
	return toFind, isFound
}

func (ll *singlyLinkedList) getNode(idx int) (toFound *singleNode, isFound bool) {
	if ll.length == 0 {
		return toFound, isFound
	}
	if idx > ll.length || idx < 0 {
		return toFound, isFound
	}
	current := ll.head
	for i := 0; i < ll.length; i++ {
		if i == idx {
			isFound = true
			toFound = current
			return toFound, isFound
		}
		current = current.next
	}
	return toFound, isFound
}

func (ll *singlyLinkedList) removeNode(idx int) (*singleNode, bool) {
	if idx < 0 || idx > ll.length {
		return nil, false
	}
	if idx == 0 {
		return ll.shiftNode()
	}
	if idx == ll.length-1 {
		return ll.popNode()
	}
	prevNode, isFound := ll.getNode(idx - 1)
	if !isFound {
		return nil, isFound
	}
	removed := prevNode.next
	prevNode.next = removed.next
	ll.length--
	return removed, true
}
