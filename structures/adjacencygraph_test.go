package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdjacencyGraph_ToArray(t *testing.T) {
	ag := prepareAdjacencyGraph()

	agArr := ag.ToArray()

	assert.NotNil(t, agArr)
	assert.Equal(t, ag.Size(), len(agArr))
}

func TestAdjacencyGraph_AddVertex(t *testing.T) {
	ag := prepareAdjacencyGraph()

	ag.AddVertex(cInt(15))
	ag.AddVertex(cInt(21))

	assert.Equal(t, prepareAdjacencyGraph().Size()+1, ag.Size())
}

func TestAdjacencyGraph_AddEdge(t *testing.T) {
	ag := prepareAdjacencyGraph()

	e1 := ag.AddEdge(cInt(15), cInt(12), 0)
	e2 := ag.AddEdge(cInt(15), cInt(45), 0)
	e3 := ag.AddEdge(cInt(15), cInt(12), 0)

	assert.True(t, e1)
	assert.False(t, e2)
	assert.False(t, e3)
}

func TestAdjacencyGraph_RemoveVertex(t *testing.T) {
	ag := prepareAdjacencyGraph()

	e1 := ag.RemoveVertex(cInt(13))
	e2 := ag.RemoveVertex(cInt(53))

	assert.True(t, e1)
	assert.False(t, e2)
	assert.False(t, ag.Contains(cInt(13)))
	assert.Equal(t, prepareAdjacencyGraph().Size()-1, ag.Size())
}

func TestAdjacencyGraph_RemoveEdge(t *testing.T) {
	ag := prepareAdjacencyGraph()

	e1 := ag.RemoveEdge(cInt(12), cInt(11))
	e2 := ag.RemoveEdge(cInt(45), nil)
	edges := ag.GetEdges(cInt(12))

	assert.True(t, e1)
	assert.False(t, e2)
	for _, v := range edges {
		assert.False(t, v.Equal(cInt(11)))
	}
}

func TestAdjacencyGraph_Contains(t *testing.T) {
	ag := prepareAdjacencyGraph()

	f1 := ag.Contains(cInt(11))
	f2 := ag.Contains(cInt(41))

	assert.True(t, f1)
	assert.False(t, f2)
}

func TestAdjacencyGraph_GetEdges(t *testing.T) {
	ag := prepareAdjacencyGraph()

	f1 := ag.GetEdges(nil)
	f2 := ag.GetEdges(cInt(51))
	f3 := ag.GetEdges(cInt(11))

	assert.Nil(t, f1)
	assert.Nil(t, f2)
	assert.NotNil(t, f3)
}
