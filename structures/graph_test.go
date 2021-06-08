package structures

import "github.com/vielendanke/go-dstructures/structures/api"

func prepareWeightedGraph() api.Graph {
	ag := NewWeightedGraph()
	return fillGraph(ag)
}

func prepareAdjacencyGraph() api.Graph {
	ag := NewAdjacencyGraph()
	return fillGraph(ag)
}

func fillGraph(g api.Graph) api.Graph {
	for i := 1; i < 21; i++ {
		g.AddVertex(cInt(i))
		if i%2 == 0 {
			g.AddEdge(cInt(i), cInt(i-1), 0)
		}
	}
	return g
}
