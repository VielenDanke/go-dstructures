package set

import (
	"github.com/vielendanke/go-dstructures/structures/api"
	"github.com/vielendanke/go-dstructures/structures/hashtable"
)

type treeSet struct {
	tm api.Map
}

func NewTreeSet(sortFunc func(left interface{}, right interface{}) bool) (api.Set, error) {
	hm, err := hashtable.NewTreeMap(16, sortFunc)
	if err != nil {
		return nil, err
	}
	return &treeSet{tm: hm}, nil
}

func (bt *treeSet) ToArray() []api.EqualHashRule {
	return bt.tm.KeySet()
}

func (bt *treeSet) Add(val api.EqualHashRule) {
	bt.tm.Put(val, nil)
}

func (bt *treeSet) Contains(val api.EqualHashRule) bool {
	return bt.tm.Contains(val)
}

func (bt *treeSet) Remove(val api.EqualHashRule) api.EqualHashRule {
	removedElem, _ := bt.tm.Remove(val)
	return removedElem
}

func (bt *treeSet) Size() int {
	return bt.tm.Size()
}
