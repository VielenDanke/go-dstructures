package hashtable

import (
	"errors"
	"github.com/vielendanke/go-dstructures/structures/api"
	"github.com/vielendanke/go-dstructures/structures/tree"
)

type treeMap struct {
	elements   []api.Tree
	sortFunc   api.Sort
	size       int
	capacity   int
	loadFactor float64
}

func NewTreeMap(capacity int, sortFunc api.Sort) (api.Map, error) {
	if sortFunc == nil {
		return nil, errors.New("sort func is not present")
	}
	if capacity < 0 {
		return nil, errors.New("capacity cannot be less than 0")
	}
	ht := &treeMap{sortFunc: sortFunc, elements: make([]api.Tree, capacity)}
	ht.capacity = capacity
	ht.loadFactor = 0.75
	return ht, nil
}

func (ht *treeMap) Remove(key api.EqualHashRule) (api.EqualHashRule, interface{}) {
	hash := ht.hashFunction(key)
	if ht.elements[hash] == nil {
		return nil, nil
	} else {
		tm := ht.elements[hash]
		ht.size--
		return tm.Remove(key)
	}
}

func (ht *treeMap) Get(key api.EqualHashRule) interface{} {
	hash := ht.hashFunction(key)
	if ht.elements[hash] == nil {
		return nil
	} else {
		tm := ht.elements[hash]
		return tm.Get(key)
	}
}

func (ht *treeMap) Put(key api.EqualHashRule, val interface{}) {
	if ht.size+1 > int(float64(ht.capacity)*ht.loadFactor) {
		ht.capacity = ht.capacity * 2
		ht.increaseMap()
	}
	hash := ht.hashFunction(key)
	if ht.elements[hash] != nil {
		tm := ht.elements[hash]
		tm.Put(key, val)
	} else {
		tm := tree.NewRBTree(ht.sortFunc)
		tm.Put(key, val)
		ht.elements[hash] = tm
	}
	ht.size++
}

func (ht *treeMap) Contains(key api.EqualHashRule) bool {
	hash := ht.hashFunction(key)
	tm := ht.elements[hash]
	if tm == nil {
		return false
	} else {
		return tm.Contains(key)
	}
}

func (ht *treeMap) Size() int {
	return ht.size
}

func (ht *treeMap) hashFunction(key api.EqualHashRule) (idx int) {
	prime := 31
	if key == nil {
		return 0
	}
	res := key.Hash() * prime
	idx = (res ^ (res >> 16)) & (ht.capacity - 1)
	return
}

func (ht *treeMap) increaseMap() {
	newElements := make([]api.Tree, ht.capacity)
	for _, v := range ht.elements {
		if v != nil {
			r, _ := v.GetRoot()
			hash := ht.hashFunction(r)
			newElements[hash] = v
		}
	}
	ht.elements = newElements
}
