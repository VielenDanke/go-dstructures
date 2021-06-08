package structures

import (
	"fmt"

	"github.com/vielendanke/go-dstructures/structures/api"
)

type linkedNode struct {
	prev *linkedNode
	next *linkedNode
	val  api.EqualHashRule
}

type linkedList struct {
	head   *linkedNode
	tail   *linkedNode
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
				dl.unlinklinkedNode(current)
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
	n := &linkedNode{val: val}
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
	n, isPopped := dl.poplinkedNode()
	if !isPopped {
		return nil, false
	}
	return n.val, true
}

func (dl *linkedList) Shift() (api.EqualHashRule, bool) {
	n, isShifted := dl.shiftlinkedNode()
	if !isShifted {
		return nil, false
	}
	return n.val, true
}

func (dl *linkedList) Unshift(val api.EqualHashRule) {
	n := &linkedNode{val: val}
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
	n, isFound := dl.getlinkedNode(idx)
	if !isFound {
		return nil, false
	}
	return n.val, true
}

func (dl *linkedList) Set(idx int, val api.EqualHashRule) bool {
	foundlinkedNode, isFound := dl.getlinkedNode(idx)
	if !isFound {
		return isFound
	}
	foundlinkedNode.val = val
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
		newlinkedNode := &linkedNode{val: val}
		prevlinkedNode, isFound := dl.getlinkedNode(idx - 1)
		if !isFound {
			return isFound
		}
		nextlinkedNode := prevlinkedNode.next
		prevlinkedNode.next = newlinkedNode
		nextlinkedNode.prev = newlinkedNode
		newlinkedNode.prev = prevlinkedNode
		newlinkedNode.next = nextlinkedNode
	}
	dl.length++
	return true
}

func (dl *linkedList) RemoveIdx(idx int) (api.EqualHashRule, bool) {
	n, isRemoved := dl.removelinkedNode(idx)
	if !isRemoved {
		return nil, false
	}
	return n.val, true
}

func (dl *linkedList) Reverse() {
	head := &linkedNode{val: dl.tail.val}
	tail := &linkedNode{val: dl.head.val}
	nextlinkedNode := dl.tail
	tempHead := head
	for i := 1; i < dl.length; i++ {
		n := &linkedNode{val: nextlinkedNode.prev.val}
		tempHead.next = n
		n.prev = tempHead
		if i == dl.length-1 {
			tempHead.next = tail
			tail.prev = tempHead
		} else {
			tempHead = tempHead.next
		}
		nextlinkedNode = nextlinkedNode.prev
	}
	dl.head = head
	dl.tail = tail
}

func (dl *linkedList) Size() int {
	return dl.length
}

func (dl *linkedList) ToArray() []api.EqualHashRule {
	arr := make([]api.EqualHashRule, 0)
	curr := dl.head
	for curr != nil {
		arr = append(arr, curr.val)
		curr = curr.next
	}
	return arr
}

func (dl *linkedList) String() string {
	return fmt.Sprintf("%v", dl.ToArray())
}

func (dl *linkedList) removelinkedNode(idx int) (foundlinkedNode *linkedNode, isFound bool) {
	if idx < 0 || idx >= dl.length {
		return
	}
	if idx == 0 {
		foundlinkedNode, isFound = dl.shiftlinkedNode()
		return
	} else if idx == dl.length-1 {
		foundlinkedNode, isFound = dl.poplinkedNode()
		return
	} else {
		foundlinkedNode, isFound = dl.getlinkedNode(idx)
		if !isFound {
			return
		}
		dl.unlinklinkedNode(foundlinkedNode)
	}
	dl.length--
	isFound = true
	return
}

func (dl *linkedList) unlinklinkedNode(foundlinkedNode *linkedNode) {
	if foundlinkedNode.prev == nil {
		dl.head = foundlinkedNode.next
		dl.head.prev = nil
		return
	}
	if foundlinkedNode.next == nil {
		dl.tail = foundlinkedNode.prev
		dl.tail.next = nil
		return
	}
	foundlinkedNode.prev.next = foundlinkedNode.next
	foundlinkedNode.next.prev = foundlinkedNode.prev
	foundlinkedNode.prev = nil
	foundlinkedNode.next = nil
}

func (dl *linkedList) poplinkedNode() (*linkedNode, bool) {
	var remove *linkedNode
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

func (dl *linkedList) shiftlinkedNode() (*linkedNode, bool) {
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

func (dl *linkedList) getlinkedNode(idx int) (*linkedNode, bool) {
	var foundlinkedNode *linkedNode
	if idx < 0 || idx >= dl.length {
		return nil, false
	}
	half := dl.length / 2
	if idx < half {
		curr := dl.head
		for i := 0; i < half; i++ {
			if i == idx {
				foundlinkedNode = curr
			}
			curr = curr.next
		}
	} else {
		curr := dl.tail
		for i := dl.length - 1; i >= half; i-- {
			if i == idx {
				foundlinkedNode = curr
			}
			curr = curr.prev
		}
	}
	return foundlinkedNode, true
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
