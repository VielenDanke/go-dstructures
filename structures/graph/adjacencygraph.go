package graph

import (
	"github.com/vielendanke/go-dstructures/structures/api"
	"github.com/vielendanke/go-dstructures/structures/hashtable"
)

type adjacencyGraph struct {
	m      api.Map
	length int
}

func NewAdjacencyGraph() api.Graph {
	return &adjacencyGraph{m: hashtable.NewHashMap(16)}
}

func (a *adjacencyGraph) AddEdge(fVertex api.EqualHashRule, sVertex api.EqualHashRule) bool {
	f := a.m.Get(fVertex)
	s := a.m.Get(sVertex)
	if f == nil && s == nil {
		return false
	}
	fArr := f.([]api.EqualHashRule)
	sArr := f.([]api.EqualHashRule)
	fArr = append(fArr, sVertex)
	sArr = append(sArr, fVertex)
	a.m.Put(fVertex, fArr)
	a.m.Put(sVertex, sArr)
	return true
}

func (a *adjacencyGraph) AddVertex(val api.EqualHashRule) bool {
	f := a.m.Get(val)
	if f != nil {
		return false
	}
	a.m.Put(val, make([]api.EqualHashRule, 0))
	a.length++
	return true
}

func (a *adjacencyGraph) RemoveEdge(fVertex api.EqualHashRule, sVertex api.EqualHashRule) bool {
	f := a.m.Get(fVertex).([]api.EqualHashRule)
	s := a.m.Get(sVertex).([]api.EqualHashRule)
	if f != nil && s != nil {
		for k, v := range f {
			if v.Equal(sVertex) {
				f = append(f[:k], f[k+1:]...)
				a.m.Put(fVertex, f)
			}
		}
		for k, v := range s {
			if v.Equal(fVertex) {
				s = append(s[:k], s[k+1:]...)
				a.m.Put(sVertex, s)
			}
		}
	} else {
		return false
	}
	return true
}

func (a *adjacencyGraph) RemoveVertex(fVertex api.EqualHashRule) bool {
	fArr := a.m.Get(fVertex).([]api.EqualHashRule)
	if fArr == nil {
		return false
	}
	for _, v := range fArr {
		temp := a.m.Get(v).([]api.EqualHashRule)
		if temp != nil {
			for k, tv := range temp {
				if tv.Equal(fVertex) {
					temp = append(temp[:k], temp[k+1:]...)
				}
			}
		}
		a.m.Put(v, temp)
	}
	a.m.Remove(fVertex)
	a.length--
	return true
}

func (a *adjacencyGraph) ToArray() []api.EqualHashRule {
	return a.m.KeySet()
}

func (a *adjacencyGraph) Size() int {
	return a.length
}

func (a *adjacencyGraph) recursiveInvoker(arr []api.EqualHashRule, m api.Map) {
	for _, v := range arr {
		if m.Get(v) != nil {
			m.Put(v, true)
			a.recursiveInvoker(a.m.Get(v).([]api.EqualHashRule), m)
		} else {
			continue
		}
	}
}
