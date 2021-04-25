package graph

import (
	"github.com/vielendanke/go-dstructures/structures/api"
)

type adjacencyGraph struct {
	m map[string][]string
}

func NewGraph() api.Graph {
	return &adjacencyGraph{m: make(map[string][]string)}
}

func (a *adjacencyGraph) AddEdge(fVertex string, sVertex string) {
	f, fOk := a.m[fVertex]
	s, sOk := a.m[sVertex]
	if fOk {
		f = append(f, sVertex)
		a.m[fVertex] = f
	} else {
		a.m[fVertex] = append(make([]string, 0), sVertex)
	}
	if sOk {
		s = append(s, fVertex)
		a.m[sVertex] = s
	} else {
		a.m[sVertex] = append(make([]string, 0), fVertex)
	}
}

func (a *adjacencyGraph) AddVertex(val string) {
	_, ok := a.m[val]
	if ok {
		return
	}
	a.m[val] = make([]string, 0)
}

func (a *adjacencyGraph) RemoveEdge(fVertex string, sVertex string) {
	f, fOk := a.m[fVertex]
	s, sOk := a.m[sVertex]
	if fOk {
		for k, v := range f {
			if v == sVertex {
				a.m[fVertex] = append(f[:k], f[k+1:]...)
			}
		}
	}
	if sOk {
		for k, v := range s {
			if v == fVertex {
				a.m[sVertex] = append(s[:k], s[k+1:]...)
			}
		}
	}
}

func (a *adjacencyGraph) RemoveVertex(fVertex string) bool {
	fArr, ok := a.m[fVertex]
	if !ok {
		return false
	}
	for _, v := range fArr {
		temp, tOk := a.m[v]
		if tOk {
			for k, tv := range temp {
				if tv == fVertex {
					a.m[v] = append(temp[:k], temp[k+1:]...)
				}
			}
		}
	}
	delete(a.m, fVertex)
	return true
}

func (a *adjacencyGraph) DepthFirstRecursive(root string) []string {
	markedNodes := make(map[string]bool)
	v, ok := a.m[root]
	if !ok {
		return nil
	}
	markedNodes[root] = true
	a.recursiveInvoker(v, markedNodes)
	resArr := make([]string, 0)
	for k := range markedNodes {
		resArr = append(resArr, k)
	}
	return resArr
}

func (a *adjacencyGraph) DepthFirstIterate(root string) []string {
	markedNodes := make(map[string]bool)
	res := make([]string, 0)
	stack := append(make([]string, 0), root)
	markedNodes[root] = true
	for len(stack) > 0 {
		currVertex := stack[0]
		stack = stack[1:]
		res = append(res, currVertex)
		tempArr, _ := a.m[currVertex]
		for _, v := range tempArr {
			if !markedNodes[v] {
				markedNodes[v] = true
				stack = append(stack, v)
			}
		}
	}
	return res
}

func (a *adjacencyGraph) BreadthFirst(root string) []string {
	_, ok := a.m[root]
	if !ok {
		return nil
	}
	res := make([]string, 0)
	markedNodes := make(map[string]bool)
	q := make([]string, 0)
	markedNodes[root] = true
	q = append(q, root)
	for len(q) != 0 {
		temp := q[0]
		res = append(res, temp)
		v := a.m[temp]
		q = q[1:]
		for _, tv := range v {
			if !markedNodes[tv] {
				markedNodes[tv] = true
				q = append(q, tv)
			}
		}
	}
	return res
}

func (a *adjacencyGraph) recursiveInvoker(arr []string, m map[string]bool) {
	for _, v := range arr {
		if !m[v] {
			m[v] = true
			a.recursiveInvoker(a.m[v], m)
		} else {
			continue
		}
	}
}
