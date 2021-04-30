package hashtable

import (
	"fmt"
	"github.com/vielendanke/go-dstructures/structures/api"
)

type node struct {
	key   api.EqualHashRule
	value interface{}
	prev  *node
	next  *node
}

func (n *node) String() string {
	return fmt.Sprintf("Node: {%v:%v}", n.key, n.value)
}

type linkedMap struct {
	head *node
	tail *node
}

type hashMap struct {
	elements   []*linkedMap
	size       int
	capacity   int
	loadFactor float64
}

func NewHashMap(capacity int) api.Map {
	return &hashMap{
		loadFactor: 0.75,
		capacity:   capacity,
		elements:   make([]*linkedMap, capacity),
	}
}

func (h *hashMap) Get(key api.EqualHashRule) interface{} {
	if key == nil && h.Size() == 0 {
		return nil
	}
	hash := h.hashFunction(key)
	l := h.elements[hash]
	if l == nil {
		return nil
	} else {
		current := l.head
		for current != nil {
			if current.key.Equal(key) {
				return current.value
			}
			current = current.next
		}
	}
	return nil
}

func (h *hashMap) Put(key api.EqualHashRule, val interface{}) {
	if key == nil {
		return
	}
	if h.size+1 > int(float64(h.capacity)*h.loadFactor) {
		h.capacity = h.capacity * 2
		h.increaseMap()
	}
	newNode := &node{key: key, value: val}
	hash := h.hashFunction(key)
	l := h.elements[hash]
	if l != nil {
		current := l.head
		isFound := false
		for current != nil {
			if current.key.Equal(key) {
				current.value = val
				isFound = true
				break
			}
			current = current.next
		}
		if !isFound {
			temp := l.tail
			temp.next = newNode
			l.tail = newNode
			l.tail.prev = temp
			h.size++
		}
	} else {
		nm := &linkedMap{}
		nm.head = newNode
		nm.tail = newNode
		h.elements[hash] = nm
		h.size++
	}
}

func (h *hashMap) Contains(key api.EqualHashRule) bool {
	if key == nil && h.Size() == 0 {
		return false
	}
	v := h.Get(key)
	if v == nil {
		return false
	}
	return true
}

func (h *hashMap) Remove(key api.EqualHashRule) (api.EqualHashRule, interface{}) {
	if key == nil && h.Size() == 0 {
		return nil, nil
	}
	hash := h.hashFunction(key)
	l := h.elements[hash]
	if l == nil {
		return nil, nil
	} else {
		if l.head.next == nil {
			if l.head.key.Equal(key) {
				temp := l.head
				h.elements[hash] = nil
				h.size--
				return temp.key, temp.value
			} else {
				return nil, nil
			}
		} else {
			current := l.head
			for current != nil {
				if current.key.Equal(key) {
					temp := current
					current.prev.next = current.next
					current.next.prev = current.prev
					h.size--
					return temp.key, temp.value
				}
				current = current.next
			}
		}
	}
	return nil, nil
}

func (h *hashMap) KeySet() []api.EqualHashRule {
	result := make([]api.EqualHashRule, 0)
	for _, v := range h.elements {
		if v != nil {
			current := v.head
			for current != nil {
				result = append(result, current.key)
				current = current.next
			}
		}
	}
	return result
}

func (h *hashMap) Size() int {
	return h.size
}

func (h *hashMap) increaseMap() {
	newElements := make([]*linkedMap, h.capacity)
	for _, v := range h.elements {
		if v != nil {
			hash := h.hashFunction(v.head.key)
			newElements[hash] = v
		}
	}
	h.elements = newElements
}

func (h *hashMap) hashFunction(key api.EqualHashRule) (idx int) {
	prime := 31
	if key == nil {
		return 0
	}
	res := key.Hash() * prime
	idx = (res ^ (res >> 16)) & (h.capacity - 1)
	return
}
