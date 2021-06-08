package structures

import (
	"fmt"

	"github.com/vielendanke/go-dstructures/structures/api"
)

type hashSet struct {
	m api.Map
}

func NewHashSet() api.Set {
	return &hashSet{m: NewHashMap(16)}
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

func (h *hashSet) String() string {
	arr := make([]string, 0)
	for _, v := range h.ToArray() {
		arr = append(arr, fmt.Sprintf("%v", v))
	}
	return fmt.Sprintf("%v", arr)
}
