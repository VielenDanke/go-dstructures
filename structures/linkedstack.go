package structures

import (
	"fmt"

	"github.com/vielendanke/go-dstructures/structures/api"
)

type stackNode struct {
	val  api.EqualHashRule
	next *stackNode
}

type linkedStack struct {
	first  *stackNode
	last   *stackNode
	length int
}

func NewLinkedStack() api.Stack {
	return &linkedStack{}
}

func (lls *linkedStack) Push(val api.EqualHashRule) {
	newstackNode := &stackNode{val: val}
	if lls.length == 0 {
		lls.first = newstackNode
		lls.last = newstackNode
	} else {
		temp := lls.first
		lls.first = newstackNode
		lls.first.next = temp
	}
	lls.length++
}

func (lls *linkedStack) Pop() (val api.EqualHashRule, isFound bool) {
	if lls.length == 0 {
		return
	}
	if lls.length == 1 {
		val = lls.first.val
		lls.first = nil
		lls.last = nil
	} else {
		val = lls.first.val
		lls.first = lls.first.next
	}
	lls.length--
	isFound = !isFound
	return
}

func (lls *linkedStack) Size() int {
	return lls.length
}

func (lls *linkedStack) String() string {
	arr := make([]api.EqualHashRule, 0)
	curr := lls.first
	for curr != nil {
		arr = append(arr, curr.val)
		curr = curr.next
	}
	return fmt.Sprintf("%v", arr)
}
