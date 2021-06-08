package structures

import (
	"github.com/vielendanke/go-dstructures/structures/api"
)

type node struct {
	vertex api.EqualHashRule
	weight int64
}

func (n node) Equal(p interface{}) bool {
	if p == nil {
		return false
	}
	incNode := p.(*node)
	return n.vertex.Equal(incNode.vertex)
}

func (n node) Hash() int {
	return n.vertex.Hash()
}

type weightedGraph struct {
	m api.Map
}

func NewWeightedGraph() api.Graph {
	return &weightedGraph{m: NewHashMap(16)}
}

func (wg *weightedGraph) Size() int {
	return wg.m.Size()
}

func (wg *weightedGraph) RemoveEdge(fVertex api.EqualHashRule, sVertex api.EqualHashRule) bool {
	f := wg.m.Get(fVertex)
	s := wg.m.Get(sVertex)

	if f != nil && s != nil {
		fConv := f.([]*node)
		sConv := s.([]*node)

		for k, v := range fConv {
			if v.vertex.Equal(sVertex) {
				fConv = append(fConv[:k], fConv[k+1:]...)
				wg.m.Put(fVertex, fConv)
			}
		}
		for k, v := range sConv {
			if v.vertex.Equal(fVertex) {
				sConv = append(sConv[:k], sConv[k+1:]...)
				wg.m.Put(sVertex, sConv)
			}
		}
	} else {
		return false
	}
	return true
}

func (wg *weightedGraph) RemoveVertex(fVertex api.EqualHashRule) bool {
	fArr := wg.m.Get(fVertex)

	if fArr == nil {
		return false
	}
	convArr := fArr.([]*node)

	for _, v := range convArr {
		temp := wg.m.Get(v.vertex).([]*node)
		for k, tv := range temp {
			if tv.vertex.Equal(fVertex) {
				temp = append(temp[:k], temp[k+1:]...)
			}
		}
		wg.m.Put(v.vertex, temp)
	}
	wg.m.Remove(fVertex)
	return true
}

func (wg *weightedGraph) Contains(vertex api.EqualHashRule) bool {
	return wg.m.Get(vertex) != nil
}

func (wg *weightedGraph) GetEdges(vertex api.EqualHashRule) []api.EqualHashRule {
	vtx := wg.m.Get(vertex)

	if vtx == nil {
		return nil
	}
	edges := make([]api.EqualHashRule, 0)
	vtxEdges := vtx.([]*node)

	for _, v := range vtxEdges {
		edges = append(edges, v.vertex)
	}
	return edges
}

func (wg *weightedGraph) ToArray() []api.EqualHashRule {
	return wg.m.KeySet()
}

func (wg *weightedGraph) AddVertex(vertex api.EqualHashRule) {
	wg.m.Put(vertex, make([]*node, 0))
}

func (wg *weightedGraph) AddEdge(fVertex, sVertex api.EqualHashRule, weight int64) bool {
	f := wg.m.Get(fVertex)
	s := wg.m.Get(sVertex)
	if f == nil || s == nil {
		return false
	}
	fNodes := f.([]*node)
	sNodes := s.([]*node)

	var isExists bool

	for _, v := range fNodes {
		if v.Equal(&node{vertex: sVertex}) {
			isExists = true
		}
	}
	if !isExists {
		fNodes = append(fNodes, &node{vertex: sVertex, weight: weight})
		sNodes = append(sNodes, &node{vertex: fVertex, weight: weight})
		wg.m.Put(fVertex, fNodes)
		wg.m.Put(sVertex, sNodes)
	} else {
		return false
	}
	return true
}

// CalculateShortestPath Dijkstra algorithm
//func (wg *weightedGraph) CalculateShortestPath(start, finish string) []string {
//	nodes := queue.NewPriorityQueue(func(elem interface{}) int {
//		n := elem.(*node)
//		return int(n.weight)
//	})
//	distances := make(map
//	[string]int64)
//	previous := make(map[string]string)
//	var smallest *node
//
//	// build initial state
//	for vertex := range wg.m {
//		if vertex == start {
//			distances[vertex] = 0
//			nodes.Enqueue(&node{vertex: vertex, weight: math.MaxInt64})
//		} else {
//			distances[vertex] = math.MaxInt64
//			nodes.Enqueue(&node{vertex: vertex, weight: math.MaxInt64})
//		}
//		previous[vertex] = ""
//	}
//	// run while queue is not empty
//	var path []string
//	for nodes.Size() > 0 {
//		val, _ := nodes.Dequeue()
//		smallest = val.(*node)
//		if smallest.vertex == finish {
//			for i := len(previous); i >= 0; i-- {
//				if smallest.vertex == "" {
//					break
//				}
//				if smallest.vertex == start {
//					continue
//				}
//				path = append(path, smallest.vertex)
//				smallest.vertex = previous[smallest.vertex]
//			}
//			break
//		}
//		if smallest != nil && distances[smallest.vertex] != math.MaxInt64 {
//			//find all neighbor node
//			for _, neighbor := range wg.m[smallest.vertex] {
//				//calculate new distance to neighbor node
//				candidate := distances[smallest.vertex] + neighbor.weight
//				if candidate < distances[neighbor.vertex] {
//					// update new smallest distance
//					distances[neighbor.vertex] = candidate
//					// update how we got to next neighbor
//					previous[neighbor.vertex] = smallest.vertex
//					// enqueue with new priority
//					nodes.Enqueue(&node{vertex: neighbor.vertex, weight: candidate})
//				}
//			}
//		}
//	}
//	path = append(path, start)
//	sort.Sort(sort.Reverse(sort.StringSlice(path)))
//	return path
//}
