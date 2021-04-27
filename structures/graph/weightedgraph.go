package graph

import (
	"github.com/vielendanke/go-dstructures/structures/queue"
	"math"
	"sort"
)

type node struct {
	vertex string
	weight int64
}

func (n node) Equal(p interface{}) bool {
	return true
}

func (n node) Hash() int {
	return 1
}

type weightedGraph struct {
	m map[string][]*node
}

func NewWeightedGraph() *weightedGraph {
	return &weightedGraph{m: make(map[string][]*node)}
}

func (wg *weightedGraph) AddVertex(vertex string) {
	_, ok := wg.m[vertex]
	if !ok {
		wg.m[vertex] = make([]*node, 0)
	}
}

func (wg *weightedGraph) AddEdge(fVertex, sVertex string, weight int64) {
	fNodes, fOk := wg.m[fVertex]
	sNodes, sOk := wg.m[sVertex]
	if !fOk || !sOk {
		return
	}
	fNodes = append(fNodes, &node{vertex: sVertex, weight: weight})
	sNodes = append(sNodes, &node{vertex: fVertex, weight: weight})
	wg.m[fVertex] = fNodes
	wg.m[sVertex] = sNodes
}

// CalculateShortestPath Dijkstra algorithm
func (wg *weightedGraph) CalculateShortestPath(start, finish string) []string {
	nodes := queue.NewPriorityQueue(func(elem interface{}) int {
		n := elem.(*node)
		return int(n.weight)
	})
	distances := make(map[string]int64)
	previous := make(map[string]string)
	var smallest *node

	// build initial state
	for vertex := range wg.m {
		if vertex == start {
			distances[vertex] = 0
			nodes.Enqueue(&node{vertex: vertex, weight: math.MaxInt64})
		} else {
			distances[vertex] = math.MaxInt64
			nodes.Enqueue(&node{vertex: vertex, weight: math.MaxInt64})
		}
		previous[vertex] = ""
	}
	// run while queue is not empty
	var path []string
	for nodes.Size() > 0 {
		val, _ := nodes.Dequeue()
		smallest = val.(*node)
		if smallest.vertex == finish {
			for i := len(previous); i >= 0; i-- {
				if smallest.vertex == "" {
					break
				}
				if smallest.vertex == start {
					continue
				}
				path = append(path, smallest.vertex)
				smallest.vertex = previous[smallest.vertex]
			}
			break
		}
		if smallest != nil && distances[smallest.vertex] != math.MaxInt64 {
			//find all neighbor node
			for _, neighbor := range wg.m[smallest.vertex] {
				//calculate new distance to neighbor node
				candidate := distances[smallest.vertex] + neighbor.weight
				if candidate < distances[neighbor.vertex] {
					// update new smallest distance
					distances[neighbor.vertex] = candidate
					// update how we got to next neighbor
					previous[neighbor.vertex] = smallest.vertex
					// enqueue with new priority
					nodes.Enqueue(&node{vertex: neighbor.vertex, weight: candidate})
				}
			}
		}
	}
	path = append(path, start)
	sort.Sort(sort.Reverse(sort.StringSlice(path)))
	return path
}