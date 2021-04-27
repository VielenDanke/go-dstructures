package linkedlist

import (
	"fmt"
	"github.com/vielendanke/go-dstructures/structures/api"
)

type node struct {
	prev *node
	next *node
	val  api.EqualHashRule
}

type linkedList struct {
	head   *node
	tail   *node
	length int
}

func NewLinkedList() api.LinkedList {
	return &linkedList{}
}

func (dl *linkedList) Remove(val api.EqualHashRule) (api.EqualHashRule, bool) {
	if dl.head == nil {
		return nil, false
	}
	if dl.length == 1 {
		fVal := dl.head.val
		if equalVal(fVal, val) {
			dl.head = nil
			dl.tail = nil
			dl.length--
			return fVal, true
		} else {
			return nil, false
		}
	} else {
		current := dl.head
		for current != nil {
			if equalVal(current.val, val) {
				dl.length--
				dl.unlinkNode(current)
				return current.val, true
			}
			current = current.next
		}
	}
	return nil, false
}

func (dl *linkedList) Contains(val api.EqualHashRule) bool {
	if dl.head == nil {
		return false
	}
	if dl.length == 1 {
		return equalVal(dl.head.val, val)
	}
	current := dl.head
	for current != nil {
		if equalVal(current.val, val) {
			return true
		}
		current = current.next
	}
	return false
}

func (dl *linkedList) Enqueue(val api.EqualHashRule) {
	dl.Push(val)
}

func (dl *linkedList) Dequeue() (api.EqualHashRule, bool) {
	return dl.Shift()
}

func (dl *linkedList) Push(val api.EqualHashRule) {
	n := &node{val: val}
	if dl.head == nil {
		dl.head = n
		dl.tail = n
	} else {
		dl.tail.next = n
		n.prev = dl.tail
		dl.tail = n
	}
	dl.length++
}

func (dl *linkedList) Pop() (api.EqualHashRule, bool) {
	n, isPopped := dl.popNode()
	if !isPopped {
		return nil, false
	}
	return n.val, true
}

func (dl *linkedList) Shift() (api.EqualHashRule, bool) {
	n, isShifted := dl.shiftNode()
	if !isShifted {
		return nil, false
	}
	return n.val, true
}

func (dl *linkedList) Unshift(val api.EqualHashRule) {
	n := &node{val: val}
	if dl.length == 0 {
		dl.head = n
		dl.tail = n
	} else {
		temp := dl.head
		dl.head = n
		n.next = temp
		temp.prev = n
	}
	dl.length++
}

func (dl *linkedList) Get(idx int) (api.EqualHashRule, bool) {
	n, isFound := dl.getNode(idx)
	if !isFound {
		return nil, false
	}
	return n.val, true
}

func (dl *linkedList) Set(idx int, val api.EqualHashRule) bool {
	foundNode, isFound := dl.getNode(idx)
	if !isFound {
		return isFound
	}
	foundNode.val = val
	return isFound
}

func (dl *linkedList) Insert(idx int, val api.EqualHashRule) bool {
	if idx < 0 || idx > dl.length {
		return false
	}
	if idx == 0 {
		dl.Unshift(val)
	} else if idx == dl.length {
		dl.Push(val)
	} else {
		newNode := &node{val: val}
		prevNode, isFound := dl.getNode(idx - 1)
		if !isFound {
			return isFound
		}
		nextNode := prevNode.next
		prevNode.next = newNode
		nextNode.prev = newNode
		newNode.prev = prevNode
		newNode.next = nextNode
	}
	dl.length++
	return true
}

func (dl *linkedList) RemoveIdx(idx int) (api.EqualHashRule, bool) {
	n, isRemoved := dl.removeNode(idx)
	if !isRemoved {
		return nil, false
	}
	return n.val, true
}

func (dl *linkedList) Reverse() {
	head := &node{val: dl.tail.val}
	tail := &node{val: dl.head.val}
	nextNode := dl.tail
	tempHead := head
	for i := 1; i < dl.length; i++ {
		n := &node{val: nextNode.prev.val}
		tempHead.next = n
		n.prev = tempHead
		if i == dl.length-1 {
			tempHead.next = tail
			tail.prev = tempHead
		} else {
			tempHead = tempHead.next
		}
		nextNode = nextNode.prev
	}
	dl.head = head
	dl.tail = tail
}

func (dl *linkedList) Size() int {
	return dl.length
}

func (dl *linkedList) String() string {
	arr := make([]api.EqualHashRule, 0)
	curr := dl.head
	for curr != nil {
		arr = append(arr, curr.val)
		curr = curr.next
	}
	return fmt.Sprintf("%v", arr)
}

func (dl *linkedList) removeNode(idx int) (foundNode *node, isFound bool) {
	if idx < 0 || idx >= dl.length {
		return
	}
	if idx == 0 {
		foundNode, isFound = dl.shiftNode()
	} else if idx == dl.length-1 {
		foundNode, isFound = dl.popNode()
	} else {
		foundNode, isFound = dl.getNode(idx)
		if !isFound {
			return
		}
		dl.unlinkNode(foundNode)
	}
	dl.length--
	isFound = true
	return
}

func (dl *linkedList) unlinkNode(foundNode *node) {
	if foundNode.prev == nil {
		dl.head = foundNode.next
		dl.head.prev = nil
		return
	}
	if foundNode.next == nil {
		dl.tail = foundNode.prev
		dl.tail.next = nil
		return
	}
	foundNode.prev.next = foundNode.next
	foundNode.next.prev = foundNode.prev
	foundNode.prev = nil
	foundNode.next = nil
}

func (dl *linkedList) popNode() (*node, bool) {
	var remove *node
	if dl.head == nil {
		return nil, false
	}
	remove = dl.tail
	if dl.length == 1 {
		dl.head = nil
		dl.tail = nil
	} else {
		dl.tail = remove.prev
		dl.tail.next = nil
		remove.prev = nil
	}
	dl.length--
	return remove, true
}

func (dl *linkedList) shiftNode() (*node, bool) {
	if dl.length == 0 {
		return nil, false
	}
	head := dl.head
	if dl.length == 1 {
		dl.head = nil
		dl.tail = nil
	} else {
		dl.head = head.next
		dl.head.prev = nil
		head.next = nil
	}
	dl.length--
	return head, true
}

func (dl *linkedList) getNode(idx int) (*node, bool) {
	var foundNode *node
	if idx < 0 || idx >= dl.length {
		return nil, false
	}
	half := dl.length / 2
	if idx < half {
		curr := dl.head
		for i := 0; i < half; i++ {
			if i == idx {
				foundNode = curr
			}
			curr = curr.next
		}
	} else {
		curr := dl.tail
		for i := dl.length - 1; i >= half; i-- {
			if i == idx {
				foundNode = curr
			}
			curr = curr.prev
		}
	}
	return foundNode, true
}

func equalVal(fVal, sVal api.EqualHashRule) bool {
	if fVal == nil && sVal == nil {
		return true
	}
	if fVal != nil && sVal != nil {
		return fVal.Equal(sVal)
	}
	return false
}
