package set

import (
	"github.com/vielendanke/go-dstructures/structures/api"
	"github.com/vielendanke/go-dstructures/structures/hashtable"
)

type hashSet struct {
	m api.Map
}

func NewHashSet() api.Set {
	return &hashSet{m: hashtable.NewHashMap(16)}
}

func (h *hashSet) ToArray() []api.EqualHashRule {
	return h.m.KeySet()
}

func (h *hashSet) Size() int {
	return h.m.Size()
}

func (h *hashSet) Add(val api.EqualHashRule) {
	h.m.Put(val, nil)
}

func (h *hashSet) Contains(val api.EqualHashRule) bool {
	return h.m.Contains(val)
}

func (h *hashSet) Remove(val api.EqualHashRule) api.EqualHashRule {
	key, _ := h.m.Remove(val)
	return key
}
