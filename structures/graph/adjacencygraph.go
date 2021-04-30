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

func (a *adjacencyGraph) AddEdge(fVertex api.EqualHashRule, sVertex api.EqualHashRule, weight int64) bool {
	f := a.m.Get(fVertex)
	s := a.m.Get(sVertex)
	if f == nil || s == nil {
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
	f := a.m.Get(fVertex)
	s := a.m.Get(sVertex)

	if f != nil && s != nil {
		fConv := f.([]api.EqualHashRule)
		sConv := s.([]api.EqualHashRule)

		for k, v := range fConv {
			if v.Equal(sVertex) {
				fConv = append(fConv[:k], fConv[k+1:]...)
				a.m.Put(fVertex, fConv)
			}
		}
		for k, v := range sConv {
			if v.Equal(fVertex) {
				sConv = append(sConv[:k], sConv[k+1:]...)
				a.m.Put(sVertex, sConv)
			}
		}
	} else {
		return false
	}
	return true
}

func (a *adjacencyGraph) RemoveVertex(fVertex api.EqualHashRule) bool {
	fArr := a.m.Get(fVertex)

	if fArr == nil {
		return false
	}
	convArr := fArr.([]api.EqualHashRule)

	for _, v := range convArr {
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

func (a *adjacencyGraph) GetEdges(vertex api.EqualHashRule) []api.EqualHashRule {
	edges := a.m.Get(vertex)

	if edges == nil {
		return nil
	}
	return edges.([]api.EqualHashRule)
}

func (a *adjacencyGraph) ToArray() []api.EqualHashRule {
	return a.m.KeySet()
}

func (a *adjacencyGraph) Size() int {
	return a.m.Size()
}

func (a *adjacencyGraph) Contains(vertex api.EqualHashRule) bool {
	for _, v := range a.ToArray() {
		if v.Equal(vertex) {
			return true
		}
	}
	return false
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
