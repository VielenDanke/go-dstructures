package queue

import (
	"fmt"
	"github.com/vielendanke/go-dstructures/structures/api"
)

type node struct {
	val api.EqualHashRule
	next *node
}

type linkedQueue struct {
	first *node
	last *node
	length int
}

func NewLinkedQueue() api.Queue {
	return &linkedQueue{}
}

func (lq *linkedQueue) Enqueue(val api.EqualHashRule) {
	n := &node{val: val}
	if lq.length == 0 {
		lq.first = n
		lq.last = n
	} else {
		lq.last.next = n
		lq.last = n
	}
	lq.length++
}

func (lq *linkedQueue) Dequeue() (api.EqualHashRule, bool) {
	if lq.length == 0 {
		return nil, false
	}
	n := lq.first
	if lq.length == 1 {
		lq.first = nil
		lq.last = nil
	} else {
		lq.first = lq.first.next
	}
	lq.length--
	return n.val, true
}

func (lq *linkedQueue) Size() int {
	return lq.length
}

func (lq *linkedQueue) String() string {
	var arr []interface{}
	current := lq.first
	for current != nil {
		arr = append(arr, current.val)
		current = current.next
	}
	return fmt.Sprintf("%v", arr)
}